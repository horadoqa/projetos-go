package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main()  {
	
	server := gin.Default()
	server.GET("/", handRoute)
	server.Run(":8080")
}

func handRoute(c *gin.Context)  {
	response := Response{

		Message: "WORKING!",
	}

	c.JSON(200, response)
}

