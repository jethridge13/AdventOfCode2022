package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

func part1(path string) int {
	x := 1
	cycle := 0
	var ins []int = make([]int, 0)
	cyclesToCheck := []int{20, 60, 100, 140, 180, 220}
	doneCycles := make([][]int, 0)
	scanner := util.GetFileScanner(path)
	cont := scanner.Scan()
	for cont || len(ins) > 0 {
		cycle += 1
		if len(cyclesToCheck) > 0 && cycle == cyclesToCheck[0] {
			doneCycles = append(doneCycles, []int{cyclesToCheck[0], x})
			cyclesToCheck = cyclesToCheck[1:]
		}
		// If there is an instruction currently working, finish it before adding a new instruction
		if len(ins) > 0 {
			if ins[0] == cycle {
				x += ins[1]
				ins = make([]int, 0)
			}
			continue
		}
		line := scanner.Text()
		cont = scanner.Scan()
		if line == "noop" {
			continue
		}
		parts := strings.Fields(line)
		val, _ := strconv.Atoi(parts[1])
		ins = []int{cycle + 1, val}
	}
	s := 0
	for _, val := range doneCycles {
		s += val[0] * val[1]
	}
	return s
}

func part2(path string) {
	x := 1
	cycle := 0
	var ins []int = make([]int, 0)
	crt := make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
	}
	scanner := util.GetFileScanner(path)
	cont := scanner.Scan()
	for cont || len(ins) > 0 {
		crtX := cycle / 40
		crtY := cycle % 40
		if crtY == x || crtY == x-1 || crtY == x+1 {
			crt[crtX][crtY] = "#"
		} else {
			crt[crtX][crtY] = "."
		}
		cycle += 1
		// If there is an instruction currently working, finish it before adding a new instruction
		if len(ins) > 0 {
			if ins[0] == cycle {
				x += ins[1]
				ins = make([]int, 0)
			}
			continue
		}
		line := scanner.Text()
		cont = scanner.Scan()
		if line == "noop" {
			continue
		}
		parts := strings.Fields(line)
		val, _ := strconv.Atoi(parts[1])
		ins = []int{cycle + 1, val}
	}
	for _, row := range crt {
		fmt.Println(row)
	}
}

func main() {
	file := "input.txt"
	// Part 1: 12460
	fmt.Println(part1(file))
	// Part 2: EZFPRAKL
	part2(file)
}
