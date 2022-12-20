package main

import (
	"fmt"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Node struct {
	n     int
	moved bool
	index *int
}

func getSequence(path string) []Node {
	seq := make([]Node, 0)
	scanner := util.GetFileScanner(path)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		n := util.ParseInt(line)
		count := i
		node := Node{n: n, index: &count}
		i += 1
		seq = append(seq, node)
	}
	return seq
}

func findZero(seq []Node) int {
	for i, n := range seq {
		if n.n == 0 {
			return i
		}
	}
	return -1
}

func mix(seq []Node) {
	i := 0
	length := len(seq)
	for i < length {
		if seq[i].moved {
			i += 1
			continue
		}
		wrap := length - 1
		newI := (i + seq[i].n) % wrap
		if newI < 0 {
			newI = wrap + newI
		}
		var move Node
		move, seq = util.SlicePop(seq, i)
		move.moved = true
		seq = util.SliceInsert(seq, newI, move)
	}
}

func mixWithOrder(seq, order []Node) {
	wrap := len(order) - 1
	for _, n := range order {
		oldI := n.index
		newI := (*oldI + n.n) % wrap
		if newI < 0 {
			newI = wrap + newI
		}

		if newI < *oldI {
			for i := newI; i < *oldI; i++ {
				*seq[i].index += 1
			}
		} else if *oldI < newI {
			for i := *oldI + 1; i <= newI; i++ {
				*seq[i].index -= 1
			}
		}

		var move Node
		move, seq = util.SlicePop(seq, *oldI)
		*oldI = newI
		seq = util.SliceInsert(seq, newI, move)
	}
	for i, n := range seq {
		n.moved = false
		seq[i] = n
	}
}

func findScore(seq []Node) int {
	index := findZero(seq)
	length := len(seq)
	return seq[(index+1000)%length].n + seq[(index+2000)%length].n + seq[(index+3000)%length].n
}

func part1(path string) int {
	seq := getSequence(path)
	mix(seq)
	return findScore(seq)
}

func part2(path string) int {
	seq := getSequence(path)
	key := 811589153
	for i, n := range seq {
		n.n *= key
		seq[i] = n
	}
	order := make([]Node, len(seq))
	copy(order, seq)
	for i := 0; i < 10; i++ {
		mixWithOrder(seq, order)
	}
	return findScore(seq)
}

func main() {
	file := "input.txt"
	// Part 1: 14888
	fmt.Println(part1(file))
	// Part 2: 3760092545849
	fmt.Println(part2(file))
}
