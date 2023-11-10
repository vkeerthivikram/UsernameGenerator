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

func main() {
	// Parse command-line arguments
	numLength := flag.Int("length", 4, "length of the number")
	flag.Parse()
	var zero int = 0
	var four int = 4
	if numLength == nil || numLength == &zero {
		numLength = &four
	}
	// Generate and print the username
	username := generateUsername(*numLength)
	fmt.Println(username)
}

func generateUsername(numLength int) string {
	wordlistData, err := wordlist.ReadFile("eff_large_wordlist.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(wordlistData), "\n")
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	// Select a random word
	word := words[r.Intn(len(words))]

	// Generate a random number with the specified length
	num := r.Intn(int(math.Pow10(numLength)))

	// Concatenate the word and the number to form the username
	username := word + strconv.Itoa(num)
	removeIndex := strings.Split(username, "	")
	finalUsername := removeIndex[len(removeIndex)-1]
	convertToRune := []rune(finalUsername)
	convertToRune[0] = unicode.ToUpper(convertToRune[0])
	usernameStringTitleCase := string(convertToRune)
	return usernameStringTitleCase
}