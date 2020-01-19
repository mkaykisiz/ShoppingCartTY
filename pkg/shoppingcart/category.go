package shoppingcart

type Category struct {
	Title          string
	ParentCategory *Category
}
