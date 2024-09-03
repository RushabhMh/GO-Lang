package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Function to count the occurrence of each word in a paragraph
func countWords(paragraph string) map[string]int {
	// Function to remove punctuation and convert to lowercase
	new := ""
	for _, v := range paragraph {
		if unicode.IsLetter(v) || unicode.IsSpace(v) {
			new = new + string(v)

		}
	}
	fmt.Println(new)
	// Split the cleaned paragraph into words
	words := strings.Fields(new)

	// Count the occurrence of each word
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	paragraph := `Hello, world! This is a test. This test is only a test. "Hello" world, hello!`

	wordCount := countWords(paragraph)

	// Print the word counts
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
