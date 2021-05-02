package v1

import (
	"github.com/gin-gonic/gin"
)

func SingleJoke(c *gin.Context) {
	// get a joke from db
	// fetch the image url
	// send the image as proxied file
}

func TodayJoke(c *gin.Context) {
	// check from redis if today's joke already exists
	// send the joke if exists
	// get a new joke if it's not, then send it.
}

func JokeByID(c *gin.Context) {
	// get a joke from db by id
	// fetch image url
	// send the image as proxied file
}
