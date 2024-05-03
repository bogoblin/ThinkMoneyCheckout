package ThinkMoneyCheckout

type Deal struct {
	skus  map[string]int // SKU name to quantity
	price int            // Price after discount
}

func (deal *Deal) Apply(total int, cart map[string]int) (int, map[string]int) {
	return total, cart
}

func (deal *Deal) AppliesTo(cart map[string]int) bool {
	return false
}

type SKU struct {
	name  string
	price int
}

func CalculateTotal(cart map[string]int, unitPriceMap map[string]int, deals []Deal) int {
	total := 0

	// Apply deals
	for _, deal := range deals {
		// A deal can apply multiple times:
		for deal.AppliesTo(cart) {
			total, cart = deal.Apply(total, cart)
		}
	}
	for sku, quantity := range cart {
		price, ok := unitPriceMap[sku]
		if !ok {
			panic("No price for SKU")
		}
		total += price * quantity
	}
	return total
}
