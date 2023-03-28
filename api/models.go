package api

type VendorStatus int

const (
	VENDOR_CLOSED VendorStatus = iota
	VENDOR_OPEN
	VENDOR_BANNED
)

type DishStatus int

const (
	DISH_UNAVAILABLE DishStatus = iota
	DISH_AVAILABLE
)

type Vendor struct {
	Id            uint64       `json:"id"`
	Name          string       `json:"name"`
	Status        VendorStatus `json:"status"`
	Opening_Hours string       `json:"opening_hours"` // hh:mm-hh:mm format, assume open 7 days a week
}

type Dish struct {
	Id          uint64  `json:"id"`
	VendorId    uint64  `json:"vendor_id"`
	Name        string  `json:"name"`
	Status      int     `json:"status"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
}

// Vendors slice to seed record Vendor data.
var Vendors = []Vendor{
	{Id: 1, Name: "Fish Soup", Status: VENDOR_OPEN, Opening_Hours: "10:00-21:00"},
	{Id: 2, Name: "Koi", Status: VENDOR_OPEN, Opening_Hours: "07:00-20:00"},
	{Id: 3, Name: "Japanese Food", Status: VENDOR_OPEN, Opening_Hours: "11:00-20:30"},
	{Id: 4, Name: "Korean Food", Status: VENDOR_OPEN, Opening_Hours: "11:30-21:00"},
}

var Dishes = []Dish{
	{
		Id:          1,
		Name:        "",
		Status:      0,
		Price:       0,
		Description: "",
		Currency:    "",
	},
}
