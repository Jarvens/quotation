// auth: kunlun
// date: 2019-01-07
// description: 队列
package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type MqConfig struct {
	username string
	password string
	vhost    string
	host     string
	port     string
}

// initial MQ
func InitRMQ() {
	prop := LoadConfig()
	fmt.Printf("MQ配置文件是：%v", prop)
}

func init() {
	fmt.Println("初始化函数")
}

// load mq config
func LoadConfig() *MqConfig {
	properties := MqConfig{}
	configor.Load(&properties, "config.yml")
	return &properties
}
