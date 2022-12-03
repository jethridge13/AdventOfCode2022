package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		turns := strings.Fields(line)
		switch turns[1] {
		case "X":
			score += 1
			switch turns[0] {
			case "A":
				score += 3
			case "B":
			case "C":
				score += 6
			}
		case "Y":
			score += 2
			switch turns[0] {
			case "A":
				score += 6
			case "B":
				score += 3
			case "C":
			}
		case "Z":
			score += 3
			switch turns[0] {
			case "A":
			case "B":
				score += 6
			case "C":
				score += 3
			}
		}
	}
	return score
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		turns := strings.Fields(line)
		switch turns[1] {
		case "X":
			// Lose
			switch turns[0] {
			case "A":
				score += 3
			case "B":
				score += 1
			case "C":
				score += 2
			}
		case "Y":
			// Draw
			score += 3
			switch turns[0] {
			case "A":
				score += 1
			case "B":
				score += 2
			case "C":
				score += 3
			}
		case "Z":
			// Win
			score += 6
			switch turns[0] {
			case "A":
				score += 2
			case "B":
				score += 3
			case "C":
				score += 1
			}
		}
	}
	return score
}

func main() {
	file := "input.txt"
	// A, X = Rock
	// B, Y = Paper
	// C, Z = Scissors
	// Part 1: 9177
	fmt.Println(part1(file))
	// A = Rock			X = lose
	// B = Paper		Y = draw
	// C = Scissors		Z = win
	// Part 2: 12111
	fmt.Println(part2(file))
}
