package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Checkout(ctx *gin.Context, c pb.BookingManagementClient) {

	bookingData := &domain.BookingData{}
	if err := ctx.Bind(&bookingData); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	id, err := strconv.Atoi(ctx.GetString("userId"))
	res, err := c.Checkout(context.Background(), &pb.CheckoutRequest{
		Compartmentid: bookingData.CompartmentId,
		TrainId:       bookingData.TrainId,
		Userid:        int64(id),
	})

	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Ticket booking Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Ticket booking Succeded",
			"data":    res,
		})
	}
}

func SearchCompartment(ctx *gin.Context, c pb.BookingManagementClient) {

	trainId := ctx.Query("trainid")

	res, err := c.SearchCompartment(context.Background(), &pb.SearchCompartmentRequest{
		Trainid: trainId,
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Searching Train  Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Searching Train  Succeded",
			"data":    res,
		})
	}
}

func SearchTrainRoute(ctx *gin.Context, c pb.BookingManagementClient) {
	searchData := domain.SearchingTrainRequstedData{}
	err := ctx.Bind(&searchData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.SearchTrain(context.Background(), &pb.SearchTrainRequest{
		Date:                 searchData.Date,
		Sourcestationid:      searchData.SourceStationid.Hex(),
		Destinationstationid: searchData.DestinationStationid.Hex(),
	})

	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Searching Train  Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Searching Train  Succeded",
			"data":    res.Traindata,
		})
	}

}
