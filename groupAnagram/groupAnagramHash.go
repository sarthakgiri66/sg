//Character Count Approach (O(n * k))
//This method uses an array to count the occurrences of each character, providing a fixed-length key for the hash map.
//Fixed-Size Array: Use a fixed-size array of length 26 to count occurrences of each character (assuming only lowercase English letters).
//Counting: For each string, count the characters and use the array as a key in the hash map.
//Group Results: Append each original string to the corresponding key in the hash map.
//Time Complexity: O(n * k), where n is the number of strings and k is the maximum length of a string.
//Space Complexity: O(n) for the hash map storage.


package main

import (
	"fmt"
)

func groupAnagramsCharacterCount(strs []string) [][]string {
	anagrams := make(map[[26]int][]string) // Map to hold character counts as keys

	for _, str := range strs {
		count := [26]int{} // Initialize a fixed-size array for character counts
		for _, char := range str {
			count[char-'a']++ // Increment the count for this character
		}
		anagrams[count] = append(anagrams[count], str) // Group by character count
	}

	var result [][]string
	for _, group := range anagrams {
		result = append(result, group) // Collect the groups
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagramsCharacterCount(strs)
	fmt.Println(result) // Output: [[eat tea ate] [tan nat] [bat]]
}
