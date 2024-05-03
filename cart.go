package main

import "fmt"

type Deal interface {
	// ApplyTo modifies the cart by removing items to apply the deal.
	// Returns the amount of money the deal will cost.
	ApplyTo(cart map[string]int) int

	// Applies returns true if the cart contains all the required SKUs in the corresponding
	// quantities for the deal.
	Applies(cart map[string]int) bool
}

// MultiPrice A deal of the form "buy {quantity} of {sku} for {price}".
type MultiPrice struct {
	sku      string
	quantity int
	price    int
}

// ApplyTo removes {quantity} of {sku} from the cart and returns {price}
func (deal *MultiPrice) ApplyTo(cart map[string]int) int {
	numInCart, ok := cart[deal.sku]
	// If the cart contains the SKU in the required quantity,
	if ok && numInCart >= deal.quantity {
		// Remove the items from the cart and return the deal price:
		cart[deal.sku] -= deal.quantity
		return deal.price
	} else {
		return 0
	}
}

// Applies returns true if the cart contains at least {quantity} of {sku}
func (deal *MultiPrice) Applies(cart map[string]int) bool {
	numInCart, ok := cart[deal.sku]
	return ok && numInCart >= deal.quantity
}

func CalculateTotal(cart map[string]int, unitPriceMap map[string]int, deals []Deal) (int, error) {
	total := 0

	// Make a deep copy of the cart, so we can modify it without affecting the original.
	cartCopy := make(map[string]int, len(cart))
	for sku, quantity := range cart {
		cartCopy[sku] = quantity
	}
	cart = cartCopy

	// Apply deals
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
