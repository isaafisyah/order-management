package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/config"
	"github.com/isaafisyah/order-management/app/routes"
	"github.com/isaafisyah/order-management/app/utils/logger"
)

func main()  {
	cnf := config.Get()
	//db
	db, err := config.InitDB(cnf)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	//gin engine
	r := gin.Default()
	//logger
	logger.InitLogger(cnf.Server.Env)
	//route
	routes.RegisterRoutes(db ,r)

	//gin run
	if err := r.Run(":" + cnf.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}