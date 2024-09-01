package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
  if len(nums) <= 1 {
    return fasle
  }
    // Create a map to track the elements we've seen
	seen := make(map[int]bool)

    // Iterate over the slice
	for _, num := range nums {
        // If the element is already in the map, we found a duplicate
		if seen[num] {
			return true
		}
        // Mark the element as seen
		seen[num] = true
	}

    // If no duplicates are found, return false
	return false
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
    if containsDuplicate(nums) {
        fmt.Println("The slice contains duplicates.")
    } else {
        fmt.Println("The slice does not contain duplicates.")
    }
}
