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
		train := admin.Group("/train")
		{
			train.POST("/add", authorizeAdmin.AuthRequired, svc.AddTrain)
			train.POST("/seat/add", authorizeAdmin.AuthRequired, svc.AddSeat)
			train.GET("/view", authorizeAdmin.AuthRequired, svc.ViewTrain)
			train.PATCH("/update/route", authorizeAdmin.AuthRequired, svc.UpdateRoute)
			train.PATCH("/update/seate", authorizeAdmin.AuthRequired, svc.UpdateSeatIntoTrain)
		}
		station := admin.Group("/station")
		{
			station.POST("/add", authorizeAdmin.AuthRequired, svc.AddStaion)
			station.GET("/view", authorizeAdmin.AuthRequired, svc.ViewStation)
		}
		admin.POST("/route/add", authorizeAdmin.AuthRequired, svc.AddRoute)
	}
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
