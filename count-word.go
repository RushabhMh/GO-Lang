package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(paragraph string) {
	// Define a regex to remove all non-alphabetical characters
	// \s is for whitespaces
	reg, _ := regexp.Compile("[^a-zA-Z\\s]+")

	// Replace non-alphabetical characters with an empty string and convert to lowercase
	cleanedParagraph := reg.ReplaceAllString(paragraph, "")
	cleanedParagraph = strings.ToLower(cleanedParagraph)

	// Split the cleaned paragraph into words
	words := strings.Fields(cleanedParagraph)

	// Create a map to store word counts
	wordCount := make(map[string]int)

	// Count the occurrences of each word
	for _, word := range words {
		wordCount[word]++
	}

	// Print the word counts
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	paragraph := "Hello, world! It's a wonderful world. Isn't it?"
	countWords(paragraph)
}
