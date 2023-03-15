package settler

import (
	"math"
)

type Debt map[string]float64

type Transaction struct {
	PayerID  string
	Amount   float64
	PayeeIDs []string
}

// CalculateDebt calculates what each person owes each other spreading evenly.
// If the split is not even, the payer pays the extra cent as this simplifies the calculus.
// Currently, there's no support for different weights i.e. Someone paid 100 but instead of paying evenly they want to pay what they owe exactly.
func CalculateDebt(transactions []Transaction) Debt {
	d := Debt{}
	for _, t := range transactions {
		d[t.PayerID] += t.Amount
		v := roundUp(t.Amount / float64(len(t.PayeeIDs))) // 100/3 = 33.33
		for _, payeeID := range t.PayeeIDs {
			d[payeeID] -= v
			if isOddSplit(t.Amount, len(t.PayeeIDs)) && payeeID == t.PayerID {
				// When we split an amount sometimes the split is a rational number i.e. 33.33333333333
				// In this case since currencies are not rational, and we only have 2 decimal places.
				// Someone needs to pay a cent more, usually the payer as this simplifies things.
				d[payeeID] -= 0.01
			}
		}
	}

	return d
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
