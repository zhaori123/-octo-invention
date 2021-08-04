package main

import (
	"net/http"
)

func main() {
	objectlist.MakeList()

	//去登陆
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/regis", regis)
	http.HandleFunc("/login", login)
	http.HandleFunc("/sign", sign)
	//创建路由
	http.ListenAndServe(":8080", nil)

}
