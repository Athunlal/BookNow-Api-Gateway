package admin

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminMiddlewareConfig struct {
	svc *ServiceAdmin
}

func InitAdminMiddleware(Svc *ServiceAdmin) AdminMiddlewareConfig {
	return AdminMiddlewareConfig{
		svc: Svc,
	}
}

func (c *AdminMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := c.svc.client.Validate(context.Background(), &pb.ValidateRequest{
		Accesstoken: token[1],
	})
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": errs,
		})
	}
	str := strconv.FormatInt(res.Adminid, 10)
	ctx.Set("adminId", str)
	ctx.Next()
}
