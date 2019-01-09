// auth: kunlun
// date: 2019-01-07
// description:
package config

import (
	"testing"
)

// testing unit
func TestInitRMQ(t *testing.T) {
	InitRMQ()
}

// testing publish message
//func TestPublish(t *testing.T) {
//	for i := 0; i < 1000; i++ {
//		Publish(common.Qexchange, common.Queue, []byte("hello"))
//	}
//}
