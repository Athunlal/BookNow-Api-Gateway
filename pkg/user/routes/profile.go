package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/user/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ViewProfile(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	// userData := domain.User{}
	// err := ctx.Bind(&userData)
	// if err != nil {
	// 	utils.JsonInputValidation(ctx)
	// 	return
	// }

	res, err := c.ViewProfile(context.Background(), &pb.ViewProfileRequest{
		Id: int64(id),
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View Profile Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View Profile Successfull",
			"data":    res,
		})
	}
}

func EditProfile(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	userData := domain.User{}
	err := ctx.Bind(&userData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.EditProfile(context.Background(), &pb.EditProfileRequest{
		Id:       int64(id),
		Username: userData.Username,
		Gender:   userData.Gender,
		Dob:      userData.Dateofbirth,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Edit Profile Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit Profile Succeded",
			"data":    res,
		})
	}
}

func ChangePassword(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	passwordData := domain.Password{}
	err := ctx.Bind(&passwordData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ChangePassword(context.Background(), &pb.ChangeRequest{
		Id:          int64(id),
		Oldpassword: passwordData.Oldpassword,
		Newpassword: passwordData.Newpassword,
	})

	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(int(res.Status), gin.H{
			"Success": false,
			"Message": "Change Password Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(int(res.Status), gin.H{
			"Success": true,
			"Message": "Change Password Passed",
			"data":    res,
		})
	}

}

func AddAdress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.AddAddress(context.Background(), &pb.AddAddressRequest{
		Id:              int64(id),
		Type:            addressData.Type,
		Locationaddress: addressData.Locationaddress,
		Completeaddress: addressData.Completeaddress,
		Landmark:        addressData.Landmark,
		Floorno:         addressData.Floorno,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Add address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Add address successfull",
			"data":    res,
		})
	}

}

func ViewAddress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	res, err := c.ViewAddress(context.Background(), &pb.ViewAddressRequest{
		Id: int64(id),
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View address successfull",
			"data":    res,
		})
	}
}

func EditAddress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	addid, _ := strconv.Atoi(ctx.Query("addressid"))
	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.EditAddress(context.Background(), &pb.EditAddressRequest{
		Id:              int64(id),
		Addressid:       int64(addid),
		Type:            addressData.Type,
		Locationaddress: addressData.Locationaddress,
		Completeaddress: addressData.Completeaddress,
		Landmark:        addressData.Landmark,
		Floorno:         addressData.Floorno,
	})

	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Edit address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit address successful",
			"data":    res,
		})
	}
}

func ViewAddressById(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	addid, _ := strconv.Atoi(ctx.Query("addressid"))
	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ViewAddressById(context.Background(), &pb.ViewAddressByIdRequest{
		Addid: int64(addid),
		Uid:   int64(id),
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View address successfull",
			"data":    res,
		})
	}
}
