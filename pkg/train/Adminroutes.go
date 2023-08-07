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

	r.Use(authorizeAdmin.AuthRequired)

	admin := r.Group("/admin")
	{
		admin.POST("/addtrain", svc.AddTrain)
		admin.POST("/addstation", svc.AddStaion)
		admin.POST("/addroute", svc.AddRoute)
		admin.PATCH("/updateroute", svc.UpdateRoute)
		admin.GET("/view", svc.ViewTrain)
	}

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
