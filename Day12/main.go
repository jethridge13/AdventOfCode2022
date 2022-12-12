package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Node struct {
	pos       [2]int
	name      string
	val       int
	pathValue int
	edges     [][2]int
}

func getVal(s string) int {
	if s == "E" {
		return 27
	} else if s == "S" {
		return 0
	}
	var r rune
	for _, c := range s {
		r = c
	}
	return int(r) - 96
}

func getVal2(s string) int {
	if s == "E" {
		return 0
	} else if s == "S" {
		return 27
	}
	var r rune
	for _, c := range s {
		r = c
	}
	return (int(r) - 123) * -1
}

func parseMatrix(path string) [][]string {
	scanner := util.GetFileScanner(path)
	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}
	return matrix
}

func parseGraph(matrix [][]string) (map[[2]int]Node, [2]int, [2]int) {
	var start [2]int
	var end [2]int
	graph := make(map[[2]int]Node)
	for i := range matrix {
		for j := range matrix[i] {
			pos := [2]int{i, j}
			if matrix[i][j] == "S" {
				start = pos
			} else if matrix[i][j] == "E" {
				end = pos
			}
			val := getVal(matrix[i][j])
			n := Node{name: matrix[i][j], pos: pos, val: val, pathValue: math.MaxInt64}
			// Check neighbors for edges
			edges := make([][2]int, 0)
			// Check North
			if i > 0 && getVal(matrix[i-1][j]) <= val+1 {
				edges = append(edges, [2]int{i - 1, j})
			}
			// Check East
			if j < len(matrix[i])-1 && getVal(matrix[i][j+1]) <= val+1 {
				edges = append(edges, [2]int{i, j + 1})
			}
			// Check West
			if j > 0 && getVal(matrix[i][j-1]) <= val+1 {
				edges = append(edges, [2]int{i, j - 1})
			}
			// Check South
			if i < len(matrix)-1 && getVal(matrix[i+1][j]) <= val+1 {
				edges = append(edges, [2]int{i + 1, j})
			}
			n.edges = edges
			graph[pos] = n
		}
	}
	return graph, start, end
}

func parseGraph2(matrix [][]string) (map[[2]int]Node, [2]int, [2]int) {
	var start [2]int
	var end [2]int
	graph := make(map[[2]int]Node)
	for i := range matrix {
		for j := range matrix[i] {
			pos := [2]int{i, j}
			if matrix[i][j] == "S" {
				start = pos
			} else if matrix[i][j] == "E" {
				end = pos
			}
			val := getVal2(matrix[i][j])
			n := Node{name: matrix[i][j], pos: pos, val: val, pathValue: math.MaxInt64}
			// Check neighbors for edges
			edges := make([][2]int, 0)
			// Check North
			if i > 0 && getVal2(matrix[i-1][j]) <= val+1 {
				edges = append(edges, [2]int{i - 1, j})
			}
			// Check East
			if j < len(matrix[i])-1 && getVal2(matrix[i][j+1]) <= val+1 {
				edges = append(edges, [2]int{i, j + 1})
			}
			// Check West
			if j > 0 && getVal2(matrix[i][j-1]) <= val+1 {
				edges = append(edges, [2]int{i, j - 1})
			}
			// Check South
			if i < len(matrix)-1 && getVal2(matrix[i+1][j]) <= val+1 {
				edges = append(edges, [2]int{i + 1, j})
			}
			n.edges = edges
			graph[pos] = n
		}
	}
	return graph, start, end
}

func part1(path string) int {
	matrix := parseMatrix(path)
	graph, start, _ := parseGraph(matrix)
	// Dijkstra here we go
	visited := make(map[[2]int]bool)
	startNode := graph[start]
	startNode.pathValue = 0
	graph[start] = startNode
	toCheck := [][2]int{start}
	for len(toCheck) > 0 {
		node := graph[toCheck[0]]
		toCheck = toCheck[1:]
		if node.name == "E" {
			return node.pathValue
		}
		if visited[node.pos] {
			continue
		}
		visited[node.pos] = true
		for _, neighbor := range node.edges {
			if visited[neighbor] {
				continue
			}
			edge := graph[neighbor]
			if node.pathValue+1 < edge.pathValue {
				edge.pathValue = node.pathValue + 1
				graph[neighbor] = edge
			}
			toCheck = append(toCheck, neighbor)
		}
	}
	return -1
}

func part2(path string) int {
	matrix := parseMatrix(path)
	graph, _, start := parseGraph2(matrix)
	// Dijkstra here we go
	visited := make(map[[2]int]bool)
	startNode := graph[start]
	startNode.pathValue = 0
	graph[start] = startNode
	toCheck := [][2]int{start}
	for len(toCheck) > 0 {
		node := graph[toCheck[0]]
		toCheck = toCheck[1:]
		if visited[node.pos] {
			continue
		}
		visited[node.pos] = true
		for _, neighbor := range node.edges {
			if visited[neighbor] {
				continue
			}
			edge := graph[neighbor]
			if node.pathValue+1 < edge.pathValue {
				edge.pathValue = node.pathValue + 1
				graph[neighbor] = edge
			}
			toCheck = append(toCheck, neighbor)
		}
	}
	maxPath := math.MaxInt64
	for pos := range graph {
		node := graph[pos]
		if node.name != "a" {
			continue
		}
		if node.pathValue < maxPath {
			maxPath = node.pathValue
		}
	}
	return maxPath
}

func main() {
	file := "input.txt"
	// Part 1: 490
	fmt.Println(part1(file))
	// Part 2: 488
	fmt.Println(part2(file))
}
