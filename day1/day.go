package day1

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

type D struct {
}

func (d *D) Name() string {
	return "--- Day 1:  ---"
}

func (d *D) Part1() (_ int, err error) {
	lines := strings.Split(input, "\n")
	return len(lines), err
}

func (d *D) Part2() (int, error) {
	return 0, nil
}
