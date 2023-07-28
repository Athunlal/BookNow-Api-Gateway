package train

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
	"google.golang.org/grpc"
)

type TraiService struct {
	client pb.TrainManagementClient
}

func InitTrainService(cfg *config.Config) pb.TrainManagementClient {
	grpcConn, err := grpc.Dial(cfg.Trainsvurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewTrainManagementClient(grpcConn)
}
