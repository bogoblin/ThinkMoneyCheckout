package main

import (
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
		&MultiPrice{sku: "A", quantity: 3, price: 130},
		&MultiPrice{sku: "B", quantity: 2, price: 45},
	}

	cart := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	for {
		sku, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// Remove the newline character:
		sku = sku[:len(sku)-1]
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
