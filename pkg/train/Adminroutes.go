package train

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/routes"
	"github.com/gin-gonic/gin"
)

func TrainManagementRoutes(r *gin.Engine, cfg *config.Config, adminSVC *admin.ServiceAdmin) {

	svc := &TraiService{
		client: InitTrainService(cfg),
	}
	authorizeAdmin := admin.InitAdminMiddleware(adminSVC)
	// r.Use(authorizeAdmin.AuthRequired)

	admin := r.Group("/admin")
	{
		admin.POST("/train", authorizeAdmin.AuthRequired, svc.AddTrain)
		admin.GET("/train", authorizeAdmin.AuthRequired, svc.ViewTrain)
		train := admin.Group("/train")
		{
			train.PATCH("/route", authorizeAdmin.AuthRequired, svc.UpdateRoute)
			train.PATCH("/compartment", authorizeAdmin.AuthRequired, svc.UpdateSeatIntoTrain)
		}
		admin.POST("/route", authorizeAdmin.AuthRequired, svc.AddRoute)
		admin.GET("route", authorizeAdmin.AuthRequired, svc.ViewRoutes)

		admin.POST("/compartment", authorizeAdmin.AuthRequired, svc.AddSeat)
		admin.GET("/compartment", authorizeAdmin.AuthRequired, svc.ViewComparment)

		admin.POST("/station", authorizeAdmin.AuthRequired, svc.AddStaion)
		admin.GET("/station", authorizeAdmin.AuthRequired, svc.ViewStation)

	}
}
func (svc *TraiService) ViewComparment(ctx *gin.Context) {
	routes.ViewComparment(ctx, svc.client)
}
func (svc *TraiService) ViewRoutes(ctx *gin.Context) {
	routes.ViewRoutes(ctx, svc.client)
}
func (svc *TraiService) UpdateSeatIntoTrain(ctx *gin.Context) {
	routes.UpdateSeatIntoTrain(ctx, svc.client)
}
func (svc *TraiService) AddSeat(ctx *gin.Context) {
	routes.AddSeat(ctx, svc.client)
}
func (svc *TraiService) AddTrain(ctx *gin.Context) {
	routes.AddTrain(ctx, svc.client)
}
func (svc *TraiService) AddStaion(ctx *gin.Context) {
	routes.AddStaion(ctx, svc.client)
}
func (svc *TraiService) AddRoute(ctx *gin.Context) {
	routes.AddRoute(ctx, svc.client)
}
func (svc *TraiService) UpdateRoute(ctx *gin.Context) {
	routes.UpdateTrainRoute(ctx, svc.client)
}
func (svc *TraiService) ViewTrain(ctx *gin.Context) {
	routes.ViewTrain(ctx, svc.client)
}
func (svc *TraiService) ViewStation(ctx *gin.Context) {
	routes.ViewSation(ctx, svc.client)
}
