package main

import (
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
)

func main() {
	router := gin.Default()

	router.GET("/vendors", api.GetVendors)
	router.GET("/vendors/:id", api.GetVendorsByID)

	router.POST("/addvendor", api.AddVendor)

	router.Run("localhost:8080")
}
