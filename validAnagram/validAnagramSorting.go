//Sorting Method: Good for small strings or when simplicity is preferred.
//Time Complexity: O(n log n)
//Space Complexity: O(n) (for sorted strings)





package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	// Step 1: Check if the lengths of the strings are different
	if len(s) != len(t) {
		return false
	}

	// Step 2: Convert strings to slices of runes (to handle Unicode)
	sliceS := []rune(s)
	sliceT := []rune(t)

	// Step 3: Sort the slices
	sort.Slice(sliceS, func(i, j int) bool {
		return sliceS[i] < sliceS[j]
	})
	sort.Slice(sliceT, func(i, j int) bool {
		return sliceT[i] < sliceT[j]
	})

	// Step 4: Compare the sorted slices
	return string(sliceS) == string(sliceT)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
