/*
You are given an undirected weighted graph of n nodes (0-indexed), represented by an edge list where edges[i] = [a, b] is an undirected edge connecting the nodes a and b with a probability of success of traversing that edge succProb[i].

Given two nodes start and end, find the path with the maximum probability of success to go from start to end and return its success probability.

If there is no path from start to end, return 0. Your answer will be accepted if it differs from the correct answer by at most 1e-5.

Example 1:

Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
Output: 0.25000
Explanation: There are two paths from start to end, one having a probability of success = 0.2 and the other has 0.5 * 0.5 = 0.25.
Example 2:

Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
Output: 0.30000
Example 3:

Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
Output: 0.00000
Explanation: There is no path between 0 and 2.
*/
package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node        int
	probability float64
}

type MaxHeap []Edge

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].probability > h[j].probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make(map[int][]Edge)

	for i, edge := range edges {
		a, b := edge[0], edge[1]
		prob := succProb[i]
		graph[a] = append(graph[a], Edge{b, prob})
		graph[b] = append(graph[b], Edge{a, prob})
	}

	probabilities := make([]float64, n)
	probabilities[start] = 1.0

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Edge{start, 1.0})

	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(Edge)
		currentNode := current.node
		currentProb := current.probability

		if currentNode == end {
			return currentProb
		}

		for _, neighbor := range graph[currentNode] {
			newProb := currentProb * neighbor.probability
			if newProb > probabilities[neighbor.node] {
				probabilities[neighbor.node] = newProb
				heap.Push(maxHeap, Edge{neighbor.node, newProb})
			}
		}
	}

	return 0.0
}

func main() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	succProb := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	result := maxProbability(n, edges, succProb, start, end)
	fmt.Printf("Output: %.5f\n", result) // Expected Output: 0.25000
}
