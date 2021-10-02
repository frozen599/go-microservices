package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/:name", Index)

	router.Run()
}

// Custom index handler
func Index(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}
