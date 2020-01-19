package shoppingcart

type Delivery struct {
	CostPerDelivery float64 // per category
	CostPerProduct  float64 // per product
	FixedCost       float64 // fixed fee
}

func (d *Delivery) CalculateFor(cart *Cart) float64 {
	numberOfDeliveries := float64(cart.getNumberOfDeliveries())
	numberOfProduct := float64(cart.getNumberOfProduct())
	return (d.CostPerDelivery * numberOfDeliveries) + (d.CostPerProduct * numberOfProduct) + d.FixedCost
}
