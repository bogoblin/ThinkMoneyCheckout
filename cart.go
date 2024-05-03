package ThinkMoneyCheckout

type Deal struct {
	skus  map[string]int // SKU name to quantity
	price int            // Price after discount
}

type SKU struct {
	name  string
	price int
}

func CalculateTotal(skusInCart map[SKU]int, deals []Deal) int {
	total := 0
	for sku, quantity := range skusInCart {
		total += sku.price * quantity
	}
	return total
}
