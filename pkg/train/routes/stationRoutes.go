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

func ViewSation(ctx *gin.Context, c pb.TrainManagementClient) {
	res, err := c.ViewStation(context.Background(), &pb.ViewRequest{})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View Station Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View station Succeeded",
			"data":    res,
		})
	}
}

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

func UpdateTrainRoute(ctx *gin.Context, c pb.TrainManagementClient) {
	trainData := domain.Train{}
	err := ctx.Bind(&trainData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.UpdateTrainRoute(context.Background(), &pb.UpdateRequest{
		Trainnumber: int64(trainData.TrainNumber),
		Route:       trainData.Route.Hex(),
	})

	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Updating Train route Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Updating Train route Succeded",
			"data":    res,
		})
	}
}

func SearchTrainRoute(ctx *gin.Context, c pb.TrainManagementClient) {
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
			"data":    res,
		})
	}

}
