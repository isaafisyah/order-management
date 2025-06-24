package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/config"
	"github.com/isaafisyah/order-management/app/middleware"
	"github.com/isaafisyah/order-management/app/routes"
	"github.com/isaafisyah/order-management/app/utils/logger"
)

func main()  {
	cnf := config.Get()
	//db
	db, err := config.InitDB(cnf)
	if err != nil {
		logger.Log.WithField("Module", "Main").WithError(err).Panic("Failed to connect database")
	}
	//gin engine
	r := gin.Default()
	//logger
	logger.InitLogger(cnf.Server.Env)
	//middleware
	r.Use(
		gin.Recovery(), //handle error panic
		middleware.CorsMiddleware(),
		middleware.AuthMiddleware(),
	)
	//route user
	routes.UserRoutes(db, r)
	//route product
	routes.ProductRoutes(db, r)

	//gin run
	if err := r.Run(":" + cnf.Server.Port); err != nil {
		logger.Log.WithField("Module", "Main").WithError(err).Fatal("Failed to run server :" + cnf.Server.Port)
	}
}