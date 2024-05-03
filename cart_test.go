package ThinkMoneyCheckout

import "testing"

func TestCalculateTotal(t *testing.T) {
	type args struct {
		skusInCart map[SKU]int
		deals      []Deal
	}
	exampleDeals := []Deal{
		{
			skus: map[string]int{
				"A": 3,
			},
			price: 130,
		},
		{
			skus: map[string]int{
				"B": 2,
			},
			price: 45,
		},
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Returns 0 if no items in cart",
			args{
				skusInCart: map[SKU]int{},
				deals:      []Deal{},
			},
			0,
		},
		{
			"No deals apply",
			args{
				skusInCart: map[SKU]int{
					{name: "A", price: 50}: 2,
					{name: "B", price: 30}: 1,
					{name: "C", price: 20}: 3,
					{name: "D", price: 15}: 2,
				},
				deals: exampleDeals,
			},
			2*50 + 30 + 3*20 + 2*15,
		},
		{
			"Deals from example each apply once",
			args{
				skusInCart: map[SKU]int{
					{name: "A", price: 50}: 3,
					{name: "B", price: 30}: 2,
					{name: "C", price: 20}: 1,
					{name: "D", price: 15}: 1,
				},
				deals: exampleDeals,
			},
			130 + 45 + 20 + 15,
		},
		{
			"Deals from example apply multiple times",
			args{
				skusInCart: map[SKU]int{
					{name: "A", price: 50}: 6,
					{name: "B", price: 30}: 8,
					{name: "C", price: 20}: 1,
				},
				deals: exampleDeals,
			},
			2*130 + 4*45 + 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTotal(tt.args.skusInCart, tt.args.deals); got != tt.want {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
