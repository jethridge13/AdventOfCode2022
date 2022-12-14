package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func generateGraph(width, height int) [][]string {
	graph := make([][]string, height)
	for i := range graph {
		row := make([]string, width)
		for j := 0; j < width; j++ {
			row[j] = "."
		}
		graph[i] = row
	}
	return graph
}

func getPoints(parts []string) [][2]int {
	pts := make([][2]int, 0)
	for _, part := range parts {
		points := strings.Split(part, ",")
		x, err := strconv.Atoi(points[0])
		util.ErrCheck(err)
		y, err := strconv.Atoi(points[1])
		util.ErrCheck(err)
		coord := [2]int{x, y}
		pts = append(pts, coord)
	}
	return pts
}

func drawLine(graph [][]string, start, end [2]int) {
	graph[start[1]][start[0]] = "#"
	for start[0] != end[0] || start[1] != end[1] {
		if start[0] < end[0] {
			start[0] += 1
		} else if start[0] > end[0] {
			start[0] -= 1
		} else if start[1] < end[1] {
			start[1] += 1
		} else if start[1] > end[1] {
			start[1] -= 1
		}
		graph[start[1]][start[0]] = "#"
	}
}

func generateCave(path string) ([][]string, int) {
	graph := generateGraph(1000, 200)
	lowestPt := 0
	scanner := util.GetFileScanner(path)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		pts := getPoints(parts)
		for i := 1; i < len(pts); i++ {
			start := pts[i-1]
			end := pts[i]
			if start[1] > lowestPt {
				lowestPt = start[1]
			}
			if end[1] > lowestPt {
				lowestPt = end[1]
			}
			drawLine(graph, start, end)
		}
	}
	return graph, lowestPt
}

func simulateSand(graph [][]string, lowestPt int) int {
	count := 0
	sand := [2]int{500, 0}
	lastSandPts := make([][2]int, 0)
	for sand[1] < lowestPt {
		// Check down, then down left, then down right
		if graph[sand[1]+1][sand[0]] == "." {
			lastSandPts = append([][2]int{sand}, lastSandPts...)
			sand[1] += 1
		} else if graph[sand[1]+1][sand[0]-1] == "." {
			lastSandPts = append([][2]int{sand}, lastSandPts...)
			sand[1] += 1
			sand[0] -= 1
		} else if graph[sand[1]+1][sand[0]+1] == "." {
			lastSandPts = append([][2]int{sand}, lastSandPts...)
			sand[1] += 1
			sand[0] += 1
		} else if len(lastSandPts) > 0 && sand[0] == lastSandPts[0][0] && sand[1] == lastSandPts[0][1] {
			sand = lastSandPts[0]
			lastSandPts = lastSandPts[1:]
		} else {
			graph[sand[1]][sand[0]] = "o"
			count += 1
			if len(lastSandPts) > 0 {
				sand = lastSandPts[0]
			} else {
				return count
			}
		}
	}
	return count
}

func part1(path string) int {
	graph, lowestPt := generateCave(path)
	return simulateSand(graph, lowestPt)
}

func part2(path string) int {
	graph, lowestPt := generateCave(path)
	row := make([]string, len(graph[0]))
	for i := range row {
		row[i] = "#"
	}
	graph[lowestPt+2] = row
	return simulateSand(graph, lowestPt+3)
}

func main() {
	file := "input.txt"
	// Part 1: 728
	fmt.Println(part1(file))
	// Part 2: 27623
	fmt.Println(part2(file))
}
