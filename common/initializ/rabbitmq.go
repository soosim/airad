package initializ

import (
	"airad/common/support"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() {
	rabbitDft := support.NewRabbitMQ("default")
	err := rabbitDft.Connect()
	if err != nil {
		logs.Error("Init default RabbitMQ instance ERROR")
	}

	go startConsumer(rabbitDft, "test")
	go startConsumer(rabbitDft, "test")
	// go testNew(rabbitDft, "test")
}

func startConsumer(rabbit *support.RabbitMQ, queue string) {
	message := make(chan amqp.Delivery)
	if err := rabbit.ConsumeQueue(queue, message); nil != err {
		logs.Error(err)
		return
	}

	defer func() {
		logs.Error(queue + " Ended")
	}()

	for data := range message {
		logs.Info(data)
		data.Ack(false)
	}
}

func testNew(rabbit *support.RabbitMQ, queue string) {
	channel, err := rabbit.GetNewChannel()
	if err != nil {
		logs.Error("Init default RabbitMQ instance ERROR")
	}
	deliveries, err := channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		logs.Error("[amqp] consume queue error", err)
	}

	for d := range deliveries {
		logs.Info(
			"[NewTask] got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
}
