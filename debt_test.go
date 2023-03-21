package settler

import (
	"reflect"
	"testing"
)

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
				"1": -100,
				"2": 100,
			},
		},
		{
			name: "one payer, two payees, debt is split evenly",
			transactions: []Transaction{
				{Amount: 100, PayerID: "1", Shares: []Share{{PayeeID: "1"}, {PayeeID: "2"}}},
			},
			result: Debt{
				"1": -50,
				"2": 50,
			},
		},
		{
			name: "one payer, three payees, debt is split evenly except one has to pay a cent more",
			transactions: []Transaction{
				{Amount: 100, PayerID: "1", Shares: []Share{{PayeeID: "1"}, {PayeeID: "2"}, {PayeeID: "3"}}},
			},
			result: Debt{
				"1": -66.66,
				"2": 33.33,
				"3": 33.33,
			},
		},
		{
			name: "example from readme",
			transactions: []Transaction{
				{Amount: 100, PayerID: "A", Shares: []Share{{PayeeID: "A"}, {PayeeID: "B"}}},
				{Amount: 60, PayerID: "C", Shares: []Share{{PayeeID: "A"}, {PayeeID: "B"}, {PayeeID: "C"}}},
				{Amount: 300, PayerID: "B", Shares: []Share{{PayeeID: "A"}, {PayeeID: "B"}, {PayeeID: "C"}}},
			},
			result: Debt{
				"A": 70,
				"B": -130,
				"C": 60,
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

func TestSettleDebt(t *testing.T) {
	tests := []struct {
		name string
		debt Debt
		want []Transaction
	}{
		{
			name: "no debt, no transactions",
			debt: Debt{},
			want: []Transaction{},
		},
		{
			name: "one debtor, one creditor, one transaction",
			debt: Debt{
				"1": -100,
				"2": 100,
			},
			want: []Transaction{
				{Amount: 100, PayerID: "2", Shares: []Share{{PayeeID: "1"}}},
			},
		},
		{
			name: "one debtor, two creditors, two transactions",
			debt: Debt{
				"A": 70,
				"B": -130,
				"C": 60,
			},
			want: []Transaction{
				{Amount: 60, PayerID: "B", Shares: []Share{{PayeeID: "C"}}},
				{Amount: 70, PayerID: "B", Shares: []Share{{PayeeID: "A"}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SettleDebt(tt.debt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SettleDebt() = %v, want %v", got, tt.want)
			}
		})
	}
}
