package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddTrain(ctx *gin.Context, c pb.TrainManagementClient) {
	fmt.Println("reached Add train  -------------------->>>>>>.")

	trainData := domain.Train{}
	err := ctx.Bind(trainData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddTrain(context.Background(), &pb.AddTrainRequest{
		Trainid:     int64(trainData.TrainId),
		Trainnumber: int64(trainData.TrainNumber),
		Trainname:   trainData.TrainName,
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new train Failed",
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
