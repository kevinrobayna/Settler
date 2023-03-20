package settler

import "math"

func debit(v float64) float64 {
	return -1 * math.Abs(v)
}

func credit(v float64) float64 {
	return math.Abs(v)
}

// roundUp rounds a float64 to 2 decimal places.
func roundUp(value float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(value*ratio) / ratio
}

// isOddSplit returns true if the remainder of val/n is not zero.
// This means that the split is not even and therefore someone needs to pay a cent more.
func isOddSplit(val float64, n int) bool {
	r := math.Remainder(val, float64(n))

	return r != 0
}
