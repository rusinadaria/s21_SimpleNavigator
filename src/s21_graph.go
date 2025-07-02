package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

// func main() {
// 	LoadGraphFromFile("txt/undirected.txt")
//     err := graph.ExportGraphToDot("txt/graph.dot")
// 	  if err != nil {
//         fmt.Println("Error exporting to DOT:", err)
//     } else {
//         fmt.Println("Graph exported to txt/graph.dot")
//     }
// }

type Graph struct {
    matrix [][]int
    rows int
    cols int
}

var graph Graph

func LoadGraphFromFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        panic("Unable to open file")
    }
    defer file.Close()

    data := bufio.NewScanner(file)
    countRows := 0
    countCols := 0
    
    for data.Scan() {
        line := data.Text()
        symbols := strings.Split(line, "")
        // symbols := strings.Fields(line)
        countCols = len(symbols)

        row := make([]int, countCols)
        for i, symbol := range symbols {
            value, err := strconv.Atoi(symbol)
            if err != nil {
                fmt.Println("Conversion error:", err)
                return
            }
            row[i] = value
        }
        graph.matrix = append(graph.matrix, row) 
        countRows++
    }
    graph.rows = countRows
    graph.cols = countCols
}

func (g *Graph) ExportGraphToDot(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("error while creating file: %v", err)
    }
    defer file.Close()

    _, err = file.WriteString("graph G {\n")
    if err != nil {
        return fmt.Errorf("error while writing to file: %v", err)
    }

    for i := 0; i < graph.rows; i++ {
        for j := 0; j < graph.cols; j++ {
            if graph.matrix[i][j] != 0 {
                // _, err = file.WriteString(fmt.Sprintf("    %d -- %d;\n", i, j))
                _, err = file.WriteString(fmt.Sprintf("    %d -- %d;\n", i + 1, j + 1))
                if err != nil {
                    return fmt.Errorf("error while writing to file: %v", err)
                }
            }
        }
    }

    _, err = file.WriteString("}\n")
    if err != nil {
        return fmt.Errorf("error while writing to file: %v", err)
    }
    return nil
}