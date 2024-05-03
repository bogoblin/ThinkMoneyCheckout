package main

import (
	"fmt"
	"testing"
)

func TestCalculateTotal(t *testing.T) {
	type args struct {
		skusInCart   map[string]int
		unitPriceMap map[string]int
		deals        []Deal
	}
	exampleUnitPriceMap := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	exampleDeals := []Deal{
		&MultiPrice{sku: "A", quantity: 3, price: 130},
		&MultiPrice{sku: "B", quantity: 2, price: 45},
	}
	tests := []struct {
		name  string
		args  args
		want  int
		error error
	}{
		{
			"Returns 0 if no items in cart",
			args{
				skusInCart:   map[string]int{},
				unitPriceMap: exampleUnitPriceMap,
				deals:        []Deal{},
			},
			0,
			nil,
		},
		{
			"No deals apply",
			args{
				skusInCart: map[string]int{
					"A": 2,
					"B": 1,
					"C": 3,
					"D": 2,
				},
				unitPriceMap: exampleUnitPriceMap,
				deals:        exampleDeals,
			},
			2*50 + 30 + 3*20 + 2*15,
			nil,
		},
		{
			"Deals from example each apply once",
			args{
				skusInCart: map[string]int{
					"A": 3,
					"B": 2,
					"C": 1,
					"D": 1,
				},
				unitPriceMap: exampleUnitPriceMap,
				deals:        exampleDeals,
			},
			130 + 45 + 20 + 15,
			nil,
		},
		{
			"Deals from example apply multiple times",
			args{
				skusInCart: map[string]int{
					"A": 6,
					"B": 8,
					"C": 1,
				},
				unitPriceMap: exampleUnitPriceMap,
				deals:        exampleDeals,
			},
			2*130 + 4*45 + 20,
			nil,
		},
		{
			"Error when SKU does not exist",
			args{
				skusInCart: map[string]int{
					"E": 5,
				},
				unitPriceMap: exampleUnitPriceMap,
				deals:        exampleDeals,
			},
			0,
			fmt.Errorf("SKU E does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CalculateTotal(tt.args.skusInCart, tt.args.unitPriceMap, tt.args.deals); got != tt.want {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
