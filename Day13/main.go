package main

import (
	"fmt"
	"strconv"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Signal struct {
	list   []Signal
	number int
	t      string
}

func generatePairs(path string) [][2]string {
	scanner := util.GetFileScanner(path)
	pairs := make([][2]string, 0)
	pair := [2]string{"", ""}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if pair[0] == "" {
			pair[0] = line
		} else {
			pair[1] = line
			pairs = append(pairs, pair)
			pair = [2]string{"", ""}
		}
	}
	return pairs
}

func getTypedList(line string) ([]Signal, int) {
	signal := Signal{t: "list"}
	l := make([]Signal, 0)
	digit := make([]byte, 0)
	i := 0
	for i < len(line) {
		c := line[i]
		s := string(c)
		if s == "[" {
			subL, newI := getTypedList(line[i+1:])
			signal.list = subL
			l = append(l, signal)
			signal = Signal{t: "list"}
			i += newI
		} else if s == "]" {
			if len(digit) > 0 {
				d, err := strconv.Atoi(string(digit))
				if err != nil {
					panic(err)
				}
				digitSignal := Signal{t: "int", number: d}
				l = append(l, digitSignal)
			}
			return l, i + 1
		} else if s == "," {
			if len(digit) > 0 {
				d, err := strconv.Atoi(string(digit))
				if err != nil {
					panic(err)
				}
				digitSignal := Signal{t: "int", number: d}
				l = append(l, digitSignal)
				digit = make([]byte, 0)
			}
		} else {
			digit = append(digit, c)
		}
		i += 1
	}
	return l, len(line)
}

func isValidPair(pair [2][]Signal) bool {
	count := 0
	for i := range pair[0] {
		count += 1
		if i >= len(pair[1]) {
			return false
		}
		if pair[0][i].t == "list" || pair[1][i].t == "list" {
			var l1 []Signal
			var l2 []Signal
			// If one is not a list, convert it to a list of length 1 and check
			if pair[0][i].t == "list" {
				l1 = pair[0][i].list
			} else {
				l1 = make([]Signal, 1)
				l1[0] = pair[0][i]
			}
			if pair[1][i].t == "list" {
				l2 = pair[1][i].list
			} else {
				l2 = make([]Signal, 1)
				l2[0] = pair[1][i]
			}
			newPair := [2][]Signal{l1, l2}
			return isValidPair(newPair)
		} else {
			if pair[0][i].number <= pair[1][i].number {
				continue
			}
			return false
		}
	}
	return true
}

func part1(path string) int {
	count := 0
	pairs := generatePairs(path)
	for i, pair := range pairs {
		typedPair := [2][]Signal{make([]Signal, 0), make([]Signal, 0)}
		typedPair[0], _ = getTypedList(pair[0])
		typedPair[1], _ = getTypedList(pair[1])
		if isValidPair(typedPair) {
			count += i + 1
		}
	}
	return count
}

func main() {
	file := "input.txt"
	// Part 1: 4714 < x < 5539
	fmt.Println(part1(file))
	// Part 2:
}
