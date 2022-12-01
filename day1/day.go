package day1

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type D struct {
	calories []int
}

func (d *D) Name() string {
	return "--- Day 1: Calorie Counting ---"
}

func (d *D) Part1() (_ int, err error) {
	lines := strings.Split(input, "\n")
	d.calories = []int{0}

	for _, line := range lines {
		if len(line) == 0 {
			d.calories = append(d.calories, 0)
			continue
		}
		cal, _ := strconv.Atoi(line)
		d.calories[len(d.calories)-1] += cal
	}

	sort.Ints(d.calories)

	return d.calories[len(d.calories)-1], err
}

func (d *D) Part2() (int, error) {
	var result int

	result += d.calories[len(d.calories)-1]
	result += d.calories[len(d.calories)-2]
	result += d.calories[len(d.calories)-3]

	return result, nil
}
