package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Scanner struct {
	pos    [2]int
	b      [2]int
	radius int
}

func (s Scanner) Contains(pos [2]int) bool {
	return s.radius >= util.GetManhattanDistance(s.pos, pos)
}

func pointEqual(a, b [2]int) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func parseScanner(line string) Scanner {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	fields := strings.Fields(line)
	pos := [2]int{0, 0}
	b := [2]int{0, 0}
	sX := fields[2]
	sY := fields[3]
	sXTrim := sX[2 : len(sX)-1]
	sYTrim := sY[2 : len(sY)-1]
	pos[1] = util.ParseInt(sXTrim)
	pos[0] = util.ParseInt(sYTrim)
	bX := fields[8]
	bY := fields[9]
	bXTrim := bX[2 : len(bX)-1]
	bYTrim := bY[2:]
	b[1] = util.ParseInt(bXTrim)
	b[0] = util.ParseInt(bYTrim)
	mDis := util.GetManhattanDistance(pos, b)
	return Scanner{pos: pos, b: b, radius: mDis}
}

func getScanners(path string) []Scanner {
	scanners := make([]Scanner, 0)
	scanner := util.GetFileScanner(path)
	for scanner.Scan() {
		line := scanner.Text()
		scanner := parseScanner(line)
		scanners = append(scanners, scanner)
	}
	return scanners
}

func part1(path string, row int) int {
	scanners := getScanners(path)
	covered := 0
	xMin := math.MaxInt
	xMax := 0
	for _, s := range scanners {
		if s.pos[1]-s.radius < xMin {
			xMin = s.pos[1] - s.radius
		}
		if s.pos[1]+s.radius > xMax {
			xMax = s.pos[1] + s.radius
		}
	}
	test := [2]int{row, xMin}
	for x := xMin; x <= xMax; x++ {
		test[1] = x
		for _, s := range scanners {
			if s.Contains(test) && !pointEqual(s.pos, test) && !pointEqual(s.b, test) {
				covered += 1
				break
			}
		}
	}
	return covered
}

func part2(path string, max int) int {
	scanners := getScanners(path)
	coord := [2]int{0, 0}
	cover := Scanner{}
	for {
		covered := false
		for _, s := range scanners {
			covered = s.Contains(coord)
			if covered {
				cover = s
				break
			}
		}

		if !covered {
			break
		}

		yDis := util.Abs(cover.pos[0] - coord[0])
		xDis := util.Abs(cover.pos[1] - coord[1])
		skip := cover.radius - yDis + xDis + 1

		if coord[1]+skip > max {
			coord[1] = 0
			coord[0] += 1
		} else {
			coord[1] += skip
		}
	}
	return coord[1]*4000000 + coord[0]
}

func main() {
	file := "input.txt"
	row := 2000000
	r := 4000000
	// Part 1: 4737567
	fmt.Println(part1(file, row))
	// Part 2: 13267474686239
	fmt.Println(part2(file, r))
}
