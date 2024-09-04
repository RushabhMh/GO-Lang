package main

import (
	"fmt"
)

// Function to find the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0
	startindex := 0
	endindex := 0
	for i, char := range s {
		if pos, found := charIndexMap[char]; found && pos >= start {
			start = pos + 1
		}
		charIndexMap[char] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			startindex = start
			endindex = i
		}
	}
	fmt.Println("Arrrayyyyy map rune int")
	fmt.Println(charIndexMap)
	if maxLength != 0 {
		fmt.Println(s[startindex : endindex+1])
	}
	return maxLength
}

func main() {
	// Test cases
	testStrings := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"au",
		"dvdf",
	}

	for _, s := range testStrings {
		fmt.Printf("Input: %s, Length of Longest Substring Without Repeating Characters: %d\n", s, lengthOfLongestSubstring(s))
	}
}
