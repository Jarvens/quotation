// auth: kunlun
// date: 2019-01-07
// description: 队列
package config

import (
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

type Config struct {
	Username string
	Password string
	Vhost    string
	Host     string
	Port     string
}

// initial MQ
func InitRMQ() {
	prop := loadConfig()
	fmt.Printf("config info username: %s, password: %s, vhost: %s, host: %s, port: %s ", prop.Username, prop.Password, prop.Vhost, prop.Host, prop.Port)

}

// load mq config
func loadConfig() *Config {
	var config = Config{}

	// test must use absolution address
	configor.Load(&config, "rabbitmq.yml")
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
	connect()
}

func connect() (err error) {
	prop := loadConfig()
	url := "amqp://" + prop.Username + ":" + prop.Password + "@" + prop.Host + ":" + prop.Port + "/"
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

// 闭包函数接收消息队列

func Receive() {
	if channel == nil {
		connect()
	}
	msg, err := channel.Consume(common.Queue, "", true, false, false, false, nil)
	sync := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Println(string(d))
		}
	}()
}
