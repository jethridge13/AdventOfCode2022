package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Pos struct {
	x   int
	y   int
	dir int
}

func getGraph(path string) ([][]string, []string) {
	graph := make([][]string, 0)
	var ins []string
	endOfGraph := false
	scanner := util.GetFileScanner(path)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			endOfGraph = true
			continue
		}
		if endOfGraph {
			re := regexp.MustCompile(`\d+|\w`)
			ins = re.FindAllString(line, -1)
			break
		}
		row := strings.Split(line, "")
		graph = append(graph, row)
	}
	return graph, ins
}

func printPos(graph [][]string, pos *Pos) {
	var c string
	switch pos.dir {
	case 0:
		c = ">"
	case 1:
		c = "v"
	case 2:
		c = "<"
	case 3:
		c = "^"
	}
	graph[pos.x][pos.y] = c
}

func wrapRight(graph [][]string, pos *Pos) bool {
	y := 0
	for y < len(graph[pos.x])-1 && graph[pos.x][y] == " " {
		y += 1
	}
	if graph[pos.x][y] == "." {
		pos.y = y
	} else if graph[pos.x][y] == "#" {
		return false
	}
	return true
}

func wrapDown(graph [][]string, pos *Pos) bool {
	x := 0
	for x < len(graph)-1 && graph[x][pos.y] == " " {
		x += 1
	}
	if graph[x][pos.y] == "." {
		pos.x = x
	} else if graph[x][pos.y] == "#" {
		return false
	}
	return true
}

func wrapLeft(graph [][]string, pos *Pos) bool {
	y := len(graph[pos.x]) - 1
	for y > 0 && graph[pos.x][y] == " " {
		y -= 1
	}
	if graph[pos.x][y] == "." {
		pos.y = y
	} else if graph[pos.x][y] == "#" {
		return false
	}
	return true
}

func wrapUp(graph [][]string, pos *Pos) bool {
	x := len(graph) - 1
	for x > 0 && (pos.y >= len(graph[x]) || graph[x][pos.y] == " ") {
		x -= 1
	}
	if graph[x][pos.y] == "." {
		pos.x = x
	} else if graph[x][pos.y] == "#" {
		return false
	}
	return true
}

func movePos(graph [][]string, pos *Pos, dis int) {
	for i := 0; i < dis; i++ {
		switch pos.dir {
		case 0:
			// Right
			if pos.y+1 < len(graph[pos.x]) && graph[pos.x][pos.y+1] != "" {
				switch graph[pos.x][pos.y+1] {
				case ".":
					pos.y += 1
				case "#":
					return
				default:
					if !wrapRight(graph, pos) {
						return
					}
				}
			} else {
				// At edge, wrap around if possible
				if !wrapRight(graph, pos) {
					return
				}
			}
		case 1:
			// Down
			if pos.x+1 < len(graph) && pos.y < len(graph[pos.x+1]) {
				switch graph[pos.x+1][pos.y] {
				case ".":
					pos.x += 1
				case "#":
					return
				default:
					if !wrapDown(graph, pos) {
						return
					}
				}
			} else {
				// At edge, wrap around if possible
				if !wrapDown(graph, pos) {
					return
				}
			}
		case 2:
			// Left
			if pos.y-1 >= 0 && graph[pos.x][pos.y-1] != "" {
				switch graph[pos.x][pos.y-1] {
				case ".":
					pos.y -= 1
				case "#":
					return
				default:
					if !wrapLeft(graph, pos) {
						return
					}
				}
			} else {
				// At edge, wrap around if possible
				if !wrapLeft(graph, pos) {
					return
				}
			}
		case 3:
			// Up
			if pos.x-1 >= 0 && graph[pos.x-1][pos.y] != "" {
				switch graph[pos.x-1][pos.y] {
				case ".":
					pos.x -= 1
				case "#":
					return
				default:
					if !wrapUp(graph, pos) {
						return
					}
				}
			} else {
				// At edge, wrap around if possible
				if !wrapUp(graph, pos) {
					return
				}
			}
		default:
			panic(pos.dir)
		}
		//printPos(graph, pos)
	}
}

func movePos2(graph [][]string, pos *Pos, dis int) {
	for i := 0; i < dis; i++ {
		switch pos.dir {
		case 0:
			// Right
		case 1:
			// Down
		case 2:
			// Left
		case 3:
			// Up
		default:
			panic(pos.dir)
		}
	}
}

func move(graph [][]string, pos *Pos, step string, part2 bool) {
	if step == "R" {
		dir := pos.dir
		dir += 1
		dir %= 4
		pos.dir = dir
		return
	} else if step == "L" {
		dir := pos.dir
		dir -= 1
		if dir < 0 {
			dir = 3
		}
		pos.dir = dir
		return
	}
	dis := util.ParseInt(step)
	if part2 {
		movePos2(graph, pos, dis)
	} else {
		movePos(graph, pos, dis)
	}
}

func travel(graph [][]string, ins []string, part2 bool) Pos {
	pos := Pos{}
	// Find start
	for i, space := range graph[0] {
		if space == "." {
			pos.y = i
			break
		}
	}
	for _, step := range ins {
		move(graph, &pos, step, part2)
	}
	return pos
}

func part1(path string) int {
	graph, ins := getGraph(path)
	pos := travel(graph, ins, false)
	return ((pos.x + 1) * 1000) + ((pos.y + 1) * 4) + pos.dir
}

func part2(path string) int {
	graph, ins := getGraph(path)
	pos := travel(graph, ins, true)
	return ((pos.x + 1) * 1000) + ((pos.y + 1) * 4) + pos.dir
}

func main() {
	file := "input.txt"
	// Part 1: 58248
	fmt.Println(part1(file))
	// Part 2:
}
