package settler

type Debt map[string]float64

// Transaction represents a "bill" in which someone paid an amount for a list of people.
type Transaction struct {
	PayerID string
	Amount  float64
	Shares  []Share
}

// Share represents the portion of the bill that a person owes. The bill is split evenly unless the amount is specified.
type Share struct {
	PayeeID string
}
