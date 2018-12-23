package task

import (
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

type SentenceMqTask struct {
}

func Test(d amqp.Delivery) {
	logs.Info("接收到mq消息:", d)
	d.Ack(false)
}
