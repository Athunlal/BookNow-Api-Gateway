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

func ViewTicket(ctx *gin.Context, c pb.BookingManagementClient) {
	id := ctx.Query("ticketId")

	res, err := c.ViewTicket(context.Background(), &pb.ViewTicketRequest{
		Ticketid: id,
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View ticket Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View ticket Succeded",
			"data":    res,
		})
	}

}

func UpdateAmount(ctx *gin.Context, c pb.BookingManagementClient) {
	Wallet := domain.UserWallet{}
	if err := ctx.Bind(&Wallet); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	id, err := strconv.Atoi(ctx.GetString("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Update wallet Failed",
			"err":     err,
		})
	}

	res, err := c.UpdateAmount(context.Background(), &pb.UpdateAmountRequest{
		Userid:        int64(id),
		WalletBalance: float32(Wallet.WalletBalance),
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Update wallet Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Update wallet Succeded",
			"data":    res,
		})
	}

}

func CreateWallet(ctx *gin.Context, c pb.BookingManagementClient) {
	Wallet := &domain.UserWallet{}
	if err := ctx.Bind(&Wallet); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	id, err := strconv.Atoi(ctx.GetString("userId"))

	res, err := c.CreateWallet(context.Background(), &pb.CreateWalletRequest{
		WalletBalance: float32(Wallet.WalletBalance),
		Userid:        int64(id),
	})
	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Create wallet Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Create wallet Succeded",
			"data":    res,
		})
	}
}

func Payment(ctx *gin.Context, c pb.BookingManagementClient) {
	paymentData := &domain.Payment{}
	if err := ctx.Bind(&paymentData); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	id, err := strconv.Atoi(ctx.GetString("userId"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"Success": false,
			"Message": "conversion faild",
			"Error":   err.Error(),
		})
	}

	res, err := c.Payment(context.Background(), &pb.PaymentRequest{
		Userid:    int64(id),
		PNRnumber: paymentData.PNRnumber,
	})

	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Payment Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Payment Succeded",
			"data":    res,
		})
	}

}

func Checkout(ctx *gin.Context, c pb.BookingManagementClient) {

	bookingData := &domain.BookingData{}
	if err := ctx.Bind(&bookingData); err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	var travelers []*pb.Travelers
	for _, traveler := range bookingData.Travelers {
		travelerInstance := &pb.Travelers{
			Travelername: traveler.Travelername,
		}
		travelers = append(travelers, travelerInstance)
	}

	id, err := strconv.Atoi(ctx.GetString("userId"))
	res, err := c.Checkout(context.Background(), &pb.CheckoutRequest{
		Compartmentid:        bookingData.CompartmentId,
		TrainId:              bookingData.TrainId,
		Sourcestationid:      bookingData.SourceStationid.Hex(),
		Destinationstationid: bookingData.DestinationStationid.Hex(),
		Userid:               int64(id),
		Travelers:            travelers,
	})

	if err != nil {
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Checkout Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Checkout Succeded",
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
