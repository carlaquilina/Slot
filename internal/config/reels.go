package config

import "slotengine/internal/reel"

var reel1 = []string{"A", "A", "K", "10", "K", "J"}
var reel2 = []string{"Q", "Q", "J", "10", "X", "X"}
var reel3 = []string{"A", "K", "Q", "10", "J", "X"}
var reel4 = []string{"A", "K", "Q", "10", "J", "J"}

var Reels = []reel.Reel{
	reel.NewBasicReel(reel1),
	reel.NewBasicReel(reel2),
	reel.NewBasicReel(reel3),
	reel.NewBasicReel(reel4),
}
