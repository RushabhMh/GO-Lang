package main

import (
	"fmt"
	"sort"
)

// Check for duplicate values and return a map with counts of each integer
func findDuplicates(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	return counts
}

// Find the count of each duplicate integer
func countDuplicates(counts map[int]int) map[int]int {
	duplicates := make(map[int]int)
	for num, count := range counts {
		if count > 1 {
			duplicates[num] = count
		}
	}
	return duplicates
}

// Return the top 'k' duplicate integers
func topKDuplicates(duplicates map[int]int, k int) []int {
	type kv struct {
		Key   int
		Value int
	}

	var sorted []kv
	for k, v := range duplicates {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	var topK []int
	for i := 0; i < k && i < len(sorted); i++ {
		topK = append(topK, sorted[i].Key)
	}

	return topK
}

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 6, 3, 3, 7, 8, 9, 1, 1, 1}

	// Step 1: Check for duplicate values and get counts
	counts := findDuplicates(nums)
	fmt.Println("Counts:", counts)

	// Step 2: Find the count of each duplicate integer
	duplicates := countDuplicates(counts)
	fmt.Println("Duplicates:", duplicates)

	// Step 3: For a given value 'k', return top 'k' duplicate integers
	k := 3
	topK := topKDuplicates(duplicates, k)
	fmt.Printf("Top %d duplicates: %v\n", k, topK)
}
