package day5

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/dketterer/advent-of-code-2022/pkg/stack"
)

//go:embed input.txt
var input string

const stacks = 9

type move struct {
	quantity int
	from     int
	to       int
}

func (m *move) parse(line string) {
	expr := `.+\s(\d+).+(\d).+(\d)`
	r := regexp.MustCompile(expr)
	match := r.FindStringSubmatch(line)
	m.quantity, _ = strconv.Atoi(match[1])
	m.from, _ = strconv.Atoi(match[2])
	m.from -= 1
	m.to, _ = strconv.Atoi(match[3])
	m.to -= 1
}

type D struct {
	stacks      []stack.Stack[rune]
	stacksPart2 []stack.Stack[rune]
	moves       []move
}

func (d *D) parseStackLine(line string) {
	for i, char := range line {
		if unicode.IsUpper(char) {
			index := int((i - 1) / 4)
			d.stacks[index].Push(char)
			d.stacksPart2[index].Push(char)
		}
	}
}

func (d *D) Name() string {
	return "--- Day 5: Supply Stacks ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")

	d.stacks = make([]stack.Stack[rune], stacks)
	d.stacksPart2 = make([]stack.Stack[rune], stacks)
	for i, _ := range d.stacks {
		d.stacks[i] = stack.New[rune]()
		d.stacksPart2[i] = stack.New[rune]()
	}

	var stackLines []string
	var moveLines []string
	var secondPart bool
	for _, line := range lines {
		if !secondPart {
			if len(line) > 0 {
				stackLines = append(stackLines, line)
			} else {
				secondPart = true
			}
		} else {
			moveLines = append(moveLines, line)
		}
	}

	for i := len(stackLines) - 2; i > -1; i-- {
		d.parseStackLine(stackLines[i])
	}

	d.moves = make([]move, len(moveLines))
	for i, line := range moveLines {
		d.moves[i].parse(line)
	}

	for _, m := range d.moves {
		for i := 0; i < m.quantity; i++ {
			char, _ := d.stacks[m.from].Pop()
			d.stacks[m.to].Push(char)
		}
	}

	var result string

	for _, s := range d.stacks {
		top, _ := s.Pop()
		result = fmt.Sprintf("%s%c", result, top)
	}

	return result, err
}

func (d *D) Part2() (string, error) {
	for _, m := range d.moves {
		temp := stack.New[rune]()
		for i := 0; i < m.quantity; i++ {
			char, _ := d.stacksPart2[m.from].Pop()
			temp.Push(char)
		}
		for i := 0; i < m.quantity; i++ {
			char, _ := temp.Pop()
			d.stacksPart2[m.to].Push(char)
		}
	}

	var result string

	for _, s := range d.stacksPart2 {
		top, _ := s.Pop()
		result = fmt.Sprintf("%s%c", result, top)
	}

	return result, nil
}
