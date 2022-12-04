package main

import (
	"fmt"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func getLetterScore(letter rune) int {
	if letter >= 65 && letter <= 90 {
		return int(letter) - 38
	}
	return int(letter) - 96
}

func calcLine(line string) int {
	m := make(map[rune]bool)
	var r rune
	for i, c := range line {
		if i < (len(line) / 2) {
			m[c] = true
		} else {
			if _, ok := m[c]; ok {
				r = c
			}
		}
	}
	return getLetterScore(r)
}

func getBadge(groups []string) rune {
	m := make(map[rune]int)
	for _, group := range groups {
		count := make(map[rune]bool)
		for _, c := range group {
			if !count[c] {
				m[c] += 1
				count[c] = true
				if m[c] == len(groups) {
					return c
				}
			}
		}
	}
	panic("AAAAAA")
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += calcLine(line)
	}
	return count
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	groups := make([]string, 0)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		groups = append(groups, line)
		if len(groups) == 3 {
			badge := getBadge(groups)
			count += getLetterScore(badge)
			groups = make([]string, 0)
		}
	}
	return count
}

func main() {
	file := "input.txt"
	// Part 1: 8394
	fmt.Println(part1(file))
	// Part 2: 2413
	fmt.Println(part2(file))
}
