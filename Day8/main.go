package main

import (
	"fmt"
	"strconv"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func loadGrid(path string) [][]int {
	scanner := util.GetFileScanner(path)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i], _ = strconv.Atoi(string(c))
		}
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getGridCount(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, tree := range row {
			if tree >= 0 {
				count += 1
			}
		}
	}
	return count
}

func calcScenicScore(heights []int) int {
	score := 1
	for _, height := range heights {
		score *= height
	}
	return score
}

func calcScore(grid [][]int, row int, column int) int {
	visGrid := make([][]int, len(grid))
	for i := range visGrid {
		blankRow := make([]int, len(grid))
		for i := range blankRow {
			blankRow[i] = -1
		}
		visGrid[i] = blankRow
	}
	visGrid[row][column] = grid[row][column]
	treeCount := make([]int, 4)
	n, s := row-1, row+1
	e, w := column+1, column-1
	max := grid[row][column]
	for n >= 0 {
		visGrid[n][column] = grid[n][column]
		treeCount[0] += 1
		if grid[n][column] >= max {
			break
		}
		n -= 1
	}
	for s < len(grid) {
		visGrid[s][column] = grid[s][column]
		treeCount[2] += 1
		if grid[s][column] >= max {
			break
		}
		s += 1
	}
	for w >= 0 {
		visGrid[row][w] = grid[row][w]
		treeCount[3] += 1
		if grid[row][w] >= max {
			break
		}
		w -= 1
	}
	for e < len(grid[row]) {
		visGrid[row][e] = grid[row][e]
		treeCount[1] += 1
		if grid[row][e] >= max {
			break
		}
		e += 1
	}
	return calcScenicScore(treeCount)
}

func part1(path string) int {
	grid := loadGrid(path)
	visGrid := make([][]int, len(grid))
	for i, row := range grid {
		// Initialize blank row to -1 because 0 is a valid tree height
		blankRow := make([]int, len(row))
		for i := range blankRow {
			blankRow[i] = -1
		}
		visGrid[i] = blankRow
		// Initialize pointers
		leftMax, rightMax := -1, -1
		lPtr, rPtr := 0, len(row)-1
		// Iterate over row
		for lPtr < len(row) {
			if row[lPtr] > leftMax {
				visGrid[i][lPtr] = row[lPtr]
				leftMax = row[lPtr]
			}
			if row[rPtr] > rightMax {
				visGrid[i][rPtr] = row[rPtr]
				rightMax = row[rPtr]
			}
			lPtr += 1
			rPtr -= 1
		}
	}
	for i := range grid[0] {
		topMax, botMax := -1, -1
		tPtr, bPtr := 0, len(grid)-1
		for tPtr < len(grid) {
			if grid[tPtr][i] > topMax {
				visGrid[tPtr][i] = grid[tPtr][i]
				topMax = grid[tPtr][i]
			}
			if grid[bPtr][i] > botMax {
				visGrid[bPtr][i] = grid[bPtr][i]
				botMax = grid[bPtr][i]
			}
			tPtr += 1
			bPtr -= 1
		}
	}
	return getGridCount(visGrid)
}

func part2(path string) int {
	maxScenicScore := 0
	grid := loadGrid(path)
	for i := range grid {
		for j := range grid[i] {
			// Skip edges because they will have a score of 0 on one side
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				continue
			}
			score := calcScore(grid, i, j)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	return maxScenicScore
}

func main() {
	file := "input.txt"
	// Part 1: 1820
	fmt.Println(part1(file))
	// Part 2: 385112
	fmt.Println(part2(file))
}
