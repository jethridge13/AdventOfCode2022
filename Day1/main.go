package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func groupByCalories(path string) []int {
	scanner := util.GetFileScanner(path)
	cals := make([]int, 0)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			cals = append(cals, count)
			count = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			count += val
		}
	}
	cals = append(cals, count)
	return cals
}

func part1(path string) int {
	cals := groupByCalories(path)
	return util.FindMax(cals)
}

func part2(path string) int {
	cals := groupByCalories(path)
	sort.Slice(cals, func(i, j int) bool {
		return cals[i] > cals[j]
	})
	return cals[0] + cals[1] + cals[2]
}

func main() {
	file := "input.txt"
	// Part 1: 71934
	fmt.Println(part1(file))
	// Part 2: 211447
	fmt.Println(part2(file))
}
