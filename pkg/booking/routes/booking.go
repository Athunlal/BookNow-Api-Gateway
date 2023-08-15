package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Booking(ctx *gin.Context, c pb.BookingManagementClient) {

	trainId := ctx.Query("trainid")

	res, err := c.Booking(context.Background(), &pb.BookingRequest{
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
