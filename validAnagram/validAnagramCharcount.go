package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	counts := make([]int, 26) // Assuming only lowercase English letters
	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}


//Character Count Method: Best for scenarios where you can restrict to a known character set (e.g., lowercase letters).
//Time Complexity: O(n)
//Space Complexity: O(1) (fixed size for 26 lowercase letters)
