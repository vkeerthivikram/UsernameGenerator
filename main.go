package main

import (
	"embed"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

//go:embed eff_large_wordlist.txt
var wordlist embed.FS

// main is the entry point of the application for generating and printing a username based on a random word and number.
func main() {
	// Parse command-line arguments
	numLength := flag.Int("length", 4, "length of the number")
	flag.Parse()

	// Ensure minimum length of 1, default to 4 if 0 or negative
	if *numLength <= 0 {
		*numLength = 4
	}

	// Generate and print the username
	username := generateUsername(*numLength)
	fmt.Println(username)
}

// generateUsername creates a username by combining a random word from a wordlist with a zero-padded random number.
// numLength defines the length of the numeric component in the username.
// Panics if the wordlist cannot be read, is empty, or has invalid formatting.
func generateUsername(numLength int) string {
	wordlistData, err := wordlist.ReadFile("eff_large_wordlist.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(wordlistData)), "\n")
	if len(lines) == 0 {
		panic("empty wordlist")
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	// Select random line and extract word (after tab)
	randomLine := lines[r.Intn(len(lines))]
	parts := strings.Split(randomLine, "\t")
	if len(parts) != 2 {
		panic("invalid wordlist format")
	}
	word := parts[1] // Get the word part

	// Generate random number with proper zero-padding
	maxNum := int(math.Pow10(numLength))
	num := r.Intn(maxNum)
	numStr := fmt.Sprintf("%0*d", numLength, num)

	// Capitalize first letter of word
	if len(word) > 0 {
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		word = string(runes)
	}

	return word + numStr
}
