package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input interface{}

	c.BindJSON(&input)
	fmt.Println(input)
}

func GetDrinks(c *gin.Context) {
	// given a vendor id, get drinks from a drinks vendor
	idStr := c.Param("id")

	if idStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": ErrInvalidParams.Error()})
	}

	id, _ := strconv.Atoi(idStr)

	drinks, _ := getDrinks(uint64(id))

	c.IndentedJSON(http.StatusOK, drinks)
}

// GetVendors responds with the list of all vendors as JSON.
func GetVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

// GetVendorsID locates the vendor whose ID value matches the id
// parameter sent by the client, then returns that vendor as a response.
func GetVendorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	v, err := getVendorById(uint64(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, v)
	}
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
