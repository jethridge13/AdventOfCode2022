package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func printGridPart2(inp [][]string, snake [][]int) {
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
		grid[node[0]][node[1]] = marker
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

func expandGridTail(grid *[][]string, snake [][]int) {
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
	for _, node := range snake {
		node[0] += topSize
		node[1] += leftSize
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

func moveGridTail(grid *[][]string, snake [][]int, dir string, dis int) {
	switch dir {
	case "U":
		for snake[0][0]-dis < 0 {
			expandGridTail(grid, snake)
		}
		snake[0][0] -= dis
		for i := 1; i < len(snake); i++ {
			for snake[i][0]-snake[i-1][0] > 1 {
				if snake[i-1][1] != snake[i][1] {
					snake[i][1] += snake[i-1][1] - snake[i][1]
				}
				snake[i][0] -= 1
				if i == len(snake)-1 {
					(*grid)[snake[i][0]][snake[i][1]] = "#"
				}
			}
		}
	case "R":
		for snake[0][1]+dis >= len(*grid) {
			expandGridTail(grid, snake)
		}
		snake[0][1] += dis
		for i := 1; i < len(snake); i++ {
			for snake[i-1][1]-snake[i][1] > 1 {
				if snake[i-1][0] != snake[i][0] {
					snake[i][0] += snake[i-1][0] - snake[i][0]
				}
				snake[i][1] += 1
				if i == len(snake)-1 {
					(*grid)[snake[i][0]][snake[i][1]] = "#"
				}
			}
		}
	case "D":
		for snake[0][0]+dis >= len(*grid) {
			expandGridTail(grid, snake)
		}
		snake[0][0] += dis
		for i := 1; i < len(snake); i++ {
			for snake[i-1][0]-snake[i][0] > 1 {
				if snake[i-1][1] != snake[i][1] {
					snake[i][1] += snake[i-1][1] - snake[i][1]
				}
				snake[i][0] += 1
				if i == len(snake)-1 {
					(*grid)[snake[i][0]][snake[i][1]] = "#"
				}
			}
		}
	case "L":
		for snake[0][1]-dis < 0 {
			expandGridTail(grid, snake)
		}
		snake[0][1] -= dis
		for i := 1; i < len(snake); i++ {
			for snake[i][1]-snake[i-1][1] > 1 {
				if snake[i-1][0] != snake[i][0] {
					snake[i][0] += snake[i-1][0] - snake[i][0]
				}
				snake[i][1] -= 1
				if i == len(snake)-1 {
					(*grid)[snake[i][0]][snake[i][1]] = "#"
				}
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

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	grid := makeGrid(2)
	snake := make([][]int, 10)
	for i := range snake {
		snake[i] = []int{0, 0}
	}
	grid[snake[9][0]][snake[9][1]] = "#"
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		dir := parts[0]
		dis, _ := strconv.Atoi(parts[1])
		moveGridTail(&grid, snake, dir, dis)
		printGridPart2(grid, snake)
	}
	return countPath(grid)
}

func main() {
	file := "example.txt"
	// Part 1: 6266
	fmt.Println(part1(file))
	// Part 2:
	fmt.Println(part2(file))
}
