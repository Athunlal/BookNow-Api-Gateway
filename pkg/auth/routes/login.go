package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "User successfully logged in",
			"data":    res,
		})
	}

}

func ForgotPassword(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ForgotPassword(context.Background(), &pb.ForgotPasswordRequest{
		Email: user.Email,
	})
	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Forgot Password failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Otp Successfully Sent",
			"data":    res,
		})
	}
}

func ChangePassword(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		Email:    user.Email,
		Password: user.Password,
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
