package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pd-capstone-be/db"
	_ "pd-capstone-be/db"
)

var (
	Vendor_ID     int
	Vendor_Name   string
	Status        string
	Opening_Hours string
)

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
var Vendors = []db.Vendor{
	{Vendor_ID: 1, Vendor_Name: "Fish Soup", Status: "Open", Opening_Hours: "10:00-21:00"},
	{Vendor_ID: 2, Vendor_Name: "Mix Rice", Status: "", Opening_Hours: "07:00-20:00"},
	{Vendor_ID: 3, Vendor_Name: "Japanese Food", Status: "", Opening_Hours: "11:00-20:30"},
	{Vendor_ID: 4, Vendor_Name: "Korean Food", Status: "", Opening_Hours: "11:30-21:00"},
}

// GetVendors responds with the list of all vendors as JSON.
func GetVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

// GetVendorsByID locates the vendor whose ID value matches the id
// parameter sent by the client, then returns that vendor as a response.
func GetVendorsByID(c *gin.Context) {

	//id := c.Param("id")
	//err := DB.QueryRow("SELECT vendor_id, vendor_name, status, opening_hours FROM restaurant.vendor WHERE id = ?", id).Scan(&Vendor_ID, &Vendor_Name, &Status, &Opening_Hours)
	//if err != nil {
	//	log.Fatal(err)
	//}

	c.JSON(http.StatusOK, gin.H{
		"vendor_id":     Vendor_ID,
		"vendor_name":   Vendor_Name,
		"status":        Status,
		"opening_hours": Opening_Hours,
	})

	c.JSON(http.StatusNotFound, gin.H{"message": "vendor not found"})
}

// AddVendor adds a new vendor from JSON received in the request body.
func AddVendor(c *gin.Context) {
	var newVendor db.Vendor

	if err := c.ShouldBindJSON(&newVendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.CreateVendor(newVendor)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully added data into the database."})
}
