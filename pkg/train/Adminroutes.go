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
		admin.POST("/addtrain", authorizeAdmin.AuthRequired, svc.AddTrain)
		admin.POST("/addstation", authorizeAdmin.AuthRequired, svc.AddStaion)
		admin.POST("/addroute", authorizeAdmin.AuthRequired, svc.AddRoute)
		admin.PATCH("/updateroute", authorizeAdmin.AuthRequired, svc.UpdateRoute)
		admin.GET("/view", authorizeAdmin.AuthRequired, svc.ViewTrain)
		admin.POST("/addseat", authorizeAdmin.AuthRequired, svc.AddSeat)
		admin.PATCH("/train/updateseate", authorizeAdmin.AuthRequired, svc.UpdateSeatIntoTrain)
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
