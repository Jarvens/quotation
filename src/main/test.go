// auth: kunlun
// date: 2019-01-09
// description:
package main

import (
	"config"
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(reflect.ValueOf(queue).Cap())
}
