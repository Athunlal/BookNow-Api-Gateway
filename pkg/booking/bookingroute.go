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

	user := r.Group("/user/train")
	{
		user.POST("/search", svc.SearchTrain)
		user.GET("/booking", svc.Booking)
	}
}

func (svc *BookingService) SearchTrain(ctx *gin.Context) {
	routes.SearchTrainRoute(ctx, svc.client)
}
func (svc *BookingService) Booking(ctx *gin.Context) {
	routes.Booking(ctx, svc.client)
}
