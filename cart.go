package ThinkMoneyCheckout

import "fmt"

type Deal struct {
	skus  map[string]int // SKU name to quantity
	price int            // Price after discount
}

// ApplyTo modifies the cart to remove the required quantity of each SKU in the deal.
// Returns the total price of the deal.
func (deal *Deal) ApplyTo(cart map[string]int) int {
	for sku := range cart {
		quantityRequired, inDeal := deal.skus[sku]
		if inDeal {
			cart[sku] -= quantityRequired
		}
	}

	return deal.price
}

// Applies returns true if the cart contains all the required SKUs in the corresponding
// quantities for the deal.
func (deal *Deal) Applies(cart map[string]int) bool {
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

	// Make a deep copy of the cart, so we can modify it without affecting the original.
	cartCopy := make(map[string]int, len(cart))
	for sku, quantity := range cart {
		cartCopy[sku] = quantity
	}
	cart = cartCopy

	// ApplyTo deals
	for _, deal := range deals {
		// A deal can apply multiple times:
		for deal.Applies(cart) {
			total += deal.ApplyTo(cart)
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
