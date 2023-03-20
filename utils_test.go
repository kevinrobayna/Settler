package settler

import (
	"testing"
)

func Test_credit(t *testing.T) {
	tests := []struct {
		name  string
		given float64
		want  float64
	}{
		{
			name:  "positive value, sign remains the same",
			given: 100,
			want:  100,
		},
		{
			name:  "negative value, value returned is positive",
			given: -100,
			want:  100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := credit(tt.given); got != tt.want {
				t.Errorf("credit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_debit(t *testing.T) {
	tests := []struct {
		name  string
		given float64
		want  float64
	}{
		{
			name:  "positive value, value returned is negative",
			given: 100,
			want:  -100,
		},
		{
			name:  "negative value, value returned is negative",
			given: -100,
			want:  -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := debit(tt.given); got != tt.want {
				t.Errorf("debit() = %v, want %v", got, tt.want)
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
