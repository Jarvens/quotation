/**
* @auth    kunlun
* @date    2019-01-04 10:19
* @version v1.0
* @des     描述： 配置文件加载
*
**/

package config

import (
	"fmt"
	"os"
	"regexp"
)

type Configor struct {
	*Config
}

type Config struct {
	Environment        string
	EnvPrefix          string
	Debug              bool
	Verbose            bool
	ErrorOnUnMatchKeys bool
}

//initial
func Initial(config *Config) *Configor {

	if config == nil {
		config = &Config{}
	}

	if os.Getenv("config_debug_mode") != "" {
		config.Debug = true
	}

	if os.Getenv("config_verbose_mode") != "" {
		config.Verbose = true
	}

	return &Configor{Config: config}
}

var testRegexp = regexp.MustCompile("_test|(\\.test$)")

//get Environment
func (configor *Configor) GetEnvironment() string {
	if configor.Environment == "" {
		if env := os.Getenv("config_env"); env != "" {
			return env
		}
		if testRegexp.MatchString(os.Args[0]) {
			return "test"
		}

		return "development"
	}
	return configor.Environment
}

//get unMatchKeysError
func (confifor *Configor) GetErrorOnUnmatchedKeys() bool {
	return confifor.ErrorOnUnMatchKeys
}

func (configor *Configor) Load(config interface{}, files ...string) error {
	defer func() {
		if configor.Config.Debug || configor.Config.Verbose {
			fmt.Printf("Configuration: \n %#v\n", config)
		}
	}()

	for _, file := range configor.get {

	}
}
