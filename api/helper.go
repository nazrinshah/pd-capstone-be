package api

import "fmt"

func (h *Handler) getVendorById(id uint64) (Vendor, error) {
	for _, v := range Vendors {
		if v.Id == id {
			return v, nil
		}
	}

	return Vendor{}, ErrVendorNotFound
}

func (h *Handler) getDrinks(id uint64) ([]Dish, error) {
	return Dishes, nil
}

func (h *Handler) getDishById(id uint64) (Dish, error) {
	for _, d := range Dishes {
		if d.Id == id {
			return d, nil
		}
	}

	return Dish{}, ErrDishNotFound
}

func (h *Handler) createOrder(order Order) {
	order.Id = uint64(len(Orders) + 1)
	order.OrderStatus = ORDER_PROCESSING
	Orders = append(Orders, order)

	fmt.Println(fmt.Sprintf("%+v", Orders))
}
