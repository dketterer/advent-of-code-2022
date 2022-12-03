package day3

import (
	_ "embed"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type rucksack struct {
	compartments []rune
}

func (r rucksack) c1() []rune {
	return r.compartments[:len(r.compartments)/2]
}

func (r rucksack) c2() []rune {
	return r.compartments[len(r.compartments)/2:]
}

func newRucksack() rucksack {
	return rucksack{
		compartments: []rune{},
	}
}

type D struct {
	r []rucksack
}

func (d *D) Name() string {
	return "--- Day 3: Rucksack Reorganization ---"
}

func (d *D) Part1() (_ int, err error) {
	lines := strings.Split(input, "\n")
	d.r = make([]rucksack, len(lines))

	for i, line := range lines {
		d.r[i] = newRucksack()
		for _, char := range line {
			switch unicode.IsLower(char) {
			case true:
				char -= 96
			case false:
				char -= 64 - 26
			}
			d.r[i].compartments = append(d.r[i].compartments, char)
		}
	}
	var totalScore int
	for _, rucks := range d.r {
		runes := make(map[rune]bool)
		for _, char := range rucks.c1() {
			runes[char] = true
		}
		for _, char := range rucks.c2() {
			if contain, _ := runes[char]; contain {
				// count only once
				runes[char] = false
				totalScore += int(char)
			}
		}
	}

	return totalScore, err
}

func (d *D) Part2() (int, error) {
	var totalScore int
	runesPerGroup := make(map[rune]int)
	for i, rucks := range d.r {
		runesPerRucksack := make(map[rune]bool)
		for _, char := range rucks.compartments {
			runesPerRucksack[char] = true
		}
		for char := range runesPerRucksack {
			runesPerGroup[char]++
		}
		if i%3 == 2 {
			for char, count := range runesPerGroup {
				if count == 3 {
					totalScore += int(char)
				}
			}
			runesPerGroup = make(map[rune]int)
		}
	}

	return totalScore, nil
}
