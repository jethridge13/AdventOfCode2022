package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func getRanges(ranges []string) ([]int, []int) {
	r1s := strings.Split(ranges[0], "-")
	r2s := strings.Split(ranges[1], "-")
	var r1 = []int{}
	var r2 = []int{}
	for _, i := range r1s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		r1 = append(r1, j)
	}
	for _, i := range r2s {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		r2 = append(r2, j)
	}
	return r1, r2
}

func part1(path string) int {
	scanner := util.GetFileScanner(path)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		r1, r2 := getRanges(ranges)
		if (r1[0] >= r2[0] && r1[1] <= r2[1]) || (r2[0] >= r1[0] && r2[1] <= r1[1]) {
			count += 1
		}
	}
	return count
}

func part2(path string) int {
	scanner := util.GetFileScanner(path)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		r1, r2 := getRanges(ranges)
		if (r1[0] <= r2[1] && r1[1] >= r2[0]) ||
			(r2[0] <= r1[1] && r2[1] >= r1[0]) {
			count += 1
		}
	}
	return count
}

func main() {
	file := "input.txt"
	// Part 1: 477
	fmt.Println(part1(file))
	// Part 2: 830
	fmt.Println(part2(file))
}
