package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetVendors responds with the list of all vendors as JSON.
func GetVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

// GetVendorsID locates the vendor whose ID value matches the id
// parameter sent by the client, then returns that vendor as a response.
func GetVendorByID(c *gin.Context) {
	idStr, _ := strconv.Atoi(c.Param("id"))
	id := uint64(idStr)

	// Loop over the list of vendors, looking for
	// a vendor whose ID value matches the parameter.
	for _, a := range Vendors {
		if a.Id == id {
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
