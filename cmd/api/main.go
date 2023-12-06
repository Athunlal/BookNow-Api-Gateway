package main

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/booking"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/config"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/train"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/user"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load the config file : ", err)
	}
	r := gin.Default()

	authsvc := auth.RegisterRoutes(r, &cfg)
	authAdminsvc := admin.AdminRoutes(r, &cfg)

	user.RegisterUserRoutes(r, &cfg, authsvc)
	train.UserTrainSvc(r, &cfg, authsvc)
	booking.BookingSvc(r, &cfg, authsvc)

	train.TrainManagementRoutes(r, &cfg, authAdminsvc)

	// Start the server
	r.Run(cfg.Port)
}
