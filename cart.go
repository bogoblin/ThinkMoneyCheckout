package ThinkMoneyCheckout

import "fmt"

type Deal struct {
	skus  map[string]int // SKU name to quantity
	price int            // Price after discount
}

func (deal *Deal) Apply(total int, cart map[string]int) (int, map[string]int) {
	// We return a new cart instead of modifying the original one
	cartAfterDealApplied := make(map[string]int)
	for sku, quantity := range cart {
		quantityRequired, inDeal := deal.skus[sku]
		if !inDeal {
			cartAfterDealApplied[sku] = quantity
		} else {
			cartAfterDealApplied[sku] = quantity - quantityRequired
		}
	}

	return total + deal.price, cartAfterDealApplied
}

func (deal *Deal) AppliesTo(cart map[string]int) bool {
	// Check if all SKUs in the deal are in the cart
	for sku, quantityRequired := range deal.skus {
		// Deals shouldn't have negative or zero quantities but let's check anyway,
		// since it would cause an infinite loop if somehow there was.
		if quantityRequired <= 0 {
			return false
		}
		numInCart, ok := cart[sku]
		if !ok || numInCart < quantityRequired {
			return false
		}
	}
	// If we get here, all SKUs in the deal are in the cart
	return true
}

type SKU struct {
	name  string
	price int
}

func CalculateTotal(cart map[string]int, unitPriceMap map[string]int, deals []Deal) (int, error) {
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
			return 0, fmt.Errorf("SKU %s does not exist", sku)
		}
		total += price * quantity
	}
	return total, nil
}
