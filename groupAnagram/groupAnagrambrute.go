//Brute Force Approach (O(n^2 * k))
//Time Complexity: O(n² * k), where n is the number of strings and k is the maximum length of a string (for the anagram check).
//Space Complexity: O(n) for the result storage.


package main

import (
	"fmt"
)

func groupAnagramsBruteForce(strs []string) [][]string {
	var result [][]string
	visited := make([]bool, len(strs)) // Track visited strings

	for i := 0; i < len(strs); i++ {
		if visited[i] {
			continue // Skip if already grouped
		}
		var group []string
		group = append(group, strs[i]) // Start a new group
		for j := i + 1; j < len(strs); j++ {
			if !visited[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j]) // Add anagram to group
				visited[j] = true               // Mark as visited
			}
		}
		result = append(result, group) // Add the group to results
	}
	return result
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false // Different lengths mean not anagrams
	}
	counts := make(map[rune]int) // Map to count characters
	for _, char := range s1 {
		counts[char]++ // Count characters in s1
	}
	for _, char := range s2 {
		counts[char]-- // Decrease count for characters in s2
		if counts[char] < 0 {
			return false // More characters in s2 means not anagrams
		}
	}
	return true // All counts are zero; they are anagrams
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagramsBruteForce(strs)
	fmt.Println(result) // Output: [[eat tea ate] [tan nat] [bat]]
}
