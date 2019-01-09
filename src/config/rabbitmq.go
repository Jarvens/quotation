// auth: kunlun
// date: 2019-01-07
// description: 队列
package config

import (
	"codec"
	"common"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/streadway/amqp"
	"time"
	log "utils"
)

var conn *amqp.Connection
var channel *amqp.Channel
var connected bool = false

type Exchange struct {
	Name  string
	Type  string
	Queue []*Queue
}

type Queue struct {
	Name string
	Key  string
}

type Config struct {
	Rabbitmq struct {
		Username    string
		Password    string
		Vhost       string
		Host        string
		Port        string
		Persistence []*Exchange
	}
}

// initial MQ
func InitRMQ() {
	config := loadConfig()
	fmt.Printf("config info username: %s, password: %s, vhost: %s, host: %s, port: %s \r\n", config.Rabbitmq.Username, config.Rabbitmq.Password, config.Rabbitmq.Vhost, config.Rabbitmq.Host, config.Rabbitmq.Port)
	for _, val := range config.Rabbitmq.Persistence {
		fmt.Printf("exchange: name: %s  type: %s\r\n", val.Name, val.Type)
		for _, q := range val.Queue {
			fmt.Printf("queue: name: %s  key: %s\r\n", q.Name, q.Key)
		}
	}
}

// load mq config
func loadConfig() *Config {
	var config = Config{}

	// test must use absolution address
	configor.Load(&config, "../../config.json")
	return &config
}

func errorHandler(err error, message string) {
	if err != nil {
		log.Info("%s: %s", message, err.Error())
	} else {
		log.Info("connect to AMQP success time: %v", time.Now())
	}
}

func init() {
	rabbitConn()
}

func rabbitConn() (err error) {
	config := loadConfig()
	url := "amqp://" + config.Rabbitmq.Username + ":" + config.Rabbitmq.Password + "@" + config.Rabbitmq.Host + ":" + config.Rabbitmq.Port + "/"
	if channel == nil {
		var config = amqp.Config{ChannelMax: 10}
		conn, err = amqp.DialConfig(url, config)
		if err != nil {
			return err
		}

		channel, err = conn.Channel()
		if err != nil {
			return err
		}
		channel.ExchangeDeclare(common.Qexchange, amqp.ExchangeDirect, true, false, false, true, nil)
		channel.QueueDeclare(common.Queue, true, false, false, false, nil)
		channel.QueueBind(common.Queue, common.Queue, common.Qexchange, false, nil)
		connected = true
	}
	return nil

}

func Publish(exchange, queue string, message []byte) {
	channel.Publish(exchange, queue, false, false, amqp.Publishing{ContentType: "text/plain", Body: message})
}

// receive message from mq
func Receive(queue string) {
	if channel == nil {
		rabbitConn()
	}
	msg, err := channel.Consume(queue, "", true, false, false, false, nil)
	errorHandler(err, "consume message error")
	sync := make(chan bool)
	go func() {
		for d := range msg {
			s := codec.ByteToString(&(d.Body))
			fmt.Println("receive message is: ", s)
		}
	}()
	<-sync
}

func exchangeType() {

}
