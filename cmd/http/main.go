package main

import (
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
)

func main() {
	router := gin.Default()

	router.GET("/getdrinks/:id", api.GetDrinks)

	router.POST("createorder", api.CreateOrder)

	router.GET("/vendors", api.GetVendors)
	router.GET("/vendors/:id", api.GetVendorByID)

	router.POST("/addvendor", api.AddVendor)

	router.Run("localhost:8080")
}
