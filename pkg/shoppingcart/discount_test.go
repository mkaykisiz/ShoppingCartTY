package shoppingcart

import "testing"

func TestDiscount_NewPrice(t *testing.T) {
	expected := 90.0
	discount := Discount{Quantity: 10, DiscountType: Rate}
	newPrice := discount.NewPrice(100.00)

	if newPrice != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", newPrice, expected)
	}
}
func TestDiscount_NewPrice2(t *testing.T) {
	expected := 90.0
	discount := Discount{Quantity: 10, DiscountType: Amount}
	newPrice := discount.NewPrice(100.00)

	if newPrice != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", newPrice, expected)
	}
}
