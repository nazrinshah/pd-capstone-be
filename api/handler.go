package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// Vendors slice to seed record Vendor data.
var Vendors = []Vendor{
	{Vendor_ID: "1", Vendor_Name: "Fish Soup", Status: "Open", Opening_Hours: "10:00-21:00"},
	{Vendor_ID: "2", Vendor_Name: "Mix Rice", Status: "", Opening_Hours: "07:00-20:00"},
	{Vendor_ID: "3", Vendor_Name: "Japanese Food", Status: "", Opening_Hours: "11:00-20:30"},
	{Vendor_ID: "4", Vendor_Name: "Korean Food", Status: "", Opening_Hours: "11:30-21:00"},
}

// GetVendors responds with the list of all vendors as JSON.
func GetVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

// GetVendorsID locates the vendor whose ID value matches the id
// parameter sent by the client, then returns that vendor as a response.
func GetVendorsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of vendors, looking for
	// a vendor whose ID value matches the parameter.
	for _, a := range Vendors {
		if a.Vendor_ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vendor not found"})
}

// AddVendor adds a new vendor from JSON received in the request body.
func AddVendor(c *gin.Context) {
	var newVendor Vendor

	// Call BindJSON to bind the received JSON to newVendor.
	if err := c.BindJSON(&newVendor); err != nil {
		return
	}

	// Add the new vendor to the slice.
	Vendors = append(Vendors, newVendor)
	c.IndentedJSON(http.StatusCreated, newVendor)
}
