package routes

// import (
// 	"context"
// 	"net/http"

// 	"github.com/athunlal/bookNow-Api-Gateway/pkg/domain"
// 	"github.com/athunlal/bookNow-Api-Gateway/pkg/user/pb"
// 	"github.com/athunlal/bookNow-Api-Gateway/pkg/utils"
// 	"github.com/gin-gonic/gin"
// )

// func SeatchTrainRoute(ctx *gin.Context, c pb.ProfileManagementClient) {
// 	searchData := domain.SearchingTrainRequstedData{}
// 	err := ctx.Bind(&searchData)
// 	if err != nil {
// 		utils.JsonInputValidation(ctx)
// 		return
// 	}
// 	res, err := c.SearchTrain(context.Background(), &pb.SearchTrainRequest{
// 		Date:                 searchData.Date,
// 		Sourcestationid:      searchData.SourceStationid.Hex(),
// 		Destinationstationid: searchData.DestinationStationid.Hex(),
// 	})
// 	if err != nil {
// 		errs, _ := utils.ExtractError(err)
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"Success": false,
// 			"Message": "Searching Train  Failed",
// 			"err":     errs,
// 		})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"Success": true,
// 			"Message": "Searching Train  Succeded",
// 			"data":    res,
// 		})
// 	}

// }
