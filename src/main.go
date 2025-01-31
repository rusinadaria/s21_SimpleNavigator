package main

import (
	"fmt"
)

func main() {
	LoadGraphFromFile("txt/undirected.txt")

    err := graph.ExportGraphToDot("graph.dot")
    if err != nil {
        fmt.Println("Error while creating file:", err)
    }

    visitedVertices := DepthFirstSearch(&graph, 1)
	fmt.Println("Visited vertices:", visitedVertices)
}