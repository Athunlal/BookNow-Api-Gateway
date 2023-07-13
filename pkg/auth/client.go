package auth

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceAuth struct {
	client pb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) pb.AuthServiceClient {
	grpc, err := grpc.Dial(cfg.Authsvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the server :", err)
	}
	return pb.NewAuthServiceClient(grpc)
}
