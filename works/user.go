package main

import (
	"fmt"
	"time"
)

type Data struct {
	username     string
	password     string
	Count        int
	RecentAction time.Time
}

type ObjectList struct {
	li []Data
}

var objectlist = &ObjectList{li: make([]Data, 10)}

//创建用户
func (ObjList *ObjectList) MakeList() {

	objectlist.Put("zhaoriyong", "12345678")
	objectlist.Put("dingshaohua", "12345678")
	objectlist.Put("gaoziyi", "12345678")
	objectlist.Put("huangru", "12345678")
	objectlist.Put("lihuayu", "12345678")
	objectlist.Put("wangwenjing", "12345678")
	//for _,dat := range objectlist.li{
	//	fmt.Println(dat)
	//}
	//fmt.Println("+++++++++++++++++++++++++++")

}

//注册用户
/**************************************************************
规定密码格式，
除去用户名空的情况
*/

func (ObjList *ObjectList) Put(username string, password string) int {
	if 8 <= len(password) && len(password) <= 16 && 8 <= len(username) && len(username) <= 16 {
		ObjList.li = append(ObjList.li, Data{username, password, 4, time.Now()})
		return 0
	} else {
		return 1
	}
}

//获取用户信息
func (ObjList *ObjectList) Get(username string) bool {
	for _, dat := range ObjList.li {
		if dat.username == username && username != "" {
			return true
		}

	}
	return false
}

func (ObjList *ObjectList) Lock(username string) int {
	for i, dat := range ObjList.li {
		if dat.username == username && username != "" {
			ObjList.li[i].RecentAction = time.Now()
			if dat.Count > 0 {
				ObjList.li[i].Count = dat.Count - 1
				return 0
			} else {
				fmt.Println("到0")
				return -1
			}
		}
	}
	return 1
}

func (ObjList *ObjectList) isUnLock(username string) bool {
	for i, dat := range ObjList.li {
		if dat.username == username && username != "" {
			m, _ := time.ParseDuration("10m")
			if dat.Count <= 0 && time.Now().After(dat.RecentAction.Add(m)) {
				ObjList.li[i].RecentAction = time.Now()
				ObjList.li[i].Count = 4
				return true
			} else if dat.Count > 0 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

//查找用户
//*****************************************************
//用户名为空的话也可以登录。
func (ObjList *ObjectList) Find(username string, password string) bool {

	for _, dat := range ObjList.li {
		if dat.username == username && dat.password == password && username != "" {
			return true
		}
	}
	return false

}

//
////删除用户
//func (ObjList *ObjectList)delete(username string) bool {
//	return true
//
//
//}
