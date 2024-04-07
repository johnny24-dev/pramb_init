package auth

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/config"
	"log"

	"google.golang.org/grpc"
)

type ServiceAuth struct {
	client pb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) pb.AuthServiceClient {
	grpc, err := grpc.Dial(cfg.Authsvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the server: ", err)
	}
	return pb.NewAuthServiceClient(grpc)
}
