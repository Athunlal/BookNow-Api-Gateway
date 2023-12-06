package train

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/routes"
	"github.com/gin-gonic/gin"
)

func UserTrainSvc(r *gin.Engine, cfg *config.Config, userSVC *auth.ServiceAuth) {

	svc := &TraiService{
		client: InitTrainService(cfg),
	}
	authorize := auth.InitAuthMiddleware(userSVC)
	// r.Use(authorize.AuthRequired)

	user := r.Group("/user")
	{
		train := user.Group("/train")
		{
			train.POST("/search", authorize.AuthRequired, svc.SearchTrain)
			train.GET("/search", authorize.AuthRequired, svc.SearchTrainbyName)
		}
		station := user.Group("/station")
		{
			station.GET("/view", authorize.AuthRequired, svc.ViewStationByUser)
		}
	}
}

func (svc *TraiService) ViewStationByUser(ctx *gin.Context) {
	routes.ViewSation(ctx, svc.client)
}

func (svc *TraiService) SearchTrain(ctx *gin.Context) {
	routes.SearchTrainRoute(ctx, svc.client)
}
func (svc *TraiService) SearchTrainbyName(ctx *gin.Context) {
	routes.SearchTrainbyName(ctx, svc.client)
}
