/*
Go http

路由规则：

URL 分两种:
1. 末尾是 / 表示一个子树，后面可以根其他子路径
2. 末尾不是 / 表示一个叶子节点，固定的路径
> 以 / 结尾的 URL 可以匹配它的任何子路径，比如 /images 会匹配 /images/cute-cat.jpg

采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理

如果没有找到任何匹配向，会返回 404 错误
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

// main1 -> main
func main1() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test test test test test"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w,"Hello world!")
	})

	http.HandleFunc("/time/", func(w http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}",t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080",nil)
}
