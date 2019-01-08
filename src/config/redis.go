// auth: kunlun
// date: 2019-01-07
// description: 缓存
package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/configor"
)

type RedisConf struct {
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
	}
}

var rcon redis.Conn

func init() {

}

func redisConn() (err error) {
	config := RedisConf{}
	configor.Load(&config, "../../config.json")

	addr := config.Redis.Host + ":" + fmt.Sprintf("%d", config.Redis.Port)
	pass := redis.DialPassword(config.Redis.Password)
	database := redis.DialDatabase(config.Redis.Database)

	rcon, err = redis.Dial("tcp", addr, pass, database)
	if err != nil {
		return err
	}
	defer rcon.Close()
	return nil
}

// TODO redis pool
