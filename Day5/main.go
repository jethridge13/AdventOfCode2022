package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func parseIns(line string) []int {
	words := strings.Fields(line)
	ins := make([]int, 3)
	ins[0], _ = strconv.Atoi(words[1])
	ins[1], _ = strconv.Atoi(words[3])
	ins[2], _ = strconv.Atoi(words[5])
	return ins
}

func parseStacks(path string) ([][]string, [][]int) {
	scanner := util.GetFileScanner(path)
	stacks := make([][]string, 0)
	ins := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// If on the linebreak between stack and instructions, stop
		if len(line) == 0 {
			break
		}
		// Add number of stacks to stack slice
		if len(stacks) == 0 {
			for i := 0; i <= len(line)/4; i++ {
				stacks = append(stacks, make([]string, 0))
			}
		}
		// Break line into chunks of 4 and add to stacks
		for i := 0; i <= len(line)/4; i++ {
			c := line[1+(4*i)]
			// Check if character is valid
			if c >= 65 && c <= 90 {
				stacks[i] = append([]string{string(c)}, stacks[i]...)
			}
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		ins = append(ins, parseIns(line))
	}
	return stacks, ins
}

func act(stacks [][]string, ins []int) {
	for i := 0; i < ins[0]; i++ {
		// Get from, to, and item to move
		from := stacks[ins[1]-1]
		to := stacks[ins[2]-1]
		item := from[len(from)-1]
		// Add to to, remove from from
		to = append(to, item)
		from = from[:len(from)-1]
		// Readd to stacks
		stacks[ins[1]-1] = from
		stacks[ins[2]-1] = to
	}
}

func act2(stacks [][]string, ins []int) {
	// Get from, to, and items to move
	from := stacks[ins[1]-1]
	to := stacks[ins[2]-1]
	items := from[len(from)-ins[0]:]
	// Add to to, remove from from
	to = append(to, items...)
	from = from[:len(from)-ins[0]]
	// Readd to stacks
	stacks[ins[1]-1] = from
	stacks[ins[2]-1] = to
}

func part1(path string) [][]string {
	stacks, ins := parseStacks(path)
	for i := range ins {
		act(stacks, ins[i])
	}
	return stacks
}

func part2(path string) [][]string {
	stacks, ins := parseStacks(path)
	for i := range ins {
		act2(stacks, ins[i])
	}
	return stacks
}

func main() {
	file := "input.txt"
	// Part 1: WHTLRMZRC
	fmt.Println((part1(file)))
	// Part 2: GMPMLWNMG
	fmt.Println(part2(file))
}
