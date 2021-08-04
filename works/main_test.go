package main

import (
	//"net/http"
	//"net/http/httptest"
	"testing"
)

func TestFind1(t *testing.T) {
	objectlist.MakeList()
	if !(objectlist.Find("zhaoriyong", "12345678")) {
		t.Error()
	}
} //正常获取

func TestFind2(t *testing.T) {
	objectlist.MakeList()
	if objectlist.Find("", "") {
		t.Error()
	}
} //测试空白项

func TestGet1(t *testing.T) {
	objectlist.MakeList()
	if !(objectlist.Get("zhaoriyong")) {
		t.Error()
	}
} //正常获取

func TestGet2(t *testing.T) {
	objectlist.MakeList()
	if objectlist.Get("zhaoriyong2") {
		t.Error()
	}
} //不存在项获取

func TestPut(t *testing.T) {
	objectlist.MakeList()
	if !(objectlist.Put("zhaoriyong3", "12345678") == 0) {
		t.Error()
	}
} //正常新建

func TestPut1(t *testing.T) {
	objectlist.MakeList()
	if !(objectlist.Put("zhaoriyong1", "123456") == 1) {
		t.Error()
	}
} //违规项新建
