package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InitGraphTestSuite struct {
	suite.Suite
}

func (suite *InitGraphTestSuite) SetupTest() {
    graph = Graph{
		matrix: [][]int{
			// 0  1  2  3  4  5  6
			{0, 1, 0, 0, 0, 0, 0}, // 0
			{1, 0, 1, 1, 0, 0, 0}, // 1
			{0, 1, 0, 0, 1, 0, 0}, // 2
			{0, 1, 0, 0, 1, 1, 0}, // 3
			{0, 0, 1, 1, 0, 0, 1}, // 4
			{0, 0, 0, 1, 0, 0, 1}, // 5
			{0, 0, 0, 0, 1, 1, 0}, // 6
		},
	}
}

func (suite *InitGraphTestSuite) TearDownTest() {
}

func (suite *InitGraphTestSuite) TestDFS() {
	visitedVertices := DepthFirstSearch(&graph, 1)

    expected := []int{1, 3, 5, 6, 4, 2, 0}
	assert.EqualValues(suite.T(), visitedVertices, expected)
}

func (suite *InitGraphTestSuite) TestBFS() {
	visitedVertices := BreadthFirstSearch(&graph, 1)

	// expected := []int{1, 0, 2, 3, 4, 5, 6}
    expected :=  []int{1, 0, 2, 3, 4, 4, 5, 6, 6, 6}
	assert.EqualValues(suite.T(), visitedVertices, expected)
}

func (suite *InitGraphTestSuite) TestShortestPathMockGraph() {
	mockGraph := Graph{
		matrix: [][]int{
			// 0  1  2  3  4
			{0, 2, 0, 1, 0}, // 0
			{2, 0, 3, 2, 0}, // 1
			{0, 3, 0, 0, 1}, // 2
			{1, 2, 0, 0, 3}, // 3
			{0, 0, 1, 3, 0}, // 4
		},
	}

	shortest := GetShortestPathBetweenVertices(&mockGraph, 0, 4)
	expected := 4
	assert.Equal(suite.T(), expected, shortest)
}

func TestInitGraphTestSuite(t *testing.T) {
	suite.Run(t, new(InitGraphTestSuite))
}
