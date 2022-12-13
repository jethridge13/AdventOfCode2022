package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func printGridPart2(inp [][]string, snake []Segment) {
	grid := make([][]string, len(inp))
	for i := range inp {
		grid[i] = make([]string, len(inp[i]))
		copy(grid[i], inp[i])
	}
	for i, node := range snake {
		var marker string
		if i == 0 {
			marker = "H"
		} else {
			marker = strconv.Itoa(i)
		}
		grid[node.pos[0]][node.pos[1]] = marker
	}
	fmt.Println("--- START ---")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("---- END ----")
}

func makeGrid(size int) [][]string {
	grid := make([][]string, size)
	for i := range grid {
		row := make([]string, size)
		for j := range row {
			row[j] = "."
		}
		grid[i] = row
	}
	return grid
}

func makeBlankRow(size int) []string {
	row := make([]string, size)
	for i := range row {
		row[i] = "."
	}
	return row
}

func countPath(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c == "#" {
				count += 1
			}
		}
	}
	return count
}

func expandGrid(grid *[][]string, h []int, t []int) {
	// Add space above
	topSize := len(*grid)
	topGrid := makeGrid(topSize)
	h[0] += topSize
	t[0] += topSize
	*grid = append(topGrid, *grid...)
	// Add space left
	leftSize := len((*grid)[0])
	h[1] += leftSize
	t[1] += leftSize
	for i, row := range *grid {
		blankRow := makeBlankRow(len(row))
		(*grid)[i] = append(blankRow, row...)
	}
	// Add space bottom
	botSize := len(*grid)
	botGrid := makeGrid(botSize)
	*grid = append(*grid, botGrid...)
	// Add space right
	for i, row := range *grid {
		blankRow := makeBlankRow(len(row))
		(*grid)[i] = append(row, blankRow...)
	}
}

func expandGridTail(grid *[][]string, snake []Segment) {
	// Add space above
	topSize := len(*grid)
	topGrid := makeGrid(topSize)
	*grid = append(topGrid, *grid...)
	// Add space left
	leftSize := len((*grid)[0])
	for i, row := range *grid {
		blankRow := makeBlankRow(len(row))
		(*grid)[i] = append(blankRow, row...)
	}
	// Add space bottom
	botSize := len(*grid)
	botGrid := makeGrid(botSize)
	*grid = append(*grid, botGrid...)
	// Add space right
	for i, row := range *grid {
		blankRow := makeBlankRow(len(row))
		(*grid)[i] = append(row, blankRow...)
	}
	for i, node := range snake {
		node.pos[0] += topSize
		node.lastPos[0] += topSize
		node.pos[1] += leftSize
		node.lastPos[1] += leftSize
		snake[i] = node
	}
}

func moveGrid(grid *[][]string, h []int, t []int, dir string, dis int) {
	switch dir {
	case "U":
		for h[0]-dis < 0 {
			expandGrid(grid, h, t)
		}
		h[0] -= dis
		for t[0]-h[0] > 1 {
			if h[1] != t[1] {
				t[1] += h[1] - t[1]
			}
			t[0] -= 1
			(*grid)[t[0]][t[1]] = "#"
		}
	case "R":
		for h[1]+dis >= len(*grid) {
			expandGrid(grid, h, t)
		}
		h[1] += dis
		for h[1]-t[1] > 1 {
			if h[0] != t[0] {
				t[0] += h[0] - t[0]
			}
			t[1] += 1
			(*grid)[t[0]][t[1]] = "#"
		}
	case "D":
		for h[0]+dis >= len(*grid) {
			expandGrid(grid, h, t)
		}
		h[0] += dis
		for h[0]-t[0] > 1 {
			if h[1] != t[1] {
				t[1] += h[1] - t[1]
			}
			t[0] += 1
			(*grid)[t[0]][t[1]] = "#"
		}
	case "L":
		for h[1]-dis < 0 {
			expandGrid(grid, h, t)
		}
		h[1] -= dis
		for t[1]-h[1] > 1 {
			if h[0] != t[0] {
				t[0] += h[0] - t[0]
			}
			t[1] -= 1
			(*grid)[t[0]][t[1]] = "#"
		}
	}
}

