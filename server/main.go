package main

import (
	"context"
	"github.com/gRPC_Tutorial/services/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {

}

func main()  {
	listener, err :=net.Listen("tcp", ":4040")
	if err!=nil{
		log.Fatal(err)
	}
	srv:=grpc.NewServer()
	proto.RegisterAddServiceServer(srv,&server{})
	reflection.Register(srv)
	if err:=srv.Serve(listener);err!=nil{
		log.Fatal(err)
	}
}

func (s *server)Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a,b := request.GetA(),request.GetB()
	result:= a+b
	return &proto.Response{Result: result},nil

}
func (s *server)Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a,b := request.GetA(),request.GetB()
	result:= a*b
	return &proto.Response{Result: result},nil
}