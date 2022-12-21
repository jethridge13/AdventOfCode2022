package main

import (
	"fmt"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Monkey struct {
	name  string
	n     int
	left  *Monkey
	lName string
	op    string
	right *Monkey
	rName string
	calc  bool
}

func parseLine(line string) Monkey {
	m := Monkey{}
	fields := strings.Fields(line)
	name := fields[0][:len(fields[0])-1]
	m.name = name
	if len(fields) > 2 {
		// Example: root: pppw + sjmn
		m.lName = fields[1]
		m.op = fields[2]
		m.rName = fields[3]
	} else {
		// Example: dbpl: 5
		m.n = util.ParseInt(fields[1])
		m.calc = true
	}
	return m
}

func getMonkeys(path string, part2 bool) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)
	scanner := util.GetFileScanner(path)
	for scanner.Scan() {
		line := scanner.Text()
		m := parseLine(line)
		if part2 && m.name == "humn" {
			m.n = 0
			m.calc = false
		}
		monkeys[m.name] = &m
	}
	for name, monkey := range monkeys {
		if monkey.lName != "" {
			monkey.left = monkeys[monkey.lName]
		}
		if monkey.rName != "" {
			monkey.right = monkeys[monkey.rName]
		}
		if (monkey.left != nil && monkey.left.calc) && (monkey.right != nil && monkey.right.calc) {
			n := calc(monkey.left.n, monkey.right.n, monkey.op)
			monkey.n = n
			monkey.calc = true
		}
		monkeys[name] = monkey
	}
	return monkeys
}

func calc(left, right int, op string) int {
	var n int
	switch op {
	case "+":
		n = left + right
	case "-":
		n = left - right
	case "/":
		n = left / right
	case "*":
		n = left * right
	default:
		panic(op)
	}
	return n
}

func calcMonkey(m *Monkey) int {
	if m.calc {
		return m.n
	}
	var left, right int
	if !m.left.calc {
		left = calcMonkey(m.left)
		m.left.n = left
		m.left.calc = true
	} else {
		left = m.left.n
	}
	if !m.right.calc {
		right = calcMonkey(m.right)
		m.right.n = right
		m.right.calc = true
	} else {
		right = m.right.n
	}
	n := calc(left, right, m.op)
	m.n = n
	m.calc = true
	return n
}

func resetCalc(monkeys map[string]*Monkey) {
	for name, m := range monkeys {
		if m.left != nil || m.right != nil {
			m.n = 0
			m.calc = false
		}
		monkeys[name] = m
	}
}

func calcHuman(monkeys map[string]*Monkey) int {
	root := monkeys["root"]
	human := monkeys["humn"]
	human.calc = true
	lower, upper := 1, 145167969204648
	var n int
	// Binary search
	for lower < upper {
		n = (lower + upper) / 2
		human.n = n
		calcMonkey(root)
		diff := root.left.n - root.right.n
		if diff > 0 {
			upper = n
		} else if diff < 0 {
			lower = n
		} else {
			return n
		}
		resetCalc(monkeys)
	}
	return n
}

func part1(path string) int {
	monkeys := getMonkeys(path, false)
	return calcMonkey(monkeys["root"])
}

func part2(path string) int {
	monkeys := getMonkeys(path, true)
	return calcHuman(monkeys)
}

func main() {
	file := "input.txt"
	// Part 1: 145167969204648
	fmt.Println(part1(file))
	// Part 2: 3330805295850
	fmt.Println(part2(file))
}
