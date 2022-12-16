package main

import (
	"ObjectStorage/chapter1/objects"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc来绑定对应pattern的处理函数，无法区分Method，必须自己编码区分
	http.HandleFunc("/objects/", objects.Handler)
	// 监听默认端口并启动服务
	log.Fatal(http.ListenAndServe(objects.Port, nil))
}
