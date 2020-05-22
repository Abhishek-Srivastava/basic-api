package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverStatus struct {
	Service string `json:"service"`
}

func main() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.Use(gin.Logger())
	router.GET("/status",
		func(c *gin.Context) {
			status := serverStatus{
				Service: "Hello, from MLOps",
			}
			c.JSON(http.StatusOK, status)
		})
	router.Run(":" + "5556")
}
