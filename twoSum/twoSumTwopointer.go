//Two-Pointer Technique: Works well if you can sort the input, but involves extra time for sorting (O(n log n)).
//Time Complexity: O(n log n) due to the sorting step.
//Space Complexity: O(n) for storing the pairs.


package main

import (
	"fmt"
	"sort"
)

func twoSumTwoPointers(nums []int, target int) []int {
	type Pair struct {
		Index int
		Value int
	}

	// Create a slice of pairs to store the original indices
	pairs := make([]Pair, len(nums))
	for i, num := range nums {
		pairs[i] = Pair{Index: i, Value: num}
	}

	// Sort pairs based on the values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	// Initialize two pointers
	left, right := 0, len(pairs)-1
	for left < right {
		sum := pairs[left].Value + pairs[right].Value
		if sum == target {
			return []int{pairs[left].Index, pairs[right].Index}
		} else if sum < target {
			left++ // Move the left pointer to the right
		} else {
			right-- // Move the right pointer to the left
		}
	}
	return nil // Return nil if no pair is found
}

func main() {
	fmt.Println(twoSumTwoPointers([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
}
