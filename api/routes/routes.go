package routes

import (
	v1 "github.com/aldy505/jokes-bapak2-api/api/routes/v1"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	version1 := router.Group("/v1")
	{
		version1.GET("/", v1.SingleJoke)
		version1.GET("/today", v1.TodayJoke)
		version1.GET("/:id", v1.JokeByID)
	}

	return router
}
