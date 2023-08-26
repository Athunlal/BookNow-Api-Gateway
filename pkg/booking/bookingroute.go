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

	authorize := auth.InitAuthMiddleware(authSvc)

	user := r.Group("/user")
	{
		user.POST("/train", svc.SearchTrain)
		user.GET("/train/compartment", authorize.AuthRequired, svc.SearchCompartment)
		user.POST("/train/booking/checkout", authorize.AuthRequired, svc.Checkout)
		user.POST("/train/ticket/payment", authorize.AuthRequired, svc.Payment)
		user.POST("/wallet/create", authorize.AuthRequired, svc.CreateWallet)

	}
}

func (svc *BookingService) CreateWallet(ctx *gin.Context) {
	routes.CreateWallet(ctx, svc.client)
}
func (svc *BookingService) Payment(ctx *gin.Context) {
	routes.Payment(ctx, svc.client)
}
func (svc *BookingService) SearchTrain(ctx *gin.Context) {
	routes.SearchTrainRoute(ctx, svc.client)
}
func (svc *BookingService) SearchCompartment(ctx *gin.Context) {
	routes.SearchCompartment(ctx, svc.client)
}
func (svc *BookingService) Checkout(ctx *gin.Context) {
	routes.Checkout(ctx, svc.client)
}
