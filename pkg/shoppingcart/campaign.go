package shoppingcart

type Campaign struct {
	Discount              // Extend discount
	ProductLimit uint64   // cart product limit for campaign
	Category     Category // campaign category
}
