//Time Complexity: O(n * k log k), where n is the number of strings and k is the maximum length of a string (for sorting).
//Space Complexity: O(n) for the hash map storage.



package main

import (
	"fmt"
	"sort"
)

func groupAnagramsSorting(strs []string) [][]string {
	anagrams := make(map[string][]string) // Map to hold sorted strings as keys

	for _, str := range strs {
		// Sort the string and use it as a key
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str) // Group by sorted string
	}

	var result [][]string
	for _, group := range anagrams {
		result = append(result, group) // Collect the groups
	}
	return result
}

func sortString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j] // Sort runes
	})
	return string(runes) // Convert back to string
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagramsSorting(strs)
	fmt.Println(result) // Output: [[eat tea ate] [tan nat] [bat]]
}

