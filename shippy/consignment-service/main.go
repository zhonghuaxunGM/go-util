package main

import (
	"context"
	"fmt"
	"log"

	pb "shippy/consignment-service/proto/consignment"

	"github.com/gogo/protobuf/proto"
	"github.com/micro/go-micro"
)

// type IRepository interface {
// 	Create(consignment *pb.consignment)
// }

// type Reposity struct {
// 	consignment []*pb.consignment
// }

// type Service struct {
// 	repo Reposity
// }
type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// listerner, err := net.Listen("tcp", ":50051")
	// if err != null {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Println("listener:50051")
	p := &pb.Person{
		Name: "tester",
		Id:   123,
		Mail: "qwe@qwe.com",
		Numbers: []*pb.Person_PhoneNumber{
			{Number: "119", Type: 1},
			{Number: "120", Type: 2},
		},
	}
	out, err := proto.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(out)
	p2 := &pb.Person{}
	if err := proto.Unmarshal(out, p2); err != nil {
		log.Println(err)
	}
	fmt.Println(p2)

	service := micro.NewService(
		micro.Name("hellooo"),
	)
	service.Init()
	pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println("Err:", err)
	}

}
