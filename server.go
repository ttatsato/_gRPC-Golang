package main

import (
	"google.golang.org/grpc"
	"kigyounokagaku/proto/pb"
	"kigyounokagaku/service"
	"log"
	"net"
)

func main () {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}
	// 実行したい日処理をserverに登録する。
	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
