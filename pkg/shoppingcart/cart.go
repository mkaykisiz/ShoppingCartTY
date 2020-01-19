package shoppingcart

import "fmt"

type Cart struct {
	shoppingProducts []ShoppingProduct // Cart Product info
	TotalPrice       float64           // cart Total price
	DiscountPrice    float64           //cart total discount price, after discount
}

type CategoryGroup struct {
	category         Category          // selected category
	shoppingProducts []ShoppingProduct // category products for category grouping
}

func (c *Cart) AddItem(product Product, quantity uint64) {
	hasProduct := false // if product already exist in cart; true
	totalPrice := product.Price

	// Add product or if exist increase quantity
	for i, value := range c.shoppingProducts {
		if value.Product.Title == product.Title {
			hasProduct = true
			c.shoppingProducts[i].Quantity += quantity
		}
	}
	if !hasProduct {
		c.shoppingProducts = append(c.shoppingProducts, ShoppingProduct{
			Product: product, Quantity: quantity,
			DiscountPrice: totalPrice, // discount price default total price
		})
	}

}

func (c *Cart) ApplyDiscount(campaigns []Campaign) {
	for i, sp := range c.shoppingProducts {
		// get category and parent for product for campaign
		categories := GetSelfAndParentsCategories(sp.Product.Category, []Category{})

		for _, category := range categories {
			for _, campaign := range campaigns {
				// category and limit control
				if campaign.Category == category && campaign.ProductLimit <= sp.Quantity {
					// check discount price for min price
					price := sp.Product.Price
					newPrice := campaign.NewPrice(price)
					if sp.DiscountPrice > newPrice {
						c.shoppingProducts[i].DiscountPrice = newPrice
					}
				}
			}
		}
	}
}

func (c *Cart) GetTotalAmountAfterDiscount() float64 {
	var totalPrice, totalDiscountPrice float64
	// prepare price for total and discount prices
	for _, sp := range c.shoppingProducts {
		totalPrice += sp.Product.Price * float64(sp.Quantity)
		totalDiscountPrice += sp.DiscountPrice * float64(sp.Quantity)
	}
	// update cart pricing
	c.TotalPrice = totalPrice
	c.DiscountPrice = totalDiscountPrice
	return c.DiscountPrice
}

func (c *Cart) ApplyCoupon(coupon Coupon) {
	if coupon.MinAmount <= c.DiscountPrice {
		c.DiscountPrice = coupon.NewPrice(c.DiscountPrice)
	}
}

func (c *Cart) GetCouponDiscount(coupon Coupon) float64 {
	if coupon.MinAmount <= c.DiscountPrice {
		return coupon.NewPrice(c.DiscountPrice)
	}
	return 0.0
}

func (c *Cart) GetCampaignDiscount() float64 {
	//get previously calculated prices
	var totalPrice, totalDiscountPrice float64
	for _, sp := range c.shoppingProducts {
		totalPrice += sp.Product.Price
		totalDiscountPrice += sp.DiscountPrice
	}
	return totalPrice - totalDiscountPrice
}

func (c *Cart) GetDeliveryCost(delivery Delivery) float64 {
	return delivery.CalculateFor(c)
}

func (c *Cart) GetTotalDiscount(coupon Coupon) float64 {
	// campaign and coupon
	return c.GetCampaignDiscount() + c.GetCouponDiscount(coupon)
}

func (c *Cart) GetProductsGroupByCategory() []CategoryGroup {
	//distinct category and products
	var categories []CategoryGroup
	for _, sp := range c.shoppingProducts {
		for _, category := range categories {
			if sp.Product.Category == category.category {
				category.shoppingProducts = append(category.shoppingProducts, sp)

			}
		}
		categories = append(categories, CategoryGroup{
			category: sp.Product.Category, shoppingProducts: []ShoppingProduct{sp},
		})
	}
	return categories
}

func (c *Cart) GetNumberOfDeliveries() int {
	//distinct category count
	var categories []string
	for _, sp := range c.shoppingProducts {
		if !HasElement(sp.Product.Category.Title, categories) {
			categories = append(categories, sp.Product.Category.Title)
		}
	}
	return len(categories)
}

func (c *Cart) GetNumberOfProduct() int {
	// distinct product count
	var products []string
	for _, sp := range c.shoppingProducts {
		if !HasElement(sp.Product.Title, products) {
			products = append(products, sp.Product.Title)
		}
	}
	return len(products)
}

func (c *Cart) Print(delivery Delivery) {
	productsGroupByCategory := c.GetProductsGroupByCategory()
	for _, categoryGroup := range productsGroupByCategory {
		// category info
		fmt.Println("Category: ", categoryGroup.category.Title)
		for _, sp := range categoryGroup.shoppingProducts {
			// products
			fmt.Println(sp.Product.Title, " - Unit Price: ", sp.Product.Price, " x ", sp.Quantity)
		}
		fmt.Println("---------------")
	}
	// last infos
	fmt.Println("Total Price Before Discount: ", c.TotalPrice)
	fmt.Println("Total Discount: ", c.TotalPrice-c.DiscountPrice)
	fmt.Println("Total Price: ", c.DiscountPrice)
	fmt.Println("Delivery Cost: ", c.GetDeliveryCost(delivery))
}
