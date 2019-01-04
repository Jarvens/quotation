/**
* @auth    kunlun
* @date    2019-01-04 13:41
* @version v1.0
* @des     描述：
*
**/
package config

import (
	"encoding/json"
	"io"
	"strings"
)

func unmarshalJSON(data []byte, config interface{}, errorOnUnmatchedKeys bool) error {
	reader := strings.NewReader(string(data))
	decoder := json.NewDecoder(reader)

	if errorOnUnmatchedKeys {

	}

	err := decoder.Decode(config)
	if err != nil && err != io.EOF {
		return err
	}
	return nil

}
