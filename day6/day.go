package day6

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dketterer/advent-of-code-2022/pkg/set"
)

//go:embed input.txt
var input string

type signal []rune

func (s signal) windowed(size int) (*[]rune, func() bool) {
	result := make([]rune, size)

	var offset int

	return &result, func() bool {
		if offset+size < len(s) {
			for i := 0; i < size; i++ {
				result[i] = s[offset+i]
			}
			offset++
			return true
		}
		return false
	}
}

func (s signal) detectMarker(windowSize int) int {
	window, next := s.windowed(windowSize)

	count := windowSize

	for next() {
		runeSet := set.New[rune]()
		for _, char := range *window {
			runeSet.Push(char)
		}
		if runeSet.Len() == windowSize {
			break
		} else {
			runeSet.Clear()
		}
		count++
	}

	return count
}

type D struct {
	signal signal
}

func (d *D) Name() string {
	return "--- Day 6: Tuning Trouble ---"
}

func (d *D) Part1() (_ string, err error) {
	d.signal = []rune(strings.Trim(input, "\n"))
	count := d.signal.detectMarker(4)
	result := fmt.Sprintf("%d", count)
	return result, err
}

func (d *D) Part2() (string, error) {
	count := d.signal.detectMarker(14)
	result := fmt.Sprintf("%d", count)
	return result, nil
}
