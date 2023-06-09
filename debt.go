package settler

// CalculateDebt calculates what each person owes each other spreading evenly.
// If the split is not even, the payer pays the extra cent as this simplifies the calculus.
// Currently, there's no support for different weights i.e. Someone paid 100 but instead of paying evenly they want to pay what they owe exactly.
func CalculateDebt(transactions []Transaction) Debt {
	d := Debt{}
	for _, t := range transactions {
		d[t.PayerID] += debit(t.Amount) // Debit the payer the total amount
		v := roundUp(t.Amount / float64(len(t.Shares)))
		for _, share := range t.Shares {
			d[share.PayeeID] += credit(v) // credit the payee their share
			if isOddSplit(t.Amount, len(t.Shares)) && share.PayeeID == t.PayerID {
				// When we split an amount sometimes the split is a rational number i.e. 33.33333333333
				// In this case since currencies are not rational, and we only have 2 decimal places.
				// Someone needs to pay a cent more, usually the payer as this simplifies things.
				d[t.PayerID] += credit(0.01)
			}
		}
	}

	return d
}
