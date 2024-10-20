//Hash Map: Most efficient for this problem (O(n)), great for quick lookups.
//Time Complexity: O(n) because we make a single pass through the array, performing O(1) operations for each element.
//Space Complexity: O(n) in the worst case due to the storage used in the hash map.


package main

import "fmt"

func twoSumHashMap(nums []int, target int) []int {
    numMap := make(map[int]int) // Create a hash map to store numbers and their indices

    for i, num := range nums {
        complement := target - num // Calculate the complement
        if index, found := numMap[complement]; found {
            return []int{index, i} // Return indices if complement is found
        }
        numMap[num] = i // Store the current number and its index
    }
    return nil // Return nil if no pair is found
}

func main() {
    fmt.Println(twoSumHashMap([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
}

