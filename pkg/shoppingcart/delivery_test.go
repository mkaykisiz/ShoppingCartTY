package shoppingcart

import "testing"

func TestDelivery_CalculateFor(t *testing.T) {
	// one category, one product
	expected := 5.99

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)

	delivery := Delivery{CostPerDelivery: 1, CostPerProduct: 2, FixedCost: 2.99}
	price := delivery.CalculateFor(&cart)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestDelivery_CalculateFor2(t *testing.T) {
	// two category, two product
	expected := 8.99

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)

	delivery := Delivery{CostPerDelivery: 1, CostPerProduct: 2, FixedCost: 2.99}
	price := delivery.CalculateFor(&cart)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestDelivery_CalculateFor3(t *testing.T) {
	// one category, two product
	expected := 7.99

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)

	delivery := Delivery{CostPerDelivery: 1, CostPerProduct: 2, FixedCost: 2.99}
	price := delivery.CalculateFor(&cart)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}
