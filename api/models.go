package api

type Error string

const (
	ErrVendorNotFound Error = "vendor not found"
	ErrInvalidParams  Error = "invalid params"
)

func (e Error) Error() string {
	return string(e)
}

type VendorType int

const (
	VENDORTYPE_FOOD VendorType = iota
	VENDORTYPE_DRINKS
)

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
	Type          VendorType   `json:"vendor_type"`
	Status        VendorStatus `json:"status"`
	Opening_Hours string       `json:"opening_hours"` // hh:mm-hh:mm format, assume open 7 days a week
}

type Dish struct {
	Id          uint64     `json:"id"`
	VendorId    uint64     `json:"vendor_id"`
	Name        string     `json:"name"`
	Status      DishStatus `json:"status"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Currency    string     `json:"currency"`
	ImageName   string     `json:"image_name"`
}

// Vendors slice to seed record Vendor data.
var Vendors = []Vendor{
	{Id: 1, Name: "Fish Soup", Status: VENDOR_OPEN, Opening_Hours: "10:00-21:00", Type: VENDORTYPE_FOOD},
	{Id: 2, Name: "Koi", Status: VENDOR_OPEN, Opening_Hours: "07:00-20:00", Type: VENDORTYPE_DRINKS},
}

var Dishes = []Dish{
	{
		Id:        1,
		VendorId:  2,
		Name:      "Pearl Milk Tea",
		Status:    DISH_AVAILABLE,
		Price:     5.30,
		Currency:  "SGD",
		ImageName: "fp-drink-gong-cha-pearl-milk-tea",
	},
	{
		Id:        2,
		VendorId:  2,
		Name:      "Earl Grey Milk Tea",
		Status:    DISH_AVAILABLE,
		Price:     7.30,
		Currency:  "SGD",
		ImageName: "fp-drink-gong-cha-pearl-milk-tea",
	},
	{
		Id:        3,
		VendorId:  2,
		Name:      "Mango Milk Tea",
		Status:    DISH_AVAILABLE,
		Price:     6.30,
		Currency:  "SGD",
		ImageName: "fp-drink-gong-cha-pearl-milk-tea",
	},
}
