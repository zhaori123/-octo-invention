package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//首页
func login(w http.ResponseWriter, r *http.Request) {
	t,_ :=template.ParseFiles("index.html")

	t.Execute(w,nil)

}


//注册页面
func sign(w http.ResponseWriter, r *http.Request) {
	t,_ :=template.ParseFiles("test.html")

	t.Execute(w,nil)

}

//创建处理器函数
//登录
func handler(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//fmt.Fprintln(w,"名字",r.PostFormValue("username"))
	//fmt.Fprintln(w,"密码",r.PostFormValue("password"))

	/*增加多种情况*/
	//*************************************************************************

	find := objectlist.Find(username, password)
	if find && objectlist.isUnLock(username) {
		//w.WriteHeader(200)
		fmt.Fprintln(w, "欢迎登陆")
	} else if objectlist.Get(username) && !(objectlist.isUnLock(username)) {
		w.WriteHeader(403)
		fmt.Fprintln(w, "账户被锁定，请等待十分钟后再进行尝试")
	} else if username=="" {
		w.WriteHeader(403)
		fmt.Fprintln(w, "用户名不能为空")
	}else if objectlist.Get(username) {
		w.WriteHeader(403)
		fmt.Fprintln(w, "用户名不存在")
	}else {
		objectlist.Lock(username)
		w.WriteHeader(403)
		fmt.Fprintln(w, "用户名或密码错误")
	}
}

//注册
func regis(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//fmt.Fprintln(w,"名字",r.PostFormValue("username"))
	//fmt.Fprintln(w,"密码",r.PostFormValue("password"))
	get := objectlist.Get(username)
	if get {
		w.WriteHeader(403)
		fmt.Fprintln(w, "用户名已存在")
	} else {
		if objectlist.Put(username, password) == 0 {
			w.WriteHeader(200)
			fmt.Fprintln(w, "注册成功")
		} else {
			w.WriteHeader(403)
			fmt.Fprintln(w, "用户名或密码长度不规范")
		}
	}
}
