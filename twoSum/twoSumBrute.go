//Brute Force: Simple but inefficient for large inputs (O(n²)).
//Time Complexity: O(n²) because for each element, we potentially check all other elements.
//Space Complexity: O(1) as we only use a few variables.


package main

import "fmt"

func twoSumBruteForce(nums []int, target int) []int {
	n := len(nums) // Get the length of the input array
	for i := 0; i < n; i++ { // Outer loop to iterate through each element
		for j := i + 1; j < n; j++ { // Inner loop to iterate through the remaining elements
			if nums[i]+nums[j] == target { // Check if the sum of the two elements equals the target
				return []int{i, j} // Return the indices of the two elements
			}
		}
	}
	return nil // Return nil if no pair is found
}

func main() {
	fmt.Println(twoSumBruteForce([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
}
