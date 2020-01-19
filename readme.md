### Example Code

	// Create Category
	edCategory := sc.Category{Title: "Eating&Drinking"}
	foodCategory := sc.Category{Title: "Food", ParentCategory: &edCategory}

	//products
	apple := sc.Product{Title: "Apple", Category: foodCategory, Price: 100.00}
	almond := sc.Product{Title: "Almond", Category: foodCategory, Price: 150.00}

	//cart
	cart := sc.Cart{}
	cart.AddItem(apple, 3)
	cart.AddItem(almond, 1)

	//campaign
	campaign1 := sc.Campaign{Category: foodCategory, ProductLimit: 3,
		Discount: sc.Discount{Quantity: 20.0, DiscountType: sc.Rate}}

	campaign2 := sc.Campaign{Category: foodCategory, ProductLimit: 5,
		Discount: sc.Discount{Quantity: 50.0, DiscountType: sc.Rate}}

	campaign3 := sc.Campaign{Category: foodCategory, ProductLimit: 5,
		Discount: sc.Discount{Quantity: 5.0, DiscountType: sc.Amount}}
	campaignList := []sc.Campaign{campaign1, campaign2, campaign3}

	cart.ApplyDiscount(campaignList)

	fmt.Println("TotalAmountAfterDiscount: ", cart.GetTotalAmountAfterDiscount())

	// coupon
	coupon := sc.Coupon{MinAmount:100.00, Discount: sc.Discount{Quantity: 10.0, DiscountType: sc.Rate}}
	cart.ApplyCoupon(coupon)

	// delivery
	delivery := sc.Delivery{CostPerDelivery:1, CostPerProduct: 2, FixedCost:2.99}

	//print
	cart.Print(delivery)

### Export

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
	Delivery Cost:  7.99


### Test

    mehmet@mkaykisiz:~/go/src/ShoppingCartTY/pkg/shoppingcart$ go test -cover
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
    Delivery Cost:  7.99
    PASS
    coverage: 100.0% of statements
    ok      ShoppingCartTY/pkg/shoppingcart 0.002s
