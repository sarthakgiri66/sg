package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Convert strings to slices of runes, sort, and compare
	sliceS := []rune(s)
	sliceT := []rune(t)
	sort.Slice(sliceS, func(i, j int) bool { return sliceS[i] < sliceS[j] })
	sort.Slice(sliceT, func(i, j int) bool { return sliceT[i] < sliceT[j] })

	return string(sliceS) == string(sliceT)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}


//Sorting Method: Good for small strings or when simplicity is preferred.
//Time Complexity: O(n log n)
//Space Complexity: O(n) (for sorted strings)
