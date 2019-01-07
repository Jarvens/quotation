// auth: kunlun
// date: 2019-01-07
// description: 队列
package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	log "utils"
)

type Config struct {
	Username string
	Password string
	Vhost    string
	Host     string
	Port     int
}

// initial MQ
func InitRMQ() {
	prop := LoadConfig()
	fmt.Printf("config info username: %s, password: %s, vhost: %s, host: %s, port: %d ", prop.Username, prop.Password, prop.Vhost, prop.Host, prop.Port)

}

// load mq config
func LoadConfig() *Config {
	var config = Config{}

	// test must use absolution address
	configor.Load(&config, "../../rabbitmq.yml")
	return &config
}

func errorHandler(err error, message string) {
	if err != nil {
		log.Info("%s: %s", message, err.Error())
	}
}
