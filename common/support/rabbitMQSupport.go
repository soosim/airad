package support

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
	"time"
)

const (
	// When reconnecting to the server after connection failure
	reconnectDelay = 5 * time.Second
	// When resending messages the server didn't confirm
	resendDelay = 5 * time.Second
)

var (
	errErrorConnected   = errors.New("err connected to the queue")
	errNotConnected     = errors.New("not connected to the queue")
	errNotConfirmed     = errors.New("message not confirmed")
	errAlreadyClosed    = errors.New("already closed: not connected to the queue")
	errJsonEncodeFailed = errors.New("json encode failed for send data")
)

type MessageEntity struct {
	Exchange     string `json:"exchange"`
	Key          string `json:"key"`
	DeliveryMode string `json:"deliverymode"`
	Priority     string `json:"priority"`
	Body         string `json:"body"`
}

type ExchangeEntity struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Durable   bool   `json:"durable"`
	AutoDelte bool   `json:"autodelte"`
	NoWait    bool   `json:"nowait"`
}

type QueueEntity struct {
	Name      string `json:"name"`
	Durable   bool   `json:"durable"`
	AutoDelte bool   `json:"autodelte"`
	Exclusive bool   `json:"exclusive"`
	NoWait    bool   `json:"nowait"`
}

type QueueBindEntity struct {
	Queue    string   `json:"queue"`
	Exchange string   `json:"exchange"`
	NoWait   bool     `json:"nowait"`
	Keys     []string `json:"keys"`
}

type RabbitMQ struct {
	instance      string
	conn          *amqp.Connection
	channel       *amqp.Channel
	done          chan bool
	isConnected   bool
	notifyConfirm chan amqp.Confirmation
	notifyClose   chan *amqp.Error
}

func NewRabbitMQ(instance string) (*RabbitMQ, error) {
	rabbit := &RabbitMQ{
		instance: instance,
	}
	if ok := rabbit.connect(); ok {
		return rabbit, nil
	}
	return rabbit, errErrorConnected
}

func NewRabbitMQ2(instance string) *RabbitMQ {
	rabbit := &RabbitMQ{
		instance: instance,
	}
	go rabbit.handleReconnect()
	return rabbit
}
func (rmq *RabbitMQ) Close() error {
	if !rmq.isConnected {
		return errAlreadyClosed
	}
	err := rmq.channel.Close()
	if err != nil {
		return err
	}
	err = rmq.conn.Close()
	if err != nil {
		return err
	}
	close(rmq.done)
	rmq.isConnected = false
	return nil
}
func (rmq *RabbitMQ) handleReconnect() {
	for {
		rmq.isConnected = false
		logs.Debug("Attempting to connect")
		//  连接失败尝试重连
		for !rmq.connect() {
			logs.Error("Failed to connect. Retrying...")
			time.Sleep(reconnectDelay)
		}
		select {
		case <-rmq.done:
			return
		case <-rmq.notifyClose:
		}
	}
}

func (rmq *RabbitMQ) connect() bool {
	var err error
	urlString := getUrl(rmq.instance)
	logs.Debug(urlString)
	conn, err := amqp.Dial(urlString)
	if err != nil {
		logs.Error(err)
		return false
	}

	channel, err := conn.Channel()
	channel.Confirm(false)
	rmq.changeConnection(conn, channel)
	rmq.isConnected = true
	if err != nil {
		logs.Error("[amqp] Failed to get RabbitMQ channel", err)
	}
	return true
}

func (rmq *RabbitMQ) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	rmq.conn = connection
	rmq.channel = channel
	rmq.notifyClose = make(chan *amqp.Error)
	rmq.notifyConfirm = make(chan amqp.Confirmation)
	rmq.channel.NotifyClose(rmq.notifyClose)
	rmq.channel.NotifyPublish(rmq.notifyConfirm)
}

func getUrl(instance string) string {
	host := beego.AppConfig.String("rabbitmq::rmq." + instance + ".host")
	port := beego.AppConfig.String("rabbitmq::rmq." + instance + ".port")
	username := beego.AppConfig.String("rabbitmq::rmq." + instance + ".username")
	password := beego.AppConfig.String("rabbitmq::rmq." + instance + ".password")
	vhost := beego.AppConfig.String("rabbitmq::rmq." + instance + ".vhost")

	urlString := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		username, password, host, port, vhost,
	)
	return urlString
}

func (rmq *RabbitMQ) Connect() (*RabbitMQ, error) {
	var err error
	urlString := getUrl(rmq.instance)

	rmq.conn, err = amqp.Dial(urlString)
	if err != nil {
		logs.Error("[amqp] Failed to connect to RabbitMQ", err)
		return nil, err
	}

	rmq.channel, err = rmq.conn.Channel()
	if err != nil {
		logs.Error("[amqp] Failed to get RabbitMQ channel", err)
	}
	rmq.done = make(chan bool)
	return rmq, nil
}

func (rmq *RabbitMQ) GetChannel() *amqp.Channel {
	return rmq.channel
}

func (rmq *RabbitMQ) GetNewChannel() (channel *amqp.Channel, err error) {
	if nil == rmq.conn {
		return nil, errors.New("[amqp] connection is invalid")
	}
	return rmq.conn.Channel()
}

// 消费者
func (rmq *RabbitMQ) ConsumeQueue(queue string, message chan<- amqp.Delivery) (err error) {
	deliveries, err := rmq.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		logs.Error("[amqp] consume queue error", err)
		return err
	}

	go func(deliveries <-chan amqp.Delivery, done chan bool, message chan<- amqp.Delivery) {
		for d := range deliveries {
			logs.Info(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)
			message <- d
		}
		logs.Error("消费线程关闭")
		close(message)
		done <- false
	}(deliveries, rmq.done, message)
	return nil
}

func (rmq *RabbitMQ) Push(ex string, key string, data MessageEntity) error {
	if !rmq.isConnected {
		return errNotConnected
	}

	for {
		err := rmq.UnSafePush(ex, key, data)
		if nil != err {
			logs.Error("Push failed, Retrying...")
			continue
		}

		select {
		case confirm := <-rmq.notifyConfirm:
			if confirm.Ack {
				logs.Info("Push confirmed")
				return nil
			}
		case <-time.After(resendDelay):
		}
		logs.Error("Push didn't confirm. Retrying...")
	}
}

func (rmq *RabbitMQ) UnSafePush(ex string, key string, data MessageEntity) error {
	if !rmq.isConnected {
		return errNotConnected
	}
	if sendData, err := json.Marshal(data); err != nil {
		return errJsonEncodeFailed
	} else {
		return rmq.channel.Publish(
			ex,
			key,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        sendData,
			},
		)
	}
}
