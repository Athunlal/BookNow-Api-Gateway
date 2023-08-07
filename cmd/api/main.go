package main

import (
	"log"

	"github.com/athunlal/bookNow-Api-Gateway/pkg/admin"
	"github.com/athunlal/bookNow-Api-Gateway/pkg/auth"
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

	authAdminsvc := *admin.AdminRoutes(r, &cfg)
	train.TrainManagementRoutes(r, &cfg, &authAdminsvc)

	authsvc := *auth.RegisterRoutes(r, &cfg)
	user.RegisterUserRoutes(r, &cfg, &authsvc)
	train.UserTrainSvc(r, &cfg, &authsvc)

	// Start the server
	r.Run(cfg.Port)
}
