package main

import (
	"github.com/gin-gonic/gin"
)

func v1EndpointHandler(c *gin.Context) {
	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}
func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users", v1EndpointHandler)
	v1.GET("/users/:userId", v1EndpointHandler)
	v1.POST("/users", v1EndpointHandler)
	v1.PUT("/users/:userId", v1EndpointHandler)
	v1.DELETE("/users/:userId", v1EndpointHandler)

	router.Run(":8080")
}
