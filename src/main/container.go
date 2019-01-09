/**
* @auth    kunlun
* @date    2019-01-07 14:12
* @version v1.0
* @des     描述：
*
**/
package main

import (
	"net/http"
	"server"
)

func main() {
	//server.Start()
	http.HandleFunc("/", server.Handler)
	http.ListenAndServe("0.0.0.0:1235", nil)
}
