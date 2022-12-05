package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type span struct {
	start int
	end   int
}

func newSpan(str string) span {
	nums := strings.Split(str, "-")
	start, _ := strconv.Atoi(nums[0])
	end, _ := strconv.Atoi(nums[1])
	return span{
		start: start,
		end:   end,
	}
}

func (s *span) fullyCovers(other span) bool {
	return s.start <= other.start && other.start <= s.end &&
		s.start <= other.end && other.end <= s.end
}

func (s *span) overlap(other span) bool {
	return (s.start <= other.start && other.start <= s.end) ||
		(s.start <= other.end && other.end <= s.end)
}

type pair struct {
	elf1 span
	elf2 span
}

func newPair(line string) *pair {
	parts := strings.Split(line, ",")
	return &pair{
		elf1: newSpan(parts[0]),
		elf2: newSpan(parts[1]),
	}
}

type D struct {
	pairs []*pair
}

func (d *D) Name() string {
	return "--- Day 4: Camp Cleanup ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")
	d.pairs = make([]*pair, len(lines))

	for i, line := range lines {
		d.pairs[i] = newPair(line)
	}
	var totalScore int

	for _, p := range d.pairs {
		if p.elf2.fullyCovers(p.elf1) {
			totalScore++
		} else if p.elf1.fullyCovers(p.elf2) {
			totalScore++
		}
	}

	return fmt.Sprintf("%d", totalScore), err
}

func (d *D) Part2() (string, error) {
	var totalScore int

	for _, p := range d.pairs {
		if p.elf2.overlap(p.elf1) {
			totalScore++
		} else if p.elf1.overlap(p.elf2) {
			totalScore++
		}
	}

	return fmt.Sprintf("%d", totalScore), nil
}
