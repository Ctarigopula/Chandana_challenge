package main

import (
	"testing"
)

func Test_valid(t *testing.T) {
	type creditcard struct {
		card string
	}
	tests := []struct {
		name       string
		creditcard creditcard
		wantOutput bool
	}{
		{
			name: "valid",
			creditcard: creditcard{
				card: "4123456789123456",
			},
			wantOutput: true,
		},
		{
			name: "in-valid-length",
			creditcard: creditcard{
				card: "4123456",
			},
			wantOutput: false,
		},
		{
			name: "non-digits",
			creditcard: creditcard{
				card: "Q1456 - 89123456",
			},
			wantOutput: false,
		},
		{
			name: "empty",
			creditcard: creditcard{
				card: "",
			},
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := valid(tt.creditcard.card); gotOutput != tt.wantOutput {
				t.Errorf("valid() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
