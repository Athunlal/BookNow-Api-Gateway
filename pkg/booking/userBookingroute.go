package booking

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking/routes"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func BookingSvc(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &BookingService{
		client: InitBookingService(cfg),
	}

	Store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", Store))

	authorize := auth.InitAuthMiddleware(authSvc)

	user := r.Group("/user")
	{

		train := user.Group("/train")
		{
			train.POST("/", svc.SearchTrain)
			train.GET("/compartment", authorize.AuthRequired, svc.SearchCompartment)
			train.POST("/booking/checkout", authorize.AuthRequired, svc.Checkout)
			train.POST("/ticket/payment", authorize.AuthRequired, svc.Payment)
			train.GET("/ticket/view", authorize.AuthRequired, svc.ViewTicket)
			train.PATCH("/ticket/cancel", svc.CancelletionTicket)
			train.GET("/booking/history", authorize.AuthRequired, svc.BookingHistory)
		}

		user.POST("/wallet", authorize.AuthRequired, svc.CreateWallet)
		user.PATCH("/wallet", authorize.AuthRequired, svc.UpdateAmount)
	}
}
func (svc *BookingService) BookingHistory(ctx *gin.Context) {
	routes.BookingHistory(ctx, svc.client)
}
func (svc *BookingService) CancelletionTicket(ctx *gin.Context) {
	routes.CancelletionTicket(ctx, svc.client)
}
func (svc *BookingService) ViewTicket(ctx *gin.Context) {
	routes.ViewTicket(ctx, svc.client)
}
func (svc *BookingService) UpdateAmount(ctx *gin.Context) {
	routes.UpdateAmount(ctx, svc.client)
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
