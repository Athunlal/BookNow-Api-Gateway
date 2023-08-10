package booking

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"google.golang.org/grpc"
)

type BookingService struct {
	client pb.BookingManagementClient
}

func InitBookingService(cfg *config.Config) pb.BookingManagementClient {

	grpcConn, err := grpc.Dial(cfg.Bookingsvurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewBookingManagementClient(grpcConn)
}
