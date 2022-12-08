package main

import (
	"fmt"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	scanner.Scan()
	line := scanner.Text()
	i := 0
	start, end := 0, -1
	chars := make(map[byte]int)
	for start < len(line) {
		end += 1
		i += 1
		chars[line[end]] += 1
		if len(chars) == 4 {
			return i
		}
		if end-start == 3 {
			chars[line[start]] -= 1
			if chars[line[start]] == 0 {
				delete(chars, line[start])
			}
			start += 1
		}
	}
	return -1
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	scanner.Scan()
	line := scanner.Text()
	i := 0
	start, end := 0, -1
	chars := make(map[byte]int)
	for start < len(line) {
		end += 1
		i += 1
		chars[line[end]] += 1
		if len(chars) == 14 {
			return i
		}
		if end-start == 13 {
			chars[line[start]] -= 1
			if chars[line[start]] == 0 {
				delete(chars, line[start])
			}
			start += 1
		}
	}
	return -1
}

func main() {
	file := "input.txt"
	// Part 1: 1598
	fmt.Println(part1(file))
	// Part 2: 2414
	fmt.Println(part2(file))
}
