package main

import (
    "fmt"
    "unicode/utf8"
)

// reverseString reverses a string and correctly handles Unicode characters
func reverseString(s string) string {
    // Convert the string to a slice of runes to handle Unicode characters
    runes := []rune(s)
    n := len(runes)

    // Reverse the slice of runes
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }

    // Convert the slice of runes back to a string
    return string(runes)
}

func main() {
    // Test cases
    strings := []string{
        "日本語",
        "álphãbét",
        "Hello, World!",
        "こんにちは",
        "¡Hola, mundo!",
    }

    for _, s := range strings {
        fmt.Printf("Original: %s, Reversed: %s\n", s, reverseString(s))
    }
}