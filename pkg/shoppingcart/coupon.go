package shoppingcart

type Coupon struct {
	Discount          // Extend discount
	MinAmount float64 // for apply coupon limit, for total cart discount price
}
