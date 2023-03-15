package settler

import "testing"

func TestCalculateDebt(t *testing.T) {
	tests := []struct {
		name         string
		transactions []Transaction
		result       Debt
	}{
		{
			name:         "no transactions, expects no debts",
			transactions: []Transaction{},
			result:       Debt{},
		},
		{
			name: "one payer, another payee, payee owes total value",
			transactions: []Transaction{
				{Amount: 100, PayerID: "1", Shares: []Share{{PayeeID: "2"}}},
			},
			result: Debt{
				"1": 100,
				"2": -100,
			},
		},
		{
			name: "one payer, two payees, debt is split evenly",
			transactions: []Transaction{
				{Amount: 100, PayerID: "1", Shares: []Share{{PayeeID: "1"}, {PayeeID: "2"}}},
			},
			result: Debt{
				"1": 50,
				"2": -50,
			},
		},
		{
			name: "one payer, three payees, debt is split evenly except one has to pay a cent more",
			transactions: []Transaction{
				{Amount: 100, PayerID: "1", Shares: []Share{{PayeeID: "1"}, {PayeeID: "2"}, {PayeeID: "3"}}},
			},
			result: Debt{
				"1": 66.66,
				"2": -33.33,
				"3": -33.33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			debt := CalculateDebt(tt.transactions)

			if len(debt) != len(tt.result) {
				t.Errorf("Number of elements in debt and expected debt should match. Got %v, want %v", len(debt), len(tt.result))
			}

			totalDebt := 0.0
			for id, v := range tt.result {
				if v != debt[id] {
					t.Errorf("Debt for ID %v should match the expected. Got %v, want %v", id, debt[id], v)
				}
				totalDebt += debt[id]
			}

			// If everyone paid their debt, the total debt should be zero.
			if totalDebt != 0.0 {
				t.Errorf("Total debt should be zero. Got %v", totalDebt)
			}
		})
	}
}

func Test_isOddSplit(t *testing.T) {
	tests := []struct {
		name           string
		value          float64
		numberOfPayees int
		want           bool
	}{
		{
			name:           "100 / 3 = 33.333333, is odd split",
			value:          100,
			numberOfPayees: 3,
			want:           true,
		},
		{
			name:           "100 / 6 = 33.333333, is odd split",
			value:          100,
			numberOfPayees: 6,
			want:           true,
		},
		{
			name:           "100 / 2 = 50, no odd split",
			value:          100,
			numberOfPayees: 2,
			want:           false,
		},
		{
			name:           "100 / 4 = 25, no odd split",
			value:          100,
			numberOfPayees: 4,
			want:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOddSplit(tt.value, tt.numberOfPayees)
			if tt.want != result {
				t.Errorf("isOddSplit(%v, %v) = %v, want %v", tt.value, tt.numberOfPayees, result, tt.want)
			}
		})
	}
}

func Test_roundUp(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "round 1.234 to 1.23",
			value: 1.234,
			want:  1.23,
		},
		{
			name:  "round 1.236 to 1.24",
			value: 1.236,
			want:  1.24,
		},
		{
			name:  "round 33.333333 to 33.33",
			value: 33.333333333,
			want:  33.33,
		},
		{
			name:  "round 25.00000 to 25.00",
			value: 25.0000000,
			want:  25.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundUp(tt.value)
			if tt.want != result {
				t.Errorf("roundUp(%v) = %v, want %v", tt.value, result, tt.want)
			}
		})
	}
}
