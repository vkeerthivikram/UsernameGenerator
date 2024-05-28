package main

import (
	"embed"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:embed eff_large_wordlist.txt
var wordlist embed.FS

// main function is the entry point of the program. It parses the command-line arguments,
// sets the default length of the number to 4 if no length is provided, generates a username
// using the generateUsername function, and prints the generated username.
func main() {
	// Parse command-line arguments
	numLength := flag.Int("length", 4, "length of the number")
	flag.Parse()
	var zero = 0
	var four = 4
	if numLength == nil || numLength == &zero {
		numLength = &four
	}
	// Generate and print the username
	username := generateUsername(*numLength)
	fmt.Println(username)
}

// generateUsername generates a username by selecting a random word from the wordlist and appending a random number of specified length.
// The username is then converted to title case.
//
// Parameters:
//   - numLength (int): The length of the random number to append to the word.
//
// Returns:
//   - (string): The generated username in title case.
func generateUsername(numLength int) string {
	// Read the contents of the wordlist file.
	wordlistData, err := wordlist.ReadFile("eff_large_wordlist.txt")
	if err != nil {
		panic(err)
	}

	// Split the wordlist data into individual words.
	words := strings.Split(string(wordlistData), "\n")

	// Generate a random seed for the random number generator.
	seed := time.Now().UnixNano()

	// Create a new random number generator using the seed.
	r := rand.New(rand.NewSource(seed))

	// Select a random word from the wordlist.
	word := words[r.Intn(len(words))]

	// Generate a random number with the specified length.
	num := r.Intn(int(math.Pow10(numLength)))

	// Concatenate the word and the number to form the username.
	username := word + strconv.Itoa(num)

	// Split the username by the "	" delimiter.
	removeIndex := strings.Split(username, "	")

	// Get the last element of the split username.
	finalUsername := removeIndex[len(removeIndex)-1]

	// Convert the final username to a slice of runes.
	convertToRune := []rune(finalUsername)

	// Convert the first rune of the username to uppercase.
	convertToRune[0] = unicode.ToUpper(convertToRune[0])

	// Convert the slice of runes back to a string and return it.
	usernameStringTitleCase := string(convertToRune)
	return usernameStringTitleCase
}
