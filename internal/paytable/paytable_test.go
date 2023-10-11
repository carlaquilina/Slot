package paytable

import (
	"fmt"
	"slotengine/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePayout(t *testing.T) {
	tests := []struct {
		name           string
		multipliers    map[string]float64
		patterns       [][]config.Coordinate
		matrix         [3][4]string
		bet            float64
		expectedPayout float64
		expectedError  error
	}{
		{
			name: "MiddleRowWiCol:n",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
			},
			patterns: [][]config.Coordinate{
				// Middle horizontal line
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "A", "J", "10"},
				{"A", "A", "A", "A"},
				{"Q", "Q", "J", "J"},
			},
			bet:            5.0,
			expectedPayout: 100.0, // A pays 20x, so 5*20 = 100,
			expectedError:  nil,
		},
		{
			name: "JokerSubstituteWin",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
				"X":  0, // Joker
			},
			patterns: [][]config.Coordinate{
				// Middle horizontal line
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "Q", "10", "J"},
				{"A", "A", "X", "A"},
				{"Q", "K", "J", "A"},
			},
			bet:            5.0,
			expectedPayout: 100.0, // 2 As and a Joker, so 5*20 = 100
			expectedError:  nil,
		},
		{
			name: "MultiplePatternsWithJoker",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
				"X":  0, // Joker
			},
			patterns: [][]config.Coordinate{
				// Middle & Top horizontal line
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
				{
					{Row: 0, Col: 0},
					{Row: 0, Col: 1},
					{Row: 0, Col: 2},
					{Row: 0, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "K", "K", "K"},
				{"A", "A", "X", "A"},
				{"Q", "J", "10", "J"},
			},
			bet:            5.0,
			expectedPayout: 175.0, // 2 As and a Joker for 5*20=100, and 4 Ks for 5*15=75, Total: 175
			expectedError:  nil,
		},
		{
			name: "NoJokerWin",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
				"X":  0, // Joker
			},
			patterns: [][]config.Coordinate{
				// Middle horizontal line
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "A", "J", "Q"},
				{"A", "J", "J", "A"},
				{"Q", "J", "K", "J"},
			},
			bet:            5.0,
			expectedPayout: 0, // No matching pattern
			expectedError:  nil,
		},
		{
			name: "JokerAsMultiplier",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
				"X":  0, // Joker
			},
			patterns: [][]config.Coordinate{
				// Middle horizontal line
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"J", "Q", "10", "K"},
				{"A", "X", "X", "A"},
				{"Q", "J", "10", "K"},
			},
			bet:            5.0,
			expectedPayout: 100.0, // 2 As and 2 Jokers, so 5*20 = 100
			expectedError:  nil,
		},
		{
			name: "BottomRowDifferentSymbolCol:s",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
				"X":  0, // Joker
			},
			patterns: [][]config.Coordinate{
				// Bottom horizontal line
				{
					{Row: 2, Col: 0},
					{Row: 2, Col: 1},
					{Row: 2, Col: 2},
					{Row: 2, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"J", "A", "Q", "J"},
				{"K", "10", "A", "Q"},
				{"A", "J", "10", "K"},
			},
			bet:            5.0,
			expectedPayout: 0, // No matching pattern
			expectedError:  nil,
		},
		{
			name: "InvalidSymbol",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
			},
			patterns: [][]config.Coordinate{
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"Z", "A", "K", "Q"},
				{"A", "K", "Q", "Z"},
				{"K", "Q", "Z", "A"},
			},
			bet:            5.0,
			expectedPayout: 0.0,
			expectedError:  nil,
		},
		{
			name: "DifferentMultipliers",
			multipliers: map[string]float64{
				"A":  25, // Changed
				"K":  10, // Changed
				"Q":  8,  // Changed
				"J":  4,  // Changed
				"10": 1,  // Changed
			},
			patterns: [][]config.Coordinate{
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "A", "J", "10"},
				{"A", "A", "A", "A"},
				{"Q", "Q", "J", "J"},
			},
			bet:            5.0,
			expectedPayout: 125.0, // A pays 25x, so 5*25 = 125
			expectedError:  nil,
		},
		{
			name: "DiagonalPatternWin",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
			},
			patterns: [][]config.Coordinate{
				// Diagonal from top-left to bottom-right
				{
					{Row: 0, Col: 0},
					{Row: 1, Col: 1},
					{Row: 2, Col: 2},
					{Row: 2, Col: 3}, // This will make the pattern invaCol:lid
				},
			},
			matrix: [3][4]string{
				{"A", "K", "J", "10"},
				{"K", "A", "J", "10"},
				{"K", "Q", "A", "J"},
			},
			bet:            5.0,
			expectedPayout: 0.0,
			expectedError:  nil,
		},
		{
			name: "NegativeBet",
			multipliers: map[string]float64{
				"A":  20,
				"K":  15,
				"Q":  10,
				"J":  5,
				"10": 2,
			},
			patterns: [][]config.Coordinate{
				{
					{Row: 1, Col: 0},
					{Row: 1, Col: 1},
					{Row: 1, Col: 2},
					{Row: 1, Col: 3},
				},
			},
			matrix: [3][4]string{
				{"K", "A", "J", "10"},
				{"A", "A", "A", "A"},
				{"Q", "Q", "J", "J"},
			},
			bet:            -5.0,
			expectedPayout: 0.0,
			expectedError:  fmt.Errorf("bet must be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := NewBasicPayTable(tt.multipliers, tt.patterns)

			payout, err := pt.CalculatePayout(tt.matrix, tt.bet)
			assert.Equal(t, tt.expectedPayout, payout, "The payout does not match the expected value")
			assert.Equal(t, tt.expectedError, err, "The error does not match the expected value")
		})
	}
}
