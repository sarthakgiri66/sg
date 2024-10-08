package main

import (
	"fmt"
)

func containsDuplicate(strings []string) bool {
  if len(nums) <= 1 {
    return false
  }
    // Create a map to track the strings we've seen
	seen := make(map[string]bool)

    // Iterate over the slice of strings
	for _, str := range strings {
        // If the string is already in the map, we found a duplicate
		if seen[str] {
			return true
		}
        // Mark the string as seen
		seen[str] = true
	}

    // If no duplicates are found, return false
	return false
}

func main() {
    strings := []string{"apple", "banana", "cherry", "date", "elderberry", "apple"}
    if containsDuplicate(strings) {
        fmt.Println("The slice contains duplicates.")
    } else {
        fmt.Println("The slice does not contain duplicates.")
    }
}
