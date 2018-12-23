package controller

import (
	"airad/common/base"
	"airad/common/support"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

type TestController struct {
	base.BaseController
}

// @Title Test
// @Description test
// @Success 200 {object}
// @Failure 400 {object} \
// @router /test [post]
func (c *TestController) Test() {
	logs.Info("接收到的数据为:" + string(c.Ctx.Input.RequestBody))
	rmq, err := support.NewRabbitMQ("default")
	if err != nil {
		logs.Error("get rmq connection error")
	}

	msg := amqp.Publishing{
		Body:         []byte("soosim"),
		DeliveryMode: amqp.Persistent,
	}

	err = rmq.GetChannel().Publish("soosoo", "", false, false, msg)

	if nil != err {
		logs.Error("send failed")
	}

	c.Success(nil)
}
