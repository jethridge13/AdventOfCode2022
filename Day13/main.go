package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Signal struct {
	list   []Signal
	number int
	t      string
}

func generatePairs(path string) [][2]any {
	packets := make([][2]any, 0)
	scanner := util.GetFileScanner(path)
	newPacket := [2]any{true, true}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		var unwrapped any
		json.Unmarshal([]byte(line), &unwrapped)
		if newPacket[0] == true {
			newPacket[0] = unwrapped
		} else {
			newPacket[1] = unwrapped
			packets = append(packets, newPacket)
			newPacket = [2]any{true, true}
		}
	}
	return packets
}

func isValidPair(left, right any) int {
	var left1, right1 []any
	leftGood, rightGood := false, false
	switch left.(type) {
	case float64:
		left1, leftGood = []any{left}, true
	case []any, []float64:
		left1 = left.([]any)
	}
	switch right.(type) {
	case float64:
		right1, rightGood = []any{right}, true
	case []any, []float64:
		right1 = right.([]any)
	}
	if leftGood && rightGood {
		return int(left1[0].(float64) - right1[0].(float64))
	}
	for i := 0; i < len(left1) && i < len(right1); i++ {
		if c := isValidPair(left1[i], right1[i]); c != 0 {
			return c
		}
	}
	return len(left1) - len(right1)
}

func part1(path string) int {
	count := 0
	pairs := generatePairs(path)
	for i, pair := range pairs {
		if isValidPair(pair[0], pair[1]) <= 0 {
			count += i + 1
		}
	}
	return count
}

func part2(path string) int {
	pairs := generatePairs(path)
	lines := make([]any, 0)
	for _, pair := range pairs {
		lines = append(lines, pair[0], pair[1])
	}
	lines = append(lines, []any{[]any{2.0}}, []any{[]any{6.0}})
	sort.Slice(lines, func(i, j int) bool {
		return isValidPair(lines[i], lines[j]) < 0
	})
	first, second := 0, 0
	for i, line := range lines {
		if fmt.Sprint(line) == "[[2]]" {
			first = i + 1
		} else if fmt.Sprint(line) == "[[6]]" {
			second = i + 1
		}
	}
	return first * second
}

func main() {
	file := "input.txt"
	// Part 1: 5185
	fmt.Println(part1(file))
	// Part 2: 23751
	fmt.Println(part2(file))
}
