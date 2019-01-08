// auth: kunlun
// date: 2019-01-07
// description:
package server

import "testing"

func TestStart(t *testing.T) {
	Start()
}

// 基准测试
func BenchmarkStart(b *testing.B) {
	var n int
	for n = 0; n < b.N; n++ {
		n++
	}
}
