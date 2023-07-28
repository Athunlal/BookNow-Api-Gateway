package train

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/routes"
	"github.com/gin-gonic/gin"
)

func AddTrainRoutes(train *gin.Engine, cfg *config.Config, adminSVC *admin.ServiceAdmin) {
	svc := &TraiService{
		client: InitTrainService(cfg),
	}
	authorize := admin.InitAdminMiddleware(adminSVC)

	train.Use(authorize.AuthRequired)

	admin := train.Group("/train")
	{
		admin.POST("/addtrain", svc.AddTrain)
	}

}

func (svc *TraiService) AddTrain(ctx *gin.Context) {

	routes.AddTrain(ctx, svc.client)
}
