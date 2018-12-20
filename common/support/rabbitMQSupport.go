package support

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
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
	instance string
	conn     *amqp.Connection
	channel  *amqp.Channel
	done     chan error
}

func NewRabbitMQ(instance string) *RabbitMQ {
	if "" == instance {
		panic("New RabbitMQ Error instance Param")
	}
	return &RabbitMQ{
		instance: instance,
	}
}

func (r *RabbitMQ) Connect() (err error) {
	host := beego.AppConfig.String("rabbitmq::rmq." + r.instance + ".host")
	port := beego.AppConfig.String("rabbitmq::rmq." + r.instance + ".port")
	username := beego.AppConfig.String("rabbitmq::rmq." + r.instance + ".username")
	password := beego.AppConfig.String("rabbitmq::rmq." + r.instance + ".password")
	vhost := beego.AppConfig.String("rabbitmq::rmq." + r.instance + ".vhost")

	urlString := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		username, password, host, port, vhost,
	)

	r.conn, err = amqp.Dial(urlString)
	if err != nil {
		logs.Error("[amqp] Failed to connect to RabbitMQ", err)
		return err
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		logs.Error("[amqp] Failed to get RabbitMQ channel", err)
	}
	r.done = make(chan error)

	return nil
}

func (r *RabbitMQ) GetNewChannel() (channel *amqp.Channel, err error) {
	if nil == r.conn {
		return nil, errors.New("[amqp] connection is invalid")
	}
	return r.conn.Channel()
}

// 消费者
func (r *RabbitMQ) ConsumeQueue(queue string, message chan<- amqp.Delivery) (err error) {
	deliveries, err := r.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		logs.Error("[amqp] consume queue error", err)
		return err
	}

	go func(deliveries <-chan amqp.Delivery, done chan error, message chan<- amqp.Delivery) {
		for d := range deliveries {
			logs.Info(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)
			message <- d
		}
		done <- nil
	}(deliveries, r.done, message)
	return nil
}
