package response

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
)

func MapToUpdateCompartments(updateData domain.Train) []*pb.Compartment {
	compartmentsPB := make([]*pb.Compartment, 0, len(updateData.Compartment))
	for _, compartment := range updateData.Compartment {
		compartmentsPB = append(compartmentsPB, &pb.Compartment{
			Seatid: compartment.Seatid.Hex(),
		})
	}
	return compartmentsPB
}
