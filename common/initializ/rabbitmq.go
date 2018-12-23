package initializ

import (
	"airad/common/support"
	"airad/module/life/task"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
	"strconv"
	"sync"
	"time"
)

const rmqLogPrefix = "[amqp] "

type rabbitConsumer struct {
	sync.Mutex
	instance        string
	queue           string
	workFunc        func(amqp.Delivery) // 实际处理数据的函数
	consumerCount   int                 // 消费者数量
	restartOnClosed bool                // 连接断开是否尝试重连
	maxRestartTimes int
	isRunning       bool
	autoAck         bool // 是否自动应答
}

func (r *rabbitConsumer) String() string {
	return fmt.Sprintf("%+v", *r)
}

func InitRabbitMQ() {
	testTask := &rabbitConsumer{instance: "default", queue: "test", consumerCount: 2}
	testTask.restartOnClosed = true
	testTask.maxRestartTimes = 10
	testTask.workFunc = task.Test

	ccTask := &rabbitConsumer{instance: "default", queue: "bb", consumerCount: 1}
	ccTask.restartOnClosed = true
	ccTask.maxRestartTimes = 10
	ccTask.workFunc = task.Test

	// 启动消费者任务
	startTask(testTask)
	startTask(ccTask)
}

// 创建消费连接任务
func startTask(c *rabbitConsumer) error {
	logs.Info(rmqLogPrefix + "start rabbit Task " + c.String())

	// 创建连接
	rabbitDft, err := support.NewRabbitMQ(c.instance)
	// rabbitDft := support.NewRabbitMQ2(c.instance)
	if err != nil {
		logs.Error(rmqLogPrefix + "Init " + c.instance + " RabbitMQ instance ERROR")
		return err
	}
	c.isRunning = true
	// 创建多个消费者
	for i := 0; i < c.consumerCount; i++ {
		go startConsumer(rabbitDft, c)
	}
	return nil
}

func startConsumer(rabbit *support.RabbitMQ, c *rabbitConsumer) {
	logs.Info(rmqLogPrefix + "start Consumer :" + c.String())
	message := make(chan amqp.Delivery)
	if err := rabbit.ConsumeQueue(c.queue, message); nil != err {
		logs.Error(err)
		return
	}

	// 从channel接收队列消息
	for data := range message {
		if c.autoAck {
			err := data.Ack(false)
			if err != nil {
				logs.Error(err)
			}
		}
		c.workFunc(data)
	}

	// message 关闭说明连接断开
	logs.Error(rmqLogPrefix + c.queue + " consumer task Ended")
	c.isRunning = false
	go restartTask(c)
}

// 尝试重新启动消费者连接
func restartTask(c *rabbitConsumer) {
	c.Lock()
	defer c.Unlock()
	if true == c.restartOnClosed {
		if c.isRunning {
			logs.Info(rmqLogPrefix + c.queue + "have restarted by others goroutine")
			return
		}
		for i := 1; i <= c.maxRestartTimes; i++ {
			logs.Info(rmqLogPrefix + c.queue + " try restart on times " + strconv.Itoa(i))
			err := startTask(c)
			if err != nil {
				logs.Info(rmqLogPrefix + c.queue + " restarted failed on times " + strconv.Itoa(i))
				time.Sleep(10 * time.Second)
			} else {
				logs.Info(rmqLogPrefix + c.queue + " restarted Success on times " + strconv.Itoa(i))
				return
			}
		}
	}
}
