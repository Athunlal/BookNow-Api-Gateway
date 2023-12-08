package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/response"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ViewComparment(ctx *gin.Context, c pb.TrainManagementClient) {
	res, err := c.ViewCompartment(ctx, &pb.ViewCompartmentRequest{})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View comapartment Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View comapartment Succeded",
			"data":    res,
		})
	}
}

func UpdateSeatIntoTrain(ctx *gin.Context, c pb.TrainManagementClient) {
	updateData, err := utils.BindUpdateData(ctx)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	compartmentsPB := response.MapToUpdateCompartments(updateData)

	res, err := c.UpdateSeatIntoTrain(context.Background(), &pb.UpdateSeatIntoTrainRequest{
		Trainnumber:  int64(updateData.TrainNumber),
		Compartments: compartmentsPB,
	})

	handleUpdateSeatResponse(ctx, res, err)
}

func handleUpdateSeatResponse(ctx *gin.Context, res *pb.UpdateSeatIntoTrainResponse, err error) {
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
			"Message": "Update seat Succeeded",
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

	date := make([]*pb.Date, len(trainData.Date))
	for i, val := range trainData.Date {
		date[i] = &pb.Date{Day: val.Day}
	}

	res, err := c.AddTrain(context.Background(), &pb.AddTrainRequest{
		Trainnumber:  int64(trainData.TrainNumber),
		Trainname:    trainData.TrainName,
		Traintype:    trainData.TrainType,
		Startingtime: trainData.StartingTime,
		Endingtime:   trainData.EndingtingTime,
		Date:         date,
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
