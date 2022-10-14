package main

import (
	pb2 "shippy2/service/proto/consignment"

	"golang.org/x/net/context"
)

//
type IRepository interface {
	Create(consignment *pb2.Consignment) (*pb2.Consignment, error)
}

//
func (repo *Repository) Create(consignment *pb2.Consignment) (*pb2.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

//
type Repository struct {
	consignments []*pb2.Consignment
}

func (s *service) CreateConsignment(ctx context.Context) {

}

type service struct {
	repo Repository
}

func main() {

}
