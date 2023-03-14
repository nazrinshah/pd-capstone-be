package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type Vendor struct {
	Vendor_ID     string `json:"vendor_id"`
	Vendor_Name   string `json:"vendor_name"`
	Status        string `json:"status"`
	Opening_Hours string `json:"opening_hours"` // hh:mm-hh:mm format, assume open 7 days a week
	//Dish_Name     []Dish
}

type Dish struct {
	Dish_ID     string  `json:"dish_id"`
	Dish_Name   string  `json:"dish_name"`
	Quantity    int     `json:"quantity"`
	Unit_Cost   float64 `json:"unit_cost"`
	Dish_Remark string  `json:"dish_remark"`
}

type Rider struct {
	Rider_ID     string `json:"rider_id"`
	Rider_Name   string `json:"rider_name"`
	Rider_Status string `json:"rider_status"`
}

// Vendor slice to seed record Vendor data.
var Vendors = []Vendor{
	{Vendor_ID: "R001", Vendor_Name: "Fish Soup", Status: "Open", Opening_Hours: "10:00-21:00"},
	{Vendor_ID: "R002", Vendor_Name: "Mix Rice", Status: "", Opening_Hours: "07:00-20:00"},
	{Vendor_ID: "R003", Vendor_Name: "Japanese Food", Status: "", Opening_Hours: "11:00-20:30"},
	{Vendor_ID: "R004", Vendor_Name: "Korean Food", Status: "", Opening_Hours: "11:30-21:00"},
}

// getVendors responds with the list of all vendors as JSON
func getVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

func addVendor(c *g.Context) {
	var newVendor Vendor

	if err := c.BindJSON(&newVendor); err != nil {
		return
	}

	Vendors = append(Vendors, addVendor)
	c.IndentedJSON(http.StatusCreated, newVendor)

}

func main() {
	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//
	//http.ListenAndServe(":8080", nil)

	router := gin.Default()
	router.GET("/vendors", getVendors)
	router.POST("/addvendor", addVendor)

	router.Run("localhost:8080")
}
