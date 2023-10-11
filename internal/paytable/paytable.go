package paytable

import (
	"fmt"
	"slotengine/internal/config"
)

//go:generate moq -out ../mocks/paytable_mock.go -pkg mocks . PayTable
type PayTable interface {
	CalculatePayout(matrix [3][4]string, bet float64) (float64, error)
}

type BasicPayTable struct {
	multipliers map[string]float64
	patterns    [][]config.Coordinate
}

func NewBasicPayTable(multipliers map[string]float64, patterns [][]config.Coordinate) *BasicPayTable {
	return &BasicPayTable{multipliers: multipliers, patterns: patterns}
}

func (pt BasicPayTable) CalculatePayout(matrix [3][4]string, bet float64) (float64, error) {
	if bet <= 0 {
		return 0, fmt.Errorf("bet must be greater than 0")
	}
	totalWin := 0.0

	for _, pattern := range pt.patterns {
		symbols := []string{}

		for _, coordinate := range pattern {
			symbols = append(symbols, matrix[coordinate.Row][coordinate.Col])
		}

		firstSymbol := symbols[0]
		match := true
		for _, sym := range symbols {
			if sym != "X" && sym != firstSymbol {
				match = false
				break
			}
		}

		if match {
			totalWin += pt.multipliers[firstSymbol] * bet
		}
	}

	return totalWin, nil
}
