package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
    "fmt"
    "github.com/stretchr/testify/suite"
    "os"
)

type InitGraphTestSuite struct {
    suite.Suite
}

func (suite *InitGraphTestSuite) SetupTest() {
    LoadGraphFromFile("txt/undirected.txt")
    err := graph.ExportGraphToDot("graph.dot")
    if err != nil {
        fmt.Println("Error while creating file:", err)
    }
}

func (suite *InitGraphTestSuite) TearDownTest() {
	err := os.Remove("graph.dot")
	if err != nil {
		fmt.Println("Error while deleting file:", err)
	}
}

func (suite *InitGraphTestSuite) TestDFS() {
    visitedVertices := DepthFirstSearch(&graph, 1)

    expected := [7]int{1, 3, 4, 0, 2, 2, 0}
    assert.EqualValues(suite.T(), visitedVertices, expected)
}

func (suite *InitGraphTestSuite) TestBFS() {
    visitedVertices := BreadthFirstSearch(&graph, 1)

    expected := [7]int{1, 0, 2, 3, 4, 3, 4}
    assert.EqualValues(suite.T(), visitedVertices, expected)
}

func TestInitGraphTestSuite(t *testing.T) {
	suite.Run(t, new(InitGraphTestSuite))
}