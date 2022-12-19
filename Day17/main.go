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
	line := scanner.Text()
	return strings.Split(line, "")
}

func part1(path string) int {
	sequence := getSequence(path)
	graph := getStartingGraph()
	si := 0
	for i := 0; i < 2023; i++ {
		shape := getShape(i)
		graph = append(shape.shape, graph...)
		l := 3
		r := l + shape.w
		for {
			// Push rock
			move := sequence[si]
			si += 1
			if si > len(sequence)-1 {
				si = 0
			}
			// Check boundary of tunnel, then check for any settled rocks, then move
			if move == ">" {
				if r < len(graph[0]) {

				}
			} else {
				if l > 0 {

				}
			}
			// Drop rock

		}
	}
	return 0
}

func main() {
	file := "example.txt"
	// Part 1:
	fmt.Println(part1(file))
	// Part 2:
}
