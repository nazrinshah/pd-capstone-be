package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"pd-capstone-be/api"
)

var useDB = flag.Bool("use-db", false, "set to true if want to use MySQL db")

func main() {
	flag.Parse()

	h := api.Handler{}
	h.Init(*useDB)

	router := gin.Default()

	router.GET("/drinks/:id", h.GetDrinks)
	router.GET("/dish/:id", h.GetDishById)

	router.POST("createorder", h.CreateOrder)

	router.GET("/vendors", h.GetVendors)
	router.GET("/vendors/:id", h.GetVendorByID)

	router.POST("/addvendor", h.AddVendor)

	router.Run("localhost:8080")
}
