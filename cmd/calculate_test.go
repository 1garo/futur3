package cmd

import (
	"testing"
)

func TestCalculate(t *testing.T) {
//	content := `
//years: 1
//initial_investment: 10000
//tax_percentage: 12
//monthly: 0
//is_interest_compound: true
//`
	tests := []struct {
		expected error
		content string
	}{
		{
			expected: nil,
			content: `
years: 1
initial_investment: 10000
tax_percentage: 12
monthly: 0
is_interest_compound: true
`,
		},
		{
			expected: nil,
			content: `
years: 1
initial_investment: 10000
tax_percentage: 12
monthly: 0
is_interest_compound: false
`,
		},
	}
	for _, tt := range tests {
		err := CalculateInterest([]byte(tt.content))
		if err != nil {
			t.Errorf("expected: %v, found: %v", tt.expected, err)
		}
	}

}
