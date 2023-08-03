package routes

import (
	"context"
	"net/http"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train/pb"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AddRoute(ctx *gin.Context, c pb.TrainManagementClient) {
	routeData := domain.Route{}
	err := ctx.Bind(&routeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	routeRequest := &pb.AddRouteRequest{
		Route: &pb.Route{
			Routeid:   routeData.RouteId.Hex(),
			Routename: routeData.RouteName,
			Routemap:  make([]*pb.RouteStation, len(routeData.RouteMap)),
		},
	}

	for i, rs := range routeData.RouteMap {
		timeProto := timestamppb.New(rs.Time)
		routeRequest.Route.Routemap[i] = &pb.RouteStation{
			Distance:  rs.Distance,
			Time:      timeProto,
			Stationid: rs.StationId.Hex(),
		}
	}

	res, err := c.AddRoute(context.Background(), routeRequest)

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
			"Message": "Route adding Succeeded",
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
