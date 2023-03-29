package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Init() {
	// init db instance here
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var newOrder Order

	err := c.BindJSON(&newOrder)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": ErrInvalidParams.Error()})
		return
	}

	h.createOrder(newOrder)
	
	c.IndentedJSON(http.StatusOK, newOrder)
}

func (h *Handler) GetDrinks(c *gin.Context) {
	// given a vendor id, get drinks from a drinks vendor
	idStr := c.Param("id")

	if idStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": ErrInvalidParams.Error()})
	}

	id, _ := strconv.Atoi(idStr)

	drinks, _ := h.getDrinks(uint64(id))

	c.IndentedJSON(http.StatusOK, drinks)
}

func (h *Handler) GetDishById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	v, err := h.getDishById(uint64(id))

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, v)
	}
}

// GetVendors responds with the list of all vendors as JSON.
func (h *Handler) GetVendors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Vendors)
}

// GetVendorsID locates the vendor whose ID value matches the id
// parameter sent by the client, then returns that vendor as a response.
func (h *Handler) GetVendorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	v, err := h.getVendorById(uint64(id))

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, v)
	}
}

// AddVendor adds a new vendor from JSON received in the request body.
func (h *Handler) AddVendor(c *gin.Context) {
	var newVendor Vendor

	// Call BindJSON to bind the received JSON to newVendor.
	if err := c.BindJSON(&newVendor); err != nil {
		return
	}

	// Add the new vendor to the slice.
	Vendors = append(Vendors, newVendor)
	c.IndentedJSON(http.StatusCreated, newVendor)
}
