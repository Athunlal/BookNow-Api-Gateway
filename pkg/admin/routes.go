package admin

import (
	"fmt"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin/routes"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine, cfg *config.Config) *ServiceAdmin {
	fmt.Println("reached here -------->>")

	svc := &ServiceAdmin{
		client: InitServiceClient(cfg),
	}
	admin := r.Group("/admin")
	{
		admin.POST("/login", svc.AdminLogin)
		admin.PATCH("/changepassword", svc.ChangePassword)
	}
	return svc
}

func (svc *ServiceAdmin) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.client)
}
func (svc *ServiceAdmin) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}
