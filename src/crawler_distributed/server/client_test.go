package main

import (
	"crawler/model"
	"crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSave(t *testing.T) {
	const host = ":1234"
	go Server(host, "test_user_2")
	time.Sleep(time.Second * 2)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var result string
	err = client.Call("ItemSaveServer.Save", model.User{
		Id:   "3",
		Name: "test name 3",
		Url:  "test url 3",
	}, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s,err:%s", result, err)
	}
}
