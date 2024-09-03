package main

import (
    "fmt"
    "strconv"
)

// reverseInt reverses an integer and takes into account if it is negative
func reverseInt(n int) int {
    // Check if the integer is negative
    isNegative := n < 0

    // Convert the integer to its absolute value
    if isNegative {
        n = -n
    }

    // Convert the integer to a string
    str := strconv.Itoa(n)

    // Reverse the string
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    // Convert the reversed string back to an integer
    reversedStr := string(runes)
    reversedInt, err := strconv.Atoi(reversedStr)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return 0
    }

    // If the original integer was negative, make the reversed integer negative
    if isNegative {
        reversedInt = -reversedInt
    }

    return reversedInt
}

func main() {
    // Test cases
    testCases := []int{12345, -12345, 100, -100, 0}

    for _, n := range testCases {
        fmt.Printf("Original: %d, Reversed: %d\n", n, reverseInt(n))
    }
}