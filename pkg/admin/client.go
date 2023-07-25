package admin

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceAdmin struct {
	client pb.AdminServiceClient
}

func InitServiceClient(cfg *config.Config) pb.AdminServiceClient {
	grpc, err := grpc.Dial(cfg.Authsvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the server: ", err)
	}
	return pb.NewAdminServiceClient(grpc)
}
