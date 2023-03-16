package main

import (
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
	"pd-capstone-be/db"
)

func main() {
	// database initialization
	db.InitDB()
	defer db.CloseDB()

	// router initialization
	router := gin.Default()

	router.GET("/vendors", api.GetVendors)
	router.GET("/vendors/:id", api.GetVendorsByID)

	router.POST("/addvendor", api.AddVendor)

	router.Run("localhost:8080")
}
