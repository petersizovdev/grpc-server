package main

import (
	"context"
	"log"
	"net"

	"github.com/petersizovdev/grpc-server.git/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}
func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error){
	return &invoicer.CreateResponse{
		Pdf: []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8008")
	if err != nil{
		log.Fatal("unable to connect: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegister, service)

	err = serverRegister.Serve(lis)
	if err != nil{
		log.Fatal("unable to serve: %s", err)
	}
}