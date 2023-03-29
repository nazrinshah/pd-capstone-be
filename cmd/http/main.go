package main

import (
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
)

func main() {
	h := api.Handler{}

	router := gin.Default()

	router.GET("/getdrinks/:id", h.GetDrinks)

	router.POST("createorder", h.CreateOrder)

	router.GET("/vendors", h.GetVendors)
	router.GET("/vendors/:id", h.GetVendorByID)

	router.POST("/addvendor", h.AddVendor)

	router.Run("localhost:8080")
}
