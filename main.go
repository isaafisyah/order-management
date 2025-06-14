package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/config"
)

func main()  {
	cnf := config.Get()

	_, err := config.InitDB(cnf)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := gin.Default()


	if err := r.Run(":" + cnf.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}