//Character Count Method: Best for scenarios where you can restrict to a known character set (e.g., lowercase letters).
//Time Complexity: O(n)
//Space Complexity: O(1) (fixed size for 26 lowercase letters)


package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	// Step 1: Check if the lengths of the strings are different
	if len(s) != len(t) {
		return false
	}

	// Step 2: Create an array to count character occurrences
	counts := make([]int, 26) // For lowercase English letters

	// Step 3: Count characters in the first string
	for _, char := range s {
		counts[char-'a']++ // Increment count for this character
	}

	// Step 4: Decrease counts based on the second string
	for _, char := range t {
		counts[char-'a']-- // Decrement count for this character
		// If count drops below zero, t has more of this character
		if counts[char-'a'] < 0 {
			return false
		}
	}

	// Step 5: If all counts are zero, they are anagrams
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
