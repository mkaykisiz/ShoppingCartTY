package shoppingcart

import (
	"fmt"
	"testing"
)

func TestCart_AddItem(t *testing.T) {
	expected := 1

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)

	if len(cart.shoppingProducts) != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", len(cart.shoppingProducts), expected)
	}
}
func TestCart_AddItem2(t *testing.T) {
	// one product but two request
	expected := 1

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p1, 3)

	if len(cart.shoppingProducts) != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", len(cart.shoppingProducts), expected)
	}
}

func TestCart_ApplyDiscount(t *testing.T) {
	expected := 80.00

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	//campaign
	campaign := Campaign{Category: c1, ProductLimit: 3,
		Discount: Discount{Quantity: 20.0, DiscountType: Rate}}
	campaignList := []Campaign{campaign}
	cart.ApplyDiscount(campaignList)

	total := cart.shoppingProducts[0].DiscountPrice

	if total != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", total, expected)
	}
}
func TestCart_GetTotalAmountAfterDiscount(t *testing.T) {
	expected := 240.00

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	//campaign
	campaign := Campaign{Category: c1, ProductLimit: 3,
		Discount: Discount{Quantity: 20.0, DiscountType: Rate}}
	campaignList := []Campaign{campaign}
	cart.ApplyDiscount(campaignList)
	total := cart.GetTotalAmountAfterDiscount()

	if total != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", total, expected)
	}
}

func TestCart_ApplyCoupon1(t *testing.T) {
	expected := 10.00

	cart := Cart{}
	cart.DiscountPrice = 10.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	cart.ApplyCoupon(coupon)

	price := cart.DiscountPrice
	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestCart_ApplyCoupon2(t *testing.T) {
	expected := 90.90

	cart := Cart{}
	cart.DiscountPrice = 101.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	cart.ApplyCoupon(coupon)

	price := cart.DiscountPrice
	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestCart_ApplyCoupon3(t *testing.T) {
	expected := 90.00

	cart := Cart{}
	cart.DiscountPrice = 100.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	cart.ApplyCoupon(coupon)

	price := cart.DiscountPrice
	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}
}

func TestCart_GetCouponDiscount(t *testing.T) {
	expected := 0.0

	cart := Cart{}
	cart.DiscountPrice = 10.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	price := cart.GetCouponDiscount(coupon)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}

}

func TestCart_GetCouponDiscount2(t *testing.T) {
	expected := 90.90

	cart := Cart{}
	cart.DiscountPrice = 101.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	price := cart.GetCouponDiscount(coupon)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}

}

func TestCart_GetCouponDiscount3(t *testing.T) {
	expected := 90.00

	cart := Cart{}
	cart.DiscountPrice = 100.0

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	price := cart.GetCouponDiscount(coupon)

	if price != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", price, expected)
	}

}

func TestCart_GetCampaignDiscount(t *testing.T) {
	expected := 20.00

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)
	//campaign
	campaign := Campaign{Category: c1, ProductLimit: 3,
		Discount: Discount{Quantity: 20.0, DiscountType: Rate}}
	campaignList := []Campaign{campaign}
	cart.ApplyDiscount(campaignList)

	total := cart.GetCampaignDiscount()

	if total != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", total, expected)
	}

}

func TestCart_GetDeliveryCost(t *testing.T) {
	expected := 8.99

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)

	delivery := Delivery{CostPerDelivery: 1, CostPerProduct: 2, FixedCost: 2.99}
	cost := cart.GetDeliveryCost(delivery)

	if cost != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", cost, expected)
	}
}

func TestCart_GetTotalDiscount(t *testing.T) {
	expected := 20.0

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)
	//campaign
	campaign := Campaign{Category: c1, ProductLimit: 3,
		Discount: Discount{Quantity: 20.0, DiscountType: Rate}}
	campaignList := []Campaign{campaign}
	cart.ApplyDiscount(campaignList)

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	cart.ApplyCoupon(coupon)

	total := cart.GetTotalDiscount(coupon)

	if total != expected {
		t.Errorf("Incorrect, got: %f, want: %f.", total, expected)
	}
}

func TestCart_GetProductsGroupByCategory(t *testing.T) {
	expected := 2

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)

	categories := cart.GetProductsGroupByCategory()

	if len(categories) != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", len(categories), expected)
	}
}

func TestCart_GetNumberOfDeliveries(t *testing.T) {
	expected := 1

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p1, 3)
	total := cart.GetNumberOfDeliveries()

	if total != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", total, expected)
	}
}

func TestCart_GetNumberOfDeliveries2(t *testing.T) {
	expected := 2

	c1 := Category{Title: "c1"}
	c2 := Category{Title: "c2"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c2}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)
	total := cart.GetNumberOfDeliveries()

	if total != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", total, expected)
	}
}

func TestCart_GetNumberOfProduct(t *testing.T) {
	expected := 1

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p1, 3)
	total := cart.GetNumberOfProduct()

	if total != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", total, expected)
	}
}

func TestCart_GetNumberOfProduct2(t *testing.T) {
	expected := 2

	c1 := Category{Title: "c1"}
	p1 := Product{Title: "P1", Price: 100.0, Category: c1}
	p2 := Product{Title: "P2", Price: 100.0, Category: c1}

	cart := Cart{}
	cart.AddItem(p1, 3)
	cart.AddItem(p2, 3)
	total := cart.GetNumberOfProduct()

	if total != expected {
		t.Errorf("Incorrect, got: %d, want: %d.", total, expected)
	}
}

func TestCart_Print(t *testing.T) {
	// Create Category
	edCategory := Category{Title: "Eating&Drinking"}
	foodCategory := Category{Title: "Food", ParentCategory: &edCategory}

	//products
	apple := Product{Title: "Apple", Category: foodCategory, Price: 100.00}
	almond := Product{Title: "Almond", Category: foodCategory, Price: 150.00}

	//cart
	cart := Cart{}
	cart.AddItem(apple, 3)
	cart.AddItem(almond, 1)

	//campaign
	campaign1 := Campaign{Category: foodCategory, ProductLimit: 3,
		Discount: Discount{Quantity: 20.0, DiscountType: Rate}}

	campaign2 := Campaign{Category: foodCategory, ProductLimit: 5,
		Discount: Discount{Quantity: 50.0, DiscountType: Rate}}

	campaign3 := Campaign{Category: foodCategory, ProductLimit: 5,
		Discount: Discount{Quantity: 5.0, DiscountType: Amount}}
	campaignList := []Campaign{campaign1, campaign2, campaign3}

	cart.ApplyDiscount(campaignList)

	fmt.Println("TotalAmountAfterDiscount: ", cart.GetTotalAmountAfterDiscount())

	// coupon
	coupon := Coupon{MinAmount: 100.00, Discount: Discount{Quantity: 10.0, DiscountType: Rate}}
	cart.ApplyCoupon(coupon)

	// delivery
	delivery := Delivery{CostPerDelivery: 1, CostPerProduct: 2, FixedCost: 2.99}

	//print
	cart.Print(delivery)
	t.Log(`
	TotalAmountAfterDiscount:  390
	Category:  Food
	Apple  - Unit Price:  100  x  3
	---------------
	Category:  Food
	Almond  - Unit Price:  150  x  1
	---------------
	Total Price Before Discount:  450
	Total Discount:  99
	Total Price:  351
	Delivery Cost:  7.99`)
}
