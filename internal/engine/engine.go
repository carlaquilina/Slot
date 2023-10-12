package engine

import (
	"fmt"
	"slotengine/internal/paytable"
	"slotengine/internal/reel"
)

type GameEngine interface {
	Play(bet float64) (results [3][4]string, winAmount float64, err error)
}

type BasicGameEngine struct {
	reels    []reel.Reel
	payTable paytable.PayTable
}

func NewBasicGameEngine(reels []reel.Reel, payTable paytable.PayTable) *BasicGameEngine {
	return &BasicGameEngine{reels: reels, payTable: payTable}
}

func (ge BasicGameEngine) Play(bet float64) ([3][4]string, float64, error) {
	if bet <= 0 {
		return [3][4]string{}, 0, fmt.Errorf("bet must be greater than 0")
	}
	var results [3][4]string

	for j, reel := range ge.reels {
		spinResult, err := reel.Spin()
		if err != nil {
			return results, 0, err
		}

		for i, symbol := range spinResult {
			results[i][j] = symbol
		}
	}

	winAmount, err := ge.payTable.CalculatePayout(results, bet)
	if err != nil {
		return results, 0, err
	}

	return results, winAmount, nil
}
