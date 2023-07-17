package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := p.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Phone:    body.Phone,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Registering the User Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully sent the OTP",
			"data":    res,
		})
	}

}

func RegisterValidate(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := p.RegisterValidate(context.Background(), &pb.RegisterValidateRequest{
		Email: body.Email,
		Otp:   body.Otp,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "OTP Verification Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully registered the user",
			"data":    res,
		})
	}

}
