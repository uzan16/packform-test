package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uzan16/packform-test/config"
	"github.com/uzan16/packform-test/controllers"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	router := gin.Default()

	// r.Use(static.Serve("/order", static.LocalFile("./webapp/dist", false)))

	controllers.NewOrdersController(db, router).Initialize()
	router.Run()
}
