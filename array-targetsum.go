// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 7, 11, 15, 20}
	targetSum := 26

	seen := make(map[int]int)

	for i, val := range a {
		complement := targetSum - val
		if idex, exists := seen[complement]; exists {
			fmt.Print("Found", idex, i)
		}

		seen[val] = i

	}

}
