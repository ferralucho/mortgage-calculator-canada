package app

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())
	mapUrls()

	router.Run(":8082")
}
