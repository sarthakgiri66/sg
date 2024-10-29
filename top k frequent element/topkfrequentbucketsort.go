// Bucket Sort (O(N))
//This method can achieve a time complexity of O(N) under certain conditions, particularly when the range of input values is limited.
//Time complexity:
//Counting Frequencies: O(N) for creating the frequency map.
//Creating Buckets: O(N) to distribute elements into buckets.
//Extracting Top-K Elements: O(K) in the worst case if we need to go through all the buckets (but K is usually much smaller than N).
//Space Complexity: The space complexity is O(N) for the frequency map and O(N) for the buckets in the worst case, which also leads to O(N) overall space complexity.


package main

import (
    "fmt"
)

func topKFrequent(nums []int, k int) []int {
    // Step 1: Count frequencies
    frequency := make(map[int]int)
    for _, num := range nums {
        frequency[num]++
    }

    // Step 2: Create buckets
    buckets := make([][]int, len(nums)+1)    // Max frequency can be N
    for num, freq := range frequency {
        buckets[freq] = append(buckets[freq], num)
    }

    // Step 3: Gather top-k elements
    result := []int{}
    for i := len(buckets) - 1; i >= 0; i-- {
        for _, num := range buckets[i] {
            result = append(result, num)
            if len(result) == k {
                return result
            }
        }
    }
    return result
}

func main() {
    nums := []int{1, 1, 1, 2, 2, 3}
    k := 2
    fmt.Println(topKFrequent(nums, k)) // Output: [1, 2]
}
