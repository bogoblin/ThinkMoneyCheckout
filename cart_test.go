package ThinkMoneyCheckout

import "testing"

func TestCalculateTotal(t *testing.T) {
	type args struct {
		skusInCart map[SKU]int
		deals      []Deal
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTotal(tt.args.skusInCart, tt.args.deals); got != tt.want {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
