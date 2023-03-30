package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/nazrinshah/pd-capstone-be/api"
	"log"
)

var useDB = flag.Bool("use-db", false, "set to true if want to use MySQL db")

func main() {
	flag.Parse()

	h := api.Handler{}
	err := h.Init(*useDB)

	if err != nil {
		log.Fatalf("error init db: %+v", err)
	}

	router := gin.Default()

	router.GET("/drinks/:id", h.GetDrinks)
	router.GET("/dish/:id", h.GetDishById)

	router.POST("createorder", h.CreateOrder)

	router.GET("/vendors", h.GetVendors)
	router.GET("/vendors/:id", h.GetVendorByID)

	router.POST("/addvendor", h.AddVendor)

	router.Run("localhost:8080")
}
