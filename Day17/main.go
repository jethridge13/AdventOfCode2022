package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Shape struct {
	w     int
	h     int
	shape [][7]string
}

func getBar() Shape {
	shape := [][7]string{{".", ".", "@", "@", "@", "@", "."}}
	return Shape{w: 4, h: 1, shape: shape}
}

func getPlus() Shape {
	shape := [][7]string{
		{".", ".", ".", "@", ".", ".", "."},
		{".", ".", "@", "@", "@", ".", "."},
		{".", ".", ".", "@", ".", ".", "."},
	}
	return Shape{w: 3, h: 3, shape: shape}
}

func getJ() Shape {
	shape := [][7]string{
		{".", ".", ".", ".", "@", ".", "."},
		{".", ".", ".", ".", "@", ".", "."},
		{".", ".", "@", "@", "@", ".", "."},
	}
	return Shape{w: 3, h: 3, shape: shape}
}

func getPipe() Shape {
	shape := [][7]string{
		{".", ".", "@", ".", ".", ".", "."},
		{".", ".", "@", ".", ".", ".", "."},
		{".", ".", "@", ".", ".", ".", "."},
		{".", ".", "@", ".", ".", ".", "."},
	}
	return Shape{w: 1, h: 4, shape: shape}
}

func getSquare() Shape {
	shape := [][7]string{
		{".", ".", "@", "@", ".", ".", "."},
		{".", ".", "@", "@", ".", ".", "."},
	}
	return Shape{w: 2, h: 2, shape: shape}
}

func getShape(i int) Shape {
	switch i % 5 {
	case 0:
		return getBar()
	case 1:
		return getPlus()
	case 2:
		return getJ()
	case 3:
		return getPipe()
	case 4:
		return getSquare()
	default:
		panic("HOW?")
	}
}

func getStartingGraph() [][7]string {
	graph := [][7]string{}
	for i := 0; i < 3; i++ {
		row := [7]string{".", ".", ".", ".", ".", ".", "."}
		graph = append(graph, row)
	}
	return graph
}

func getSequence(path string) []string {
	scanner := util.GetFileScanner(path)
	if scanner.Scan() {
		line := scanner.Text()
		return strings.Split(line, "")
	}
	panic("Bad input")
}

func moveRight(graph [][7]string, l, r, down, height int) bool {
	// Check if move is possible first
	if r >= len(graph[0])-1 {
		return false
	}
	for i := down; i >= down-height+1; i-- {
		if graph[i][r] == "@" && graph[i][r+1] != "." {
			return false
		}
	}
	// Move rock
	for i := down; i >= down-height+1; i-- {
		for j := r; j >= l; j-- {
			graph[i][j+1] = graph[i][j]
			graph[i][j] = "."
		}
	}
	return true
}

func moveLeft(graph [][7]string, l, r, down, height int) bool {
	// Check if move is possible first
	if l <= 0 {
		return false
	}
	for i := down; i >= down-height+1; i-- {
		if graph[i][r] == "@" && graph[i][l-1] != "." {
			return false
		}
	}
	// Move rock
	for i := down; i >= down-height+1; i-- {
		for j := l; j <= r; j++ {
			graph[i][j-1] = graph[i][j]
			graph[i][j] = "."
		}
	}
	return true
}

func solidifyRock(graph [][7]string, l, r, down int) {
	// Maximum block height is 4
	for i := down - 3; i <= down; i++ {
		if i < 0 {
			continue
		}
		for j := l; j <= r; j++ {
			if graph[i][j] == "@" {
				graph[i][j] = "#"
			}
		}
	}
}

func moveDown(graph [][7]string, l, r, down int) bool {
	// Test for collision
	if down == len(graph)-1 {
		solidifyRock(graph, l, r, down)
		return false
	}
	for i := l; i <= r; i++ {
		if graph[down][i] == "@" && graph[down+1][i] == "#" {
			solidifyRock(graph, l, r, down)
			return false
		}
	}
	// Move down
	for i := down; i < down+4 && i >= 0; i-- {
		for j := l; j <= r; j++ {
			if graph[i][j] == "@" {
				graph[i+1][j] = "@"
				graph[i][j] = "."
			}
		}
	}
	return true
}

func part1(path string) int {
	sequence := getSequence(path)
	graph := getStartingGraph()
	si := 0
	for i := 0; i < 2023; i++ {
		shape := getShape(i)
		graph = append(shape.shape, graph...)
		l := 2
		r := l + shape.w - 1
		down := 0 + shape.h - 1
		for {
			// Push rock
			move := sequence[si]
			si += 1
			if si > len(sequence)-1 {
				si = 0
			}
			// Check boundary of tunnel, then check for any settled rocks, then move
			if move == ">" {
				if moveRight(graph, l, r, down, shape.h) {
					l += 1
					r += 1
				}
			} else {
				if moveLeft(graph, l, r, down, shape.h) {
					l -= 1
					r -= 1
				}
			}
			// Drop rock
			if !moveDown(graph, l, r, down) {
				break
			} else {
				down += 1
			}
		}
		// Add clear space to top of graph
		spaceToAdd := 3
		foundTop := false
		for j := 0; !foundTop; j++ {
			for k := 0; k < len(graph[j]); k++ {
				if graph[j][k] != "." {
					foundTop = true
					break
				}
			}
			if !foundTop {
				spaceToAdd -= 1
			}
		}
		if spaceToAdd > 0 {
			space := [][7]string{}
			for j := 0; j < spaceToAdd; j++ {
				space = append(space, [7]string{".", ".", ".", ".", ".", ".", "."})
			}
			graph = append(space, graph...)
		} else if spaceToAdd < 0 {
			index := spaceToAdd * -1
			graph = graph[index:]
		}
	}
	return len(graph) - 3
}

func main() {
	file := "example.txt"
	// Part 1: >3033
	fmt.Println(part1(file))
	// Part 2:
}
