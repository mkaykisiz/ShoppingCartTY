package shoppingcart

type Discount struct {
	Quantity     float64 // quantity for rate or amount
	DiscountType DiscountType
}

func (c *Discount) NewPrice(price float64) float64 {
	// Get new price with discount
	if c.DiscountType == Rate {
		return price - (price * c.Quantity / 100)
	}
	return price - c.Quantity
}
