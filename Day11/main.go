package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type Monkey struct {
	items        []int
	operation    []string
	test         int
	trueThrow    int
	falseThrow   int
	inspectCount int
}

func generateMonkeys(path string) []Monkey {
	monkeys := make([]Monkey, 0)
	scanner := util.GetFileScanner(path)
	monkey := Monkey{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		} else if strings.HasPrefix(line, "Monkey") {
			continue
		} else if strings.HasPrefix(line, "Starting items:") {
			itemsString := line[strings.Index(line, ":")+2:]
			fields := strings.Fields(itemsString)
			items := make([]int, len(fields))
			for i, item := range fields {
				items[i], _ = strconv.Atoi(strings.Replace(item, ",", "", -1))
			}
			monkey.items = items
		} else if strings.HasPrefix(line, "Operation: ") {
			opString := line[strings.Index(line, "=")+1:]
			fields := strings.Fields(opString)
			monkey.operation = fields
		} else if strings.HasPrefix(line, "Test: ") {
			fields := strings.Fields(line)
			monkey.test, _ = strconv.Atoi(fields[len(fields)-1])
		} else if strings.HasPrefix(line, "If true:") {
			fields := strings.Fields(line)
			monkey.trueThrow, _ = strconv.Atoi(fields[len(fields)-1])
		} else if strings.HasPrefix(line, "If false:") {
			fields := strings.Fields(line)
			monkey.falseThrow, _ = strconv.Atoi(fields[len(fields)-1])
		} else {
			panic(line)
		}
	}
	monkeys = append(monkeys, monkey)
	return monkeys
}

func operate(old int, operation []string) int {
	var l int
	var r int
	if operation[0] == "old" {
		l = old
	} else {
		l, _ = strconv.Atoi(operation[0])
	}
	if operation[2] == "old" {
		r = old
	} else {
		r, _ = strconv.Atoi(operation[2])
	}
	switch operation[1] {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	default:
		panic("BAD OPERAND")
	}
}

func playMonkeys(monkeys []Monkey) {
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			// Inspect and throw item
			monkeys[i].inspectCount += 1
			item = operate(item, monkey.operation)
			item = item / 3
			var monkeyThrow int
			if item%monkey.test == 0 {
				monkeyThrow = monkey.trueThrow
			} else {
				monkeyThrow = monkey.falseThrow
			}
			monkeys[monkeyThrow].items = append(monkeys[monkeyThrow].items, item)
		}
		monkeys[i].items = make([]int, 0)
	}
}

func playMonkeys2(monkeys []Monkey, lcm int) {
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			// Inspect and throw item
			monkeys[i].inspectCount += 1
			item = item % lcm
			item = operate(item, monkey.operation)
			var monkeyThrow int
			if item%monkey.test == 0 {
				monkeyThrow = monkey.trueThrow
			} else {
				monkeyThrow = monkey.falseThrow
			}
			monkeys[monkeyThrow].items = append(monkeys[monkeyThrow].items, item)
		}
		monkeys[i].items = make([]int, 0)
	}
}

func part1(path string) int {
	monkeys := generateMonkeys(path)
	for i := 0; i < 20; i++ {
		playMonkeys(monkeys)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})
	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

func part2(path string) int {
	monkeys := generateMonkeys(path)
	lcm := 1
	for _, monkey := range monkeys {
		lcm *= monkey.test
	}
	for i := 0; i < 10000; i++ {
		playMonkeys2(monkeys, lcm)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})
	return monkeys[0].inspectCount * monkeys[1].inspectCount
}

func main() {
	file := "input.txt"
	// Part 1: 72884
	fmt.Println(part1(file))
	// Part 2: 15310845153
	fmt.Println(part2(file))
}
