package main

import (
	"context"
	"fmt"
	pb "shippy/consignment-service/proto/consignment"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()
	helloservice := pb.NewHelloService("hellooo", service.Client())
	res, err := helloservice.Ping(context.TODO(), &pb.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
