package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func UpdateSeatIntoTrain(ctx *gin.Context, c pb.TrainManagementClient) {
	updateData := domain.Train{}
	err := ctx.Bind(&updateData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.UpdateSeatIntoTrain(context.Background(), &pb.UpdateSeatIntoTrainRequest{
		Trainnumber: int64(updateData.TrainNumber),
		Seatid:      updateData.Seatsid.Hex(),
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Update seat Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Update seat Succeded",
			"data":    res,
		})
	}
}

func AddSeat(ctx *gin.Context, c pb.TrainManagementClient) {
	seatData := domain.SeatAddingData{}
	err := ctx.Bind(&seatData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	err = utils.SeatValidation(seatData.Compartment)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding seat Failed",
			"err":     err.Error(),
		})
		return
	}

	res, err := c.AddSeat(context.Background(), &pb.AddSeatRequest{
		Price:        int64(seatData.Price),
		Numberofseat: int64(seatData.NumbserOfSeat),
		Typeofseat:   seatData.TypeOfSeat,
		Compartment:  seatData.Compartment,
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding seat Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding seat Succeded",
			"data":    res,
		})
	}
}

func ViewTrain(ctx *gin.Context, c pb.TrainManagementClient) {
	res, err := c.ViewTrain(context.Background(), &pb.ViewTrainRequest{})
	if err != nil {

		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Viewing train Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Train viewing Succeded",
			"data":    res,
		})
	}
}

func AddTrain(ctx *gin.Context, c pb.TrainManagementClient) {
	trainData := domain.Train{}
	err := ctx.Bind(&trainData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddTrain(context.Background(), &pb.AddTrainRequest{
		Trainnumber: int64(trainData.TrainNumber),
		Trainname:   trainData.TrainName,
		Traintype:   trainData.TrainType,
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
			"Message": "Train adding Succeded",
			"data":    res,
		})
	}
}
