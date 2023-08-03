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

func AddRoute(ctx *gin.Context, c pb.TrainManagementClient) {
	routeData := domain.Route{}
	err := ctx.Bind(&routeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	fmt.Println(routeData.RouteMap)

	res, err := c.AddRoute(context.Background(), &pb.AddRouteRequest{})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new Route Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Route adding Succeded",
			"data":    res,
		})
	}
}

func AddStaion(ctx *gin.Context, c pb.TrainManagementClient) {
	stationData := domain.Station{}
	err := ctx.Bind(&stationData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddStation(context.Background(), &pb.AddStationRequest{
		Stationname: stationData.StationName,
		City:        stationData.City,
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new station Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Station adding Succeded",
			"data":    res,
		})
	}
}
