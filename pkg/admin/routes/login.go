package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AdminLogin(ctx *gin.Context, c pb.AdminServiceClient) {
	admin := domain.Admin{}
	err := ctx.Bind(&admin)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: admin.Username,
		Email:    admin.Email,
		Password: admin.Password,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Admin login failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Admin Successfully logged in",
			"data":    res,
		})
	}
}

func ChangePassword(ctx *gin.Context, c pb.AdminServiceClient) {
	admin := domain.User{}
	err := ctx.Bind(&admin)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		Id:       admin.Id,
		Password: admin.Password,
	}) // extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Change Password Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Password Successfully Changed",
			"data":    res,
		})
	}
}
