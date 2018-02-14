package product

import "math"

// Product returns product of two integers, without using the '*' operator.
// It doesn't use a recursive call because stack overflow possibility
func Product(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	var (
		index  int
		result int
	)

	min := absMin(a, b) // min(abs(a), abs(b))
	max := absMax(a, b) // max(abs(a), abs(b))

	for index < min {
		result += max
		index++
	}

	// if sign of result is a negative number
	if !sign(a, b) {
		return -result
	}

	return result
}

// sign returns true if both integers are positive or negative, false otherwise
func sign(a int, b int) bool {
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		return true
	}

	return false
}

// absMax returns the maximum of abs(a) and abs(b)
func absMax(a int, b int) int {
	aabs := math.Abs(float64(a))
	babs := math.Abs(float64(b))

	if aabs < babs {
		return int(babs)
	}

	return int(aabs)
}

// absMin returns the minimum of abs(a) and abs(b)
func absMin(a int, b int) int {
	aabs := math.Abs(float64(a))
	babs := math.Abs(float64(b))

	if aabs < babs {
		return int(aabs)
	}

	return int(babs)
}
