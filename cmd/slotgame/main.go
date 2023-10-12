package main

import (
	"fmt"
	"slotengine/internal/config"
	"slotengine/internal/engine"
	"slotengine/internal/paytable"
	"strings"
)

func main() {

	gameEngine := engine.NewBasicGameEngine(config.Reels, paytable.NewBasicPayTable(config.Multipliers, config.Patterns))

	fmt.Println("Welcome to the Slot Game!")

	// Main game loop
	for {
		var bet float64
		fmt.Print("Enter your betting amount: ")
		_, err := fmt.Scan(&bet)
		if err != nil {
			fmt.Println("Invalid bet amount. Please enter a numeric value.")
			continue
		}

		resultWindows, winAmount, err := gameEngine.Play(bet)
		if err != nil {
			fmt.Printf("Error gameEngine.Play: %v\n", err)
			continue
		}

		for _, window := range resultWindows {
			fmt.Println(window)
		}
		fmt.Printf("You won: %v\n", winAmount)

		var choice string
		fmt.Print("Do you want to play again? (yes/no): ")
		fmt.Scan(&choice)
		choice = strings.TrimSpace(strings.ToLower(choice))
		if choice == "yes" || choice == "y" {
			continue
		} else {
			fmt.Println("Thank you for playing!")
			break
		}
	}
}
