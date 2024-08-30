/*
2699 You are given an undirected weighted connected graph containing n nodes labeled from 0 to n - 1, and an integer array edges where edges[i] = [ai, bi, wi] indicates that there is an edge between nodes ai and bi with weight wi.

Some edges have a weight of -1 (wi = -1), while others have a positive weight (wi > 0).

Your task is to modify all edges with a weight of -1 by assigning them positive integer values in the range [1, 2 * 109] so that the shortest distance between the nodes source and destination becomes equal to an integer target. If there are multiple modifications that make the shortest distance between source and destination equal to target, any of them will be considered correct.

Return an array containing all edges (even unmodified ones) in any order if it is possible to make the shortest distance from source to destination equal to target, or an empty array if it's impossible.

Note: You are not allowed to modify the weights of edges with initial positive weights.

Example 1:

Input: n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0, destination = 1, target = 5
Output: [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
Explanation: The graph above shows a possible modification to the edges, making the distance from 0 to 1 equal to 5.
Example 2:

Input: n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
Output: []
Explanation: The graph above contains the initial edges. It is not possible to make the distance from 0 to 2 equal to 6 by modifying the edge with weight -1. So, an empty array is returned.
Example 3:

Input: n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination = 2, target = 6
Output: [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
Explanation: The graph above shows a modified graph having the shortest distance from 0 to 2 as 6.
*/
package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	n := 5
	edges := [][]int{
		{4, 1, -1},
		{2, 0, -1},
		{0, 3, -1},
		{4, 3, -1},
	}
	source := 0
	destination := 1
	target := 5

	result := modifiedGraphEdges(n, edges, source, destination, target)
	fmt.Println(result)
}

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	adjList := make([][]int, n)
	for i := range adjList {
		adjList[i] = make([]int, n)
	}
	for _, edge := range edges {
		a, b, w := edge[0], edge[1], edge[2]
		if w == -1 {
			continue
		}
		adjList[a][b] = w
		adjList[b][a] = w
	}

	distance := find(adjList, n, source, destination)
	if distance[destination] < target {
		return [][]int{}
	}

	if distance[destination] == target {
		for _, edge := range edges {
			if edge[2] == -1 {
				edge[2] = 2e9
			}
		}
		return edges
	}

	for i, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if w == -1 {
			edges[i][2] = 1
			adjList[u][v] = 1
			adjList[v][u] = 1
			distance = find(adjList, n, source, destination)
			if distance[destination] <= target {
				edges[i][2] += (target - distance[destination])
				for j := i + 1; j < len(edges); j++ {
					if edges[j][2] == -1 {
						edges[j][2] = 2e9
					}
				}
				return edges
			}
		}
	}
	return [][]int{}
}

func find(adjList [][]int, n int, src int, _ int) []int {
	distance := make([]int, n)
	for i := range distance {
		if i == src {
			continue
		}
		distance[i] = math.MaxInt32
	}
	pq := PriorityQueue{{src, 0}}
	heap.Init(&pq)
	visited := make([]bool, n)
	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*State)
		if visited[state.node] {
			continue
		}

		visited[state.node] = true
		for child, cost := range adjList[state.node] {
			if cost == 0 {
				continue
			}
			if distance[child] > state.cost+cost {
				distance[child] = state.cost + cost
				heap.Push(&pq, &State{child, distance[child]})
			}
		}
	}
	return distance
}

type State struct {
	node int
	cost int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	state := x.(*State)
	*pq = append(*pq, state)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	*pq = old[0 : n-1]
	return state
}
