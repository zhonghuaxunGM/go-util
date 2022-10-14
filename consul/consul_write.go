package main

import (
	"fmt"
	"path"

	"github.com/hashicorp/consul/api"
)

func WriteConsulInfo(key, val, consulAddr string) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Errorf("WriteConsulInfo", "写入Consul配置发生错误:", e.(error)))
		}
	}()
	if path.IsAbs(key) {
		panic(fmt.Errorf("%s", "key应为相对路径"))
	}
	value := []byte(val)
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: consulAddr})
	if err != nil {
		panic(err)
	}
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, nil)
	if err != nil {
		panic(err)
	}
}
