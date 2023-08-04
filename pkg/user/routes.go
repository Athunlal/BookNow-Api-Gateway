package user

import (
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(user *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &UserService{
		client: InitUserService(cfg),
	}
	authorize := auth.InitAuthMiddleware(authSvc)

	// Use the auth middleware for all routes under /user
	user.Use(authorize.AuthRequired)

	// train := user.Group("/train")
	// {

	// 	train.GET("/search", svc.SearchTrain)
	// }

	profile := user.Group("/profile")
	{
		profile.GET("/view", svc.ViewProfile)
		profile.PUT("/edit", svc.EditProfile)
		profile.PATCH("/change/password", svc.ChangePassword)
	}

	address := user.Group("/address")
	{
		address.POST("/add", svc.AddAdress)
		address.GET("/view", svc.ViewAddressById)
		address.GET("/view/all", svc.ViewAddress)
		address.PUT("/edit", svc.EditAddress)
	}
}

func (svc *UserService) ViewProfile(ctx *gin.Context) {
	routes.ViewProfile(ctx, svc.client)
}

func (svc *UserService) EditProfile(ctx *gin.Context) {
	routes.EditProfile(ctx, svc.client)
}
func (svc *UserService) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}

func (svc *UserService) AddAdress(ctx *gin.Context) {
	routes.AddAdress(ctx, svc.client)
}

func (svc *UserService) ViewAddress(ctx *gin.Context) {
	routes.ViewAddress(ctx, svc.client)
}

func (svc *UserService) EditAddress(ctx *gin.Context) {
	routes.EditAddress(ctx, svc.client)
}

func (svc *UserService) ViewAddressById(ctx *gin.Context) {
	routes.ViewAddressById(ctx, svc.client)
}

// func (svc *UserService) SearchTrain(ctx *gin.Context) {
// 	routes.SeatchTrainRoute(ctx, svc.client)
// }
