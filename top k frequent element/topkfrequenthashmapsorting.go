//Hash Map + Sorting
//This method is straightforward and utilizes a combination of a hash map for counting frequencies and a sorting algorithm to determine the top elements.
//Time Complexity:
//Counting frequencies takes O(N), where N is the number of elements in the array.
//Sorting the list of pairs takes O(M log M), where M is the number of unique elements. In the worst case, M could be N (if all elements are unique), leading to an overall complexity of O(N log N).
//Space Complexity:  We use O(M) space for the hash map to store frequencies and O(M) for the list of pairs, leading to a total space complexity of O(M).



package main

import (
	"fmt"
	"sort"
)

// Function to find the top K frequent elements
func topKFrequent(nums []int, k int) []int {
	// Step 1: Count the frequency of each element
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// Step 2: Create a slice to hold pairs of elements and their frequencies
	type pair struct {
		num  int
		freq int
	}

	pairs := make([]pair, 0, len(count))
	for num, freq := range count {
		pairs = append(pairs, pair{num, freq})
	}

	// Step 3: Sort pairs based on frequency in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	// Step 4: Extract the top K elements
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [1, 2]
}
