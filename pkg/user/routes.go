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

	// user.Use(authorize.AuthRequired)
	profile := user.Group("/profile")
	{
		profile.GET("/view", authorize.AuthRequired, svc.ViewProfile)
		profile.PUT("/edit", authorize.AuthRequired, svc.EditProfile)
		profile.PATCH("/change/password", authorize.AuthRequired, svc.ChangePassword)
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
