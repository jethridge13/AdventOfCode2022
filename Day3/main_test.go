package main

import "testing"

func TestGetLetterScore(t *testing.T) {
	c := 'A'
	want := 27
	get := getLetterScore(c)
	if want != get {
		t.Fatalf(`getLetterScore('A') = %d, want %d`, get, want)
	}

	c = 'Z'
	want = 52
	get = getLetterScore(c)
	if want != get {
		t.Fatalf(`getLetterScore('Z') = %d, want %d`, get, want)
	}

	c = 'a'
	want = 1
	get = getLetterScore(c)
	if want != get {
		t.Fatalf(`getLetterScore('a') = %d, want %d`, get, want)
	}

	c = 'z'
	want = 26
	get = getLetterScore(c)
	if want != get {
		t.Fatalf(`getLetterScore('z') = %d, want %d`, get, want)
	}
}

func TestCalcLine(t *testing.T) {
	line := "vJrwpWtwJgWrhcsFMMfFFhFp"
	want := getLetterScore('p')
	get := calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}

	line = "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	want = getLetterScore('L')
	get = calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}

	line = "PmmdzqPrVvPwwTWBwg"
	want = getLetterScore('P')
	get = calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}

	line = "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"
	want = getLetterScore('v')
	get = calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}

	line = "ttgJtRGJQctTZtZT"
	want = getLetterScore('t')
	get = calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}

	line = "CrZsJsPPZsGzwwsLwLmpwMDw"
	want = getLetterScore('s')
	get = calcLine(line)
	if want != get {
		t.Fatalf(`calcLine(%s) = %d, want %d`, line, get, want)
	}
}

func TestGetBadge(t *testing.T) {
	groups := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}
	want := 'r'
	get := getBadge(groups)
	if want != get {
		t.Fatalf(`getBadge(%v) = %c, want %c`, groups, get, want)
	}

	groups = []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
	want = 'Z'
	get = getBadge(groups)
	if want != get {
		t.Fatalf(`getBadge(%v) = %c, want %c`, groups, get, want)
	}
}
