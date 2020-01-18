package main

import (
	"context"
	"google.golang.org/grpc"
	"kigyounokagaku/proto/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{TargetCat: "tama"}
	res, err := client.GetMyCat(context.TODO(), message)
	log.Printf("result:%#v \n", res)
	log.Printf("err:%#v \n", err)

}