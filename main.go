package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random number between min and max (inclusive)
func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Function to get the player's name
func getName() string {
	var name string
	fmt.Println("Welcome to Casino..")
	fmt.Println("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play! \n", name)
	return name
}

// Function to get the player's bet
func getBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance= %d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("Bet cannot be larger than balance")
		} else {
			break
		}
	}
	return bet
}

// Function to generate the symbol array based on the symbol counts
func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

// Function to generate a spin result (randomly select symbols)
func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	// Generate a random spin result for each column
	for col := 0; col < cols; col++ {
		colResult := []string{}
		selected := map[int]bool{}

		// For each row in the column, randomly pick a unique symbol
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				if !selected[randomIndex] {
					selected[randomIndex] = true
					colResult = append(colResult, reel[randomIndex])
					break
				}
			}
		}
		result = append(result, colResult)
	}
	return result
}

// Function to display the spin result
func displaySpin(result [][]string) {
	for _, row := range result {
		for _, symbol := range row {
			fmt.Print(symbol + " ")
		}
		fmt.Println()
	}
}

// Function to calculate winnings (for simplicity, we check if all symbols in a row are the same)
func calculateWinnings(result [][]string, bet uint, multipliers map[string]uint) uint {
	winnings := uint(0)
	// Check each row for a win (if all symbols in a row are the same)
	for _, row := range result {
		if len(row) > 0 {
			firstSymbol := row[0]
			win := true
			for _, symbol := range row {
				if symbol != firstSymbol {
					win = false
					break
				}
			}
			if win {
				winnings += bet * multipliers[firstSymbol]
			}
		}
	}
	return winnings
}

func main() {
	// Define symbols and their counts
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}

	// Define multipliers for each symbol
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	// Generate symbol array
	symbolArr := generateSymbolArray(symbols)
	fmt.Println("Symbols array:", symbolArr)

	// Starting balance
	balance := uint(200)

	// Get player's name
	getName()

	// Main game loop
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		// Deduct the bet from the balance
		balance -= bet
		fmt.Println("Spinning...")

		// Get the spin result
		result := getSpin(symbolArr, 3, 3)

		// Display the result of the spin
		displaySpin(result)

		// Calculate winnings
		winnings := calculateWinnings(result, bet, multipliers)

		// Update balance
		balance += winnings
		fmt.Printf("You won: %d\n", winnings)
		fmt.Printf("Your balance is now: %d\n", balance)
	}

	// End of game
	fmt.Printf("You left with %d \n", balance)
}
