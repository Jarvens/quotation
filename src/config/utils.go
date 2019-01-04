/**
* @auth    kunlun
* @date    2019-01-04 11:27
* @version v1.0
* @des     描述：工具类
*
**/
package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type UnmatchedTomlKeysError struct {
	Keys []toml.Key
}

func (e *UnmatchedTomlKeysError) Error() string {

	return fmt.Sprintf("There are keys in the config file that do not match any field in the given struct: %v", e.Keys)
}

func (configor *Configor) getEnvPrefix(config interface{}) string {
	if configor.Config.EnvPrefix == "" {
		if prefix := os.Getenv("configor_env_prefix"); prefix != "" {
			return prefix
		}
		return "configor"
	}
	return configor.Config.EnvPrefix
}

func getConfigurationFileWithEnvPrefix(file, env string) (string, error) {
	var (
		envFile string
		extName = path.Ext(file)
	)
	if extName == "" {
		envFile = fmt.Sprintf("%v.%v", file, env)
	} else {
		envFile = fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extName), env, extName)
	}

	if fileInfo, err := os.Stat(envFile); err == nil && fileInfo.Mode().IsRegular() {
		return envFile, nil
	}

	return "", fmt.Errorf("Faild  find file: %v", file)

}

func processFile(config interface{}, file string, errorOnUnmatchedKeys bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		if errorOnUnmatchedKeys {
			return yaml.UnmarshalStrict(data, config)
		}
		return yaml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".toml"):
		return unmarshalToml(data, config, errorOnUnmatchedKeys)
	case strings.HasSuffix(file, ".json"):
		return unmarshalJSON(data, config, errorOnUnmatchedKeys)
	default:
		if err := unmarshalToml(data, config, errorOnUnmatchedKeys); err == nil {
			return nil
		} else if errunmatchedKeys, ok := err.(*UnmatchedTomlKeysError); ok {
			return errunmatchedKeys
		}
		if err := unmarshalJSON(data, config, errorOnUnmatchedKeys); err == nil {
			return nil
		} else if strings.Contains(err.Error(), "json: unknown field") {
			return err
		}

		var yamlError error
		if errorOnUnmatchedKeys {
			yamlError = yaml.UnmarshalStrict(data, config)
		} else {
			yamlError = yaml.Unmarshal(data, config)
		}

		if yamlError == nil {
			return nil
		} else if yErr, ok := yamlError.(*yaml.TypeError); ok {
			return yErr
		}

		return errors.New("failed to decode config")

	}
}

func unmarshalToml(data []byte, config interface{}, errorOnUnmatchedKeys bool) error {
	metadata, err := toml.Decode(string(data), config)
	if err != nil && len(metadata.Undecoded()) > 0 && errorOnUnmatchedKeys {
		return &UnmatchedTomlKeysError{Keys: metadata.Undecoded()}
	}
	return err
}

//get configuration files
func (configor *Configor) getConfigurationFiles(files ...string) []string {

	var results []string
	//print environment
	if configor.Config.Debug || configor.Config.Verbose {
		fmt.Printf("Current environment: %v\n", configor.GetEnvironment())
	}

	for i := len(files) - 1; i >= 0; i-- {
		fountFile := false
		file := files[i]

		//check
		if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
			fountFile = true
			results = append(results, file)
		}

	}
	return nil
}
