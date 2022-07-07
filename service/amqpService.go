package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	ramqp "github.com/rabbitmq/amqp091-go"
	amqp "github.com/wms3001/goAmqp"
	"github.com/wms3001/goAmqp/model"
	m "github.com/wms3001/goMongo/model"
	"io/ioutil"
	"log"
)

var aconn *model.Conn
var channel *model.Channel
var queue *model.Queue

func init() {
}

func InTestQueue(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	aconn = amqp.OpenConn()
	channel = amqp.OpenChannel(aconn)
	queue = amqp.DeclareQueue(channel, "test")
	err := channel.Channel.Publish(
		"",               // exchange
		queue.Queue.Name, // routing key
		false,            // mandatory
		false,            // immediate
		ramqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	var queueIn = model.InQueue{}
	if err != nil {
		queueIn.Code = -1
		queueIn.Message = "入队失败"
	}
	queueIn.Code = 1
	queueIn.Message = "入队成功"
	c.JSON(200, &queueIn)
}

func OutTestQueue(c *gin.Context) {
	aconn = amqp.OpenConn()
	channel = amqp.OpenChannel(aconn)
	queue = amqp.DeclareQueue(channel, "test")

	msgs, err := channel.Channel.Consume(
		queue.Queue.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Printf("%s", err)
	}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}

func InQueue() {
	log.Printf("%s", "我是主程序")
	go orderQueue()
	go InOrderQueue()
	go OutOrderQueue()
}

func orderQueue() {
	log.Printf("%s", "我是子程序")
}

func InOrderQueue() {
	aconn = amqp.OpenConn()
	channel = amqp.OpenChannel(aconn)
	queue = amqp.DeclareQueue(channel, "order")
	mData := &m.MData{}
	mData.FilterMap = map[string]interface{}{"orderStatus": 0}
	mData.Type = "and"
	many := GetManyOrders(mData)
	if many.Code == 1 {
		for _, v := range many.Data {
			mJson, _ := json.Marshal(v)
			//log.Printf("%s", mJson)
			err := channel.Channel.Publish("", queue.Queue.Name, false, false, ramqp.Publishing{ContentType: "text/plain", Body: []byte(mJson)})
			if err == nil {
				mData.FilterMap["orderId"] = v["orderId"]
				mData.UpMap = map[string]interface{}{"orderStatus": 1}
				UpdateOrderInfo(mData)
			}
		}

	}
}

func OutOrderQueue() {
	aconn = amqp.OpenConn()
	channel = amqp.OpenChannel(aconn)
	queue = amqp.DeclareQueue(channel, "order")
	msgs, err := channel.Channel.Consume(
		queue.Queue.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Printf("%s", err)
	}
	i := 1
	for d := range msgs {
		log.Printf("Received a message: %s", i)
		log.Printf("Received a message: %s", d.Body)
		i += 1
	}
}
