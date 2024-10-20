//Hash Map Method: More flexible and works with a broader range of characters.
//Time Complexity: O(n)
//Space Complexity: O(k) (where k is the number of unique characters)

package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	// Check if the lengths of the strings are different
	if len(s) != len(t) {
		return false
	}

	// Create a map to count character occurrences
	counts := make(map[rune]int)

	// Count characters in the first string
	for _, char := range s {
		counts[char]++ // Increment count for this character
	}

	// Decrease counts based on the second string
	for _, char := range t {
		counts[char]-- // Decrement count for this character
		// If count drops below zero, t has more of this character
		if counts[char] < 0 {
			return false
		}
	}

	// If all counts are zero, they are anagrams
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