func moveGridTail(grid *[][]string, snake []Segment, dir string, dis int) {
	moved := 0
	for moved < dis {
		switch dir {
		case "U":
			for snake[0].pos[0]-dis < 0 {
				expandGridTail(grid, snake)
			}
			snake[0].lastPos = snake[0].pos
			snake[0].pos[0] -= 1
			moved += 1
		case "R":
			for snake[0].pos[1]+dis >= len(*grid) {
				expandGridTail(grid, snake)
			}
			snake[0].lastPos = snake[0].pos
			snake[0].pos[1] += 1
			moved += 1
		case "D":
			for snake[0].pos[0]+dis >= len(*grid) {
				expandGridTail(grid, snake)
			}
			snake[0].lastPos = snake[0].pos
			snake[0].pos[0] += 1
			moved += 1
		case "L":
			for snake[0].pos[1]-dis < 0 {
				expandGridTail(grid, snake)
			}
			snake[0].lastPos = snake[0].pos
			snake[0].pos[1] -= 1
			moved += 1
		}
		// Catch up snake tail
		for i := 1; i < len(snake); i++ {
			head := snake[i-1]
			dx := head.pos[0] - snake[i].pos[0]
			dy := head.pos[1] - snake[i].pos[1]
			if util.Abs(dx) == 2 && dy == 0 {
				if dx == 2 {
					snake[i].pos[0] += 1
				} else {
					snake[i].pos[0] -= 1
				}
			} else if util.Abs(dy) == 2 && dx == 0 {
				if dy == 2 {
					snake[i].pos[1] += 1
				} else {
					snake[i].pos[1] -= 1
				}
			} else if (dx == 1 && dy == 2) || (dy == 1 && dx == 2) || (dx == 2 && dy == 2) {
				snake[i].pos[0] += 1
				snake[i].pos[1] += 1
			} else if (dx == 1 && dy == -2) || (dy == -1 && dx == 2) || (dx == 2 && dy == -2) {
				snake[i].pos[0] += 1
				snake[i].pos[1] -= 1
			} else if (dx == -1 && dy == 2) || (dy == 1 && dx == -2) || (dx == -2 && dy == 2) {
				snake[i].pos[0] -= 1
				snake[i].pos[1] += 1
			} else if (dx == -1 && dy == -2) || (dy == -1 && dx == -2) || (dx == -2 && dy == -2) {
				snake[i].pos[0] -= 1
				snake[i].pos[1] -= 1
			}
			if i == len(snake)-1 {
				(*grid)[snake[i].pos[0]][snake[i].pos[1]] = "#"
			}
		}
	}
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	grid := makeGrid(2)
	h, t := []int{0, 0}, []int{0, 0}
	grid[t[0]][t[1]] = "#"
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		dir := parts[0]
		dis, _ := strconv.Atoi(parts[1])
		moveGrid(&grid, h, t, dir, dis)
	}
	return countPath(grid)
}

type Segment struct {
	pos     [2]int
	lastPos [2]int
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	grid := makeGrid(2)
	snake := make([]Segment, 10)
	for i := range snake {
		s := Segment{pos: [2]int{0, 0}}
		snake[i] = s
	}
	grid[snake[9].pos[0]][snake[9].pos[1]] = "#"
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		dir := parts[0]
		dis, _ := strconv.Atoi(parts[1])
		moveGridTail(&grid, snake, dir, dis)
	}
	return countPath(grid)
}

func main() {
	file := "input.txt"
	// Part 1: 6266
	fmt.Println(part1(file))
	// Part 2: 2369
	fmt.Println(part2(file))
}
