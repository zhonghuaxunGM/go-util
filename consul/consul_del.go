package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func DelConsulInfo(key, consulAddr string) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Errorf("DelConsulInfo", "删除Consul配置发生错误:", e.(error)))
		}
	}()
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	if err != nil {
		panic(err)
	}
	_, err = consulClient.KV().Delete(key, nil)
	if err != nil {
		panic(err)
	}
}
