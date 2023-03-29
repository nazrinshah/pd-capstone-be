package api

type Handler struct {
	// add DB ptr here
}

type Error string

func (e Error) Error() string {
	return string(e)
}

type VendorType int

const (
	VENDORTYPE_FOOD VendorType = iota
	VENDORTYPE_DRINKS
)

type VendorStatus int32

const (
	VENDOR_CLOSED VendorStatus = iota
	VENDOR_OPEN
	VENDOR_BANNED
)

type DishStatus int32

const (
	DISH_UNAVAILABLE DishStatus = iota
	DISH_AVAILABLE
)

type Vendor struct {
	Id           uint64       `json:"id"`
	Name         string       `json:"name"`
	Type         VendorType   `json:"vendor_type"`
	Status       VendorStatus `json:"status"`
	OpeningHours string       `json:"opening_hours"` // hh:mm-hh:mm format, assume open 7 days a week
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

type OrderItem struct {
	Id        uint64  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int32   `json:"quantity"`
	Remarks   string  `json:"remarks"`
	ImageName string  `json:"image_name"`
}

type OrderStatus int32

const (
	ORDER_PROCESSING = iota
	ORDER_COOKING
	ORDER_DELIVERING
	ORDER_COMPLETED
	ORDER_CANCELLED
)

type Order struct {
	Id          uint64      `json:"id"`
	OrderStatus OrderStatus `json:"status"`
	Items       []OrderItem `json:"items"`
	Subtotal    float64     `json:"subtotal"`
	PlatformFee float64     `json:"plaform_fee"`
	DeliveryFee float64     `json:"delivery_fee"`
}

// Vendors slice to seed record Vendor data.
var Vendors = []Vendor{
	{Id: 1, Name: "Fish Soup", Status: VENDOR_OPEN, OpeningHours: "10:00-21:00", Type: VENDORTYPE_FOOD},
	{Id: 2, Name: "Koi", Status: VENDOR_OPEN, OpeningHours: "07:00-20:00", Type: VENDORTYPE_DRINKS},
}

var Dishes = []Dish{
	{
		Id:          1,
		VendorId:    2,
		Name:        "Pearl Milk Tea",
		Status:      DISH_AVAILABLE,
		Price:       4.30,
		Description: "Delicious boba with milk tea",
		Currency:    "SGD",
		ImageName:   "fp-drink-gong-cha-pearl-milk-tea",
	},
	{
		Id:          2,
		VendorId:    2,
		Name:        "Brown Sugar Milk Tea",
		Status:      DISH_AVAILABLE,
		Price:       7.30,
		Description: "Healthier choice",
		Currency:    "SGD",
		ImageName:   "fp-drink-gong-cha-brown-sugar-mlik-tea-with-pearl",
	},
	{
		Id:          3,
		VendorId:    2,
		Name:        "Strawberry Milk Tea",
		Status:      DISH_AVAILABLE,
		Price:       6.30,
		Description: "Fruity",
		Currency:    "SGD",
		ImageName:   "fp-drink-gong-cha-strawberry-milk-tea",
	},
}

var Orders = []Order{}
