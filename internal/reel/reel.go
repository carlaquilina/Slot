package reel

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate moq -out ../mocks/reel_mock.go -pkg mocks . Reel
type Reel interface {
	Spin() ([]string, error)
}

type BasicReel struct {
	symbols []string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewBasicReel(symbols []string) *BasicReel {
	return &BasicReel{symbols: symbols}
}

func (r BasicReel) Spin() ([]string, error) {
	if len(r.symbols) < 3 {
		return r.symbols, fmt.Errorf("reel must have at least 3 symbols")
	}
	startIndex := rand.Intn(len(r.symbols) - 2) // Ensure we have space for 3 symbols
	return r.symbols[startIndex : startIndex+3], nil
}
