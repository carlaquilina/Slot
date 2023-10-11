package reel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpin(t *testing.T) {
	tests := []struct {
		name           string
		symbols        []string
		expectedError  error
		expectedResult []string
	}{
		{
			name:    "StandardReel",
			symbols: []string{"A", "A", "K", "10", "K", "J"},
		},
		{
			name:    "ReelWithJoker",
			symbols: []string{"Q", "Q", "J", "10", "X", "X"},
		},
		{
			name:           "EmptyReel",
			symbols:        []string{},
			expectedError:  fmt.Errorf("reel must have at least 3 symbols"),
			expectedResult: []string{},
		},
		{
			name:           "SingleSymbolReel",
			symbols:        []string{"A"},
			expectedError:  fmt.Errorf("reel must have at least 3 symbols"),
			expectedResult: []string{"A"},
		},
		{
			name:           "TwoSymbolsReel",
			symbols:        []string{"A", "B"},
			expectedError:  fmt.Errorf("reel must have at least 3 symbols"),
			expectedResult: []string{"A", "B"},
		},
		{
			name:    "VeryLongReel",
			symbols: makeLongReel(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBasicReel(tt.symbols)
			result, err := r.Spin()

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
				assert.Equal(t, tt.expectedResult, result)
				return // Exit the current sub-test early
			}

			assert.NoError(t, err)

			// Check if all symbols in the result are valid
			for _, symbol := range result {
				if !contains(tt.symbols, symbol) {
					t.Errorf("Unexpected symbol %s in result", symbol)
				}
			}

			assert.True(t, isConsecutive(tt.symbols, result), "Result symbols are not consecutive in original symbols")
		})
	}
}

// Utility function to generate a long reel for testing
func makeLongReel() []string {
	reel := []string{}
	for i := 0; i < 100; i++ {
		reel = append(reel, fmt.Sprintf("A%d", i))
	}
	return reel
}

// Utility function to check if a slice contains an item
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Utility function to check if the result is a sub-slice of the original symbols
func isConsecutive(symbols []string, result []string) bool {
	for i := 0; i <= len(symbols)-len(result); i++ {
		if equal(symbols[i:i+len(result)], result) {
			return true
		}
	}
	return false
}

// Utility function to check if two slices are equal
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
