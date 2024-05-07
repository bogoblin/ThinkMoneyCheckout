package main

import (
	. "ThinkMoneyCheckout/cart"
	"bufio"
	"fmt"
	"os"
)

func main() {
	prices := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	deals := []Deal{
		&MultiPrice{Sku: "A", Quantity: 3, Price: 130},
		&MultiPrice{Sku: "B", Quantity: 2, Price: 45},
	}

	cart := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sku := scanner.Text()
		// Check if the SKU is in the unit price map:
		_, ok := prices[sku]
		if !ok {
			fmt.Printf("SKU %s does not exist\n", sku)
			continue
		}
		currentlyInCart, _ := cart[sku]
		// Add the SKU to the cart:
		cart[sku] = currentlyInCart + 1
	}
	total, err := CalculateTotal(cart, prices, deals)
	if err != nil {
		// Print the error to stderr:
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(total)
	}
}
