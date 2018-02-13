package product

import (
	"testing"
	"math"
)

func TestProduct(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"the first is zero", 0, 42, 0},
		{"the second is zero", 55, 0, 0},
		{"both are positive", 100, 42, 4200},
		{"both are negative", -10, -123456, 1234560},
		{"the first is positive, the second is negative", 1, -789, -789},
		{"the first is negative, the second is positive", -2, 21, -42},
		{"the first is too large", math.MaxInt64, 2, 0},
		{"the second is too large", 99999, math.MaxInt64, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Product(tt.a, tt.b); got != tt.want {
				t.Errorf("Name: %s, Product() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
