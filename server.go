package main

import (
	"time"

	"main/api"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()

	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}
	router.Use(cors.Middleware(config))

	router.Static("/images", "./uploaded/images")

	api.Setup(router)
	router.Run(":8081")

}
