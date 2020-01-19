package shoppingcart

import "testing"

func TestShoppingProduct_TotalDiscount(t *testing.T) {
	expected := 3

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2", ParentCategory: &c1}
	c3 := Category{Title: "c3", ParentCategory: &c2}

	categories := GetSelfAndParentsCategories(c3, []Category{})

	if len(categories) != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", len(categories), expected)
	}
}
func TestShoppingProduct_TotalPrice(t *testing.T) {
	expected := 800.0
	product := Product{Title: "test", Price: 100.0}
	sp := ShoppingProduct{Quantity: 10, Product: product, DiscountPrice: 80}
	price := sp.TotalPrice()

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}
