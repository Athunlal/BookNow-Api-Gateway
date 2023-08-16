package booking

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking/routes"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func BookingSvc(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &BookingService{
		client: InitBookingService(cfg),
	}

	user := r.Group("/user")
	{
		user.POST("/train", svc.SearchTrain)
		user.GET("/train/compartment", svc.SearchCompartment)
	}
}

func (svc *BookingService) SearchTrain(ctx *gin.Context) {
	routes.SearchTrainRoute(ctx, svc.client)
}
func (svc *BookingService) SearchCompartment(ctx *gin.Context) {
	routes.SearchCompartment(ctx, svc.client)
}
