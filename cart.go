package ThinkMoneyCheckout

type Deal struct {
	skus  map[SKU]int // SKU to quantity
	price int         // Price after discount
}

type SKU struct {
	name  string
	price int
}

func CalculateTotal(skusInCart map[SKU]int, deals []Deal) int {
	return 0
}
