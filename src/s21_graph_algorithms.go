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
	// visited[start_vertex] = true
	s.Push(start_vertex)
	visited[start_vertex] = true
	for !s.IsEmpty() {
		vertex, _ := s.Pop()
		// visited[vertex] = true
		result = append(result, vertex)

		for j, isConnected := range graph.matrix[vertex] {
			if isConnected == 1 && !visited[j] {
				s.Push(j)
				visited[j] = true
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

func GetShortestPathBetweenVertices(graph *Graph, vertex1 int, vertex2 int) int { // алгоитм Дейкстры
    n := len(graph.matrix)

    d := make([]int, n)
    for i := 0; i < n; i++ {
        d[i] = 1<<31 - 1
    }
    d[vertex1] = 0

    a := make([]int, n)

    for {
        u := -1
        minDist := 1<<31 - 1
        for i := 0; i < n; i++ {
            if a[i] == 0 && d[i] < minDist {
                minDist = d[i]
                u = i
            }
        }

        if u == -1 {
            break
        }

        a[u] = 1

        for v := 0; v < n; v++ {
            weight := graph.matrix[u][v]
            if weight > 0 && a[v] == 0 {
                if d[v] > d[u]+weight {
                    d[v] = d[u] + weight
                }
            }
        }
    }

   if d[vertex2] == 1<<31-1 {
		return -1
	}
	return d[vertex2]
}
