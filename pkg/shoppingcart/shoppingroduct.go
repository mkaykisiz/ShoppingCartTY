package shoppingcart

type ShoppingProduct struct {
	Product       Product // cart element product
	Quantity      uint64  // cart product count
	DiscountPrice float64 // after price campaign
}

func (sp *ShoppingProduct) TotalPrice() float64 {
	return sp.DiscountPrice * float64(sp.Quantity)
}

func (sp *ShoppingProduct) TotalDiscount() float64 {
	return (sp.Product.Price - sp.DiscountPrice) * float64(sp.Quantity)
}
