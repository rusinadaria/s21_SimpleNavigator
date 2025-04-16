package main

import (
	queue "graph/queue"
	"graph/stack"
	// "math"
)

type GraphAlgorithms struct {
}

func DepthFirstSearch(graph *Graph, start_vertex int) []int { //stack
	var result []int
	visited := make(map[int]bool)
	s := stack.NewStack()
	s.Push(start_vertex)
	for !s.IsEmpty() {
		vertex, _ := s.Pop()
		visited[vertex] = true
		result = append(result, vertex)

		for j, isConnected := range graph.matrix[vertex] {
			if isConnected == 1 && !visited[j] {
				s.Push(j)
			}
		}
	}
	return result
}

func BreadthFirstSearch(graph *Graph, start_vertex int) []int { //queue
	var result []int
	visited := make(map[int]bool)
	q := queue.NewQueue()
	q.Push(start_vertex)
	for !q.IsEmpty() {
		vertex, _ := q.Pop()
		visited[vertex] = true
		result = append(result, vertex)

		for j, isConnected := range graph.matrix[vertex] {
			if isConnected == 1 && !visited[j] {
				q.Push(j)
			}
		}
	}
	return result
}

// func GetShortestPathBetweenVertices(graph *Graph, vertex1 int, vertex2 int) { // алгоитм Дейкстры
// 	// var visited map[string]bool
// 	visited := make(map[string]bool)
// 	// var distance map[string]int

// 	// for key, value := range visited {
// 	// 	// visited[graph.matrix] = 
// 	// }

// }
