package service

import (
	"context"
	"errors"
	"kigyounokagaku/proto/pb"
)

type MyCatService struct {

}

func (s *MyCatService) GetMyCat (ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error){
	switch message.TargetCat {
	case "tama":
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}
	return nil, errors.New("Not Found YourCat")
}