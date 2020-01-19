package shoppingcart

import "testing"

func TestGetSelfAndParentsCategories(t *testing.T) {
	expected := 200.0
	product := Product{Title: "test", Price: 100.0}
	sp := ShoppingProduct{Quantity: 10, Product: product, DiscountPrice: 80}
	price := sp.TotalDiscount()

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestHasElement(t *testing.T) {
	expected := true
	elementList := []string{"1", "2", "3"}
	hasElement := HasElement("1", elementList)

	if hasElement != expected {
		t.Errorf("Incorrect, got: %t, want: %t.", hasElement, expected)
	}
}

func TestHasElement2(t *testing.T) {
	expected := false
	elementList := []string{"10", "2", "3"}
	hasElement := HasElement("1", elementList)

	if hasElement != expected {
		t.Errorf("Incorrect, got: %t, want: %t.", hasElement, expected)
	}
}
