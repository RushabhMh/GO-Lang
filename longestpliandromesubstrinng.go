package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	n := len(s)
	// Create a new string with inserted '#' to handle even-length palindromes
	newStr := make([]rune, 2*n+1)
	i := 0
	newStr[i] = '#'
	i++
	for _, c := range s {
		newStr[i] = c
		i++
		newStr[i] = '#'
		i++
	}

	// Array to store the length of the palindrome centered at each position
	p := make([]int, 2*n+1)
	center, right := 0, 0
	longestLength, longestCenter := 0, 0

	// Loop through the newStr array to find the longest palindromic substring
	for i := 0; i < len(newStr); i++ {
		mirror := 2*center - i

		if right > i {
			p[i] = min(p[mirror], right-i)
		}

		// Expand around the center
		a := i + (p[i] + 1)
		b := i - (p[i] + 1)
		for b >= 0 && a < len(newStr) && newStr[a] == newStr[b] {
			p[i]++
			a++
			b--
		}

		// Update the longest palindrome found
		if p[i] >= longestLength {
			longestCenter = i
			longestLength = p[i]
		}

		// Update the right boundary of the current longest palindrome
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	// Extract the palindrome from the modified string
	start := longestCenter - longestLength
	end := longestCenter + longestLength
	palindrome := string(newStr[start:end])
	// Remove '#' from the extracted palindrome
	return strings.ReplaceAll(palindrome, "#", "")
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example usage
	s := "babad"
	fmt.Println("Longest Palindromic Substring:", longestPalindrome(s))
}
