//Hash Map + Min Heap (O(N log K))
//Time Complexity: O(N log k)
//Counting frequencies takes O(N).
//Each push and pop operation on the heap takes O(log k), and since we process each unique element (up to N), the overall complexity is O(N log k).
//Space Complexity: O(N) in the worst case for the hash map and O(k) for the heap.




package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int
	frequency int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	minHeap := &MinHeap{}
	for value, freq := range frequency {
		heap.Push(minHeap, Item{value, freq})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(minHeap).(Item)
		result[i] = item.value
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // Output: [1, 2]
}
