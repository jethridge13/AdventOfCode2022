package main

import (
	"reflect"
	"testing"
)

func TestGetTypedList(t *testing.T) {
	inp := "[1,1,3,1,1]"
	want := make([]Signal, 1)
	l := make([]Signal, 5)
	l[0] = Signal{t: "int", number: 1}
	l[1] = Signal{t: "int", number: 1}
	l[2] = Signal{t: "int", number: 3}
	l[3] = Signal{t: "int", number: 1}
	l[4] = Signal{t: "int", number: 1}
	want[0] = Signal{t: "list", list: l}
	get, _ := getTypedList(inp)
	if !reflect.DeepEqual(get, want) {
		t.Fatalf(`getTypedList(%s) = %v, want %v`, inp, get, want)
	}
}

func TestIsValidPair(t *testing.T) {
	typedPair := [2][]Signal{make([]Signal, 0), make([]Signal, 0)}
	typedPair[0], _ = getTypedList("[1,1,3,1,1]")
	typedPair[1], _ = getTypedList("[1,1,5,1,1]")
	if !isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = false, want true`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[1],[2,3,4]]")
	typedPair[1], _ = getTypedList("[[1],4]")
	if !isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = false, want true`, typedPair)
	}

	typedPair[0], _ = getTypedList("[9]")
	typedPair[1], _ = getTypedList("[[8,7,6]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[4,4],4,4]")
	typedPair[1], _ = getTypedList("[[4,4],4,4,4]")
	if !isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = false, want true`, typedPair)
	}

	typedPair[0], _ = getTypedList("[7,7,7,7]")
	typedPair[1], _ = getTypedList("[7,7,7]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[]")
	typedPair[1], _ = getTypedList("[3]")
	if !isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = false, want true`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[[]]]")
	typedPair[1], _ = getTypedList("[[]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	typedPair[1], _ = getTypedList("[1,[2,[3,[4,[5,6,0]]]],8,9]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[[[2],9,1,[2,2,4,8]],[1,8,8],9,[7,2,[7,0,1],0,[10,9,10,3]]]]")
	typedPair[1], _ = getTypedList("[[1,6,[[0,0,10,9]],[[10,6,0,2],[7,4],[2]],[[],3]],[[2,7,2],5,[[7,0,5],[8],1,[],3],2]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[10,10,10,4,[8,[8],6,[]]]]")
	typedPair[1], _ = getTypedList("[[],[[[8,2,6],0,4,10]],[8,[10,[10,4],[6]],2,7,[[8]]]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[7,2,[5],[[0,5,10],[4,10,10,6,1],5],4],[],[[2],4,8],[]]")
	typedPair[1], _ = getTypedList("[[1,3,[10,10,[3,5,8]],5,[6,[3,5,2,2],[1,2,10,2,10]]],[9,[[9,3,0,10],1,2],[[8,7],[0,3,9],[3],8],[5,[],8,[1,1,6],[]],2],[3,0]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[4,10,8,5,[[2,3],[8,0,2],[8,8,0,4,7],2]],[[5,[7],0,8]],[],[[],[[5,8],[1,6],[7]],[[],5],[4,6]]]")
	typedPair[1], _ = getTypedList("[[[[9,0,10],10]],[[7,5,1],7]]")
	if !isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[[1,[6,6,[7,3]]],[[],[],8],[2],[0,5,[[1,1,0,0,10]],8]]")
	typedPair[1], _ = getTypedList("[[[],[2,10,[10],[8,5]],[[10,6,5],[],1],[]],[[[],2,[7,1,4,8,4],7,[10,10]],1,[1],[5,8]],[],[2,5,[9,[8,2,0],8]],[[[]],1,[4,0,5],2,9]]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}

	typedPair[0], _ = getTypedList("[3,3,9,9,9]")
	typedPair[1], _ = getTypedList("[3,3,9,9]")
	if isValidPair(typedPair) {
		t.Fatalf(`isValidPair(%v) = true, want false`, typedPair)
	}
}

func TestPart1(t *testing.T) {
	want := 13
	get := part1("example.txt")
	if want != get {
		t.Fatalf(`part1("example.txt") = %d, want %d`, get, want)
	}
}
