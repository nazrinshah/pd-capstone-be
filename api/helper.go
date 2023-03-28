package api

func getVendorById(id uint64) (Vendor, error) {
	for _, v := range Vendors {
		if v.Id == id {
			return v, nil
		}
	}

	return Vendor{}, ErrVendorNotFound
}

func getDrinks(id uint64) ([]Dish, error) {
	return Dishes, nil
}
