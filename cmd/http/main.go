package main

import (
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
)

func main() {
	h := api.Handler{}
	h.Init()

	router := gin.Default()

	router.GET("/drinks/:id", h.GetDrinks)
	router.GET("/dish/:id", h.GetDishById)

	router.POST("createorder", h.CreateOrder)

	router.GET("/vendors", h.GetVendors)
	router.GET("/vendors/:id", h.GetVendorByID)

	router.POST("/addvendor", h.AddVendor)

	router.Run("localhost:8080")
}
