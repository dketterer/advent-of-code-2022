package day7

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/dketterer/advent-of-code-2022/pkg/stack"
)

//go:embed input.txt
var input string

type node struct {
	name     string
	size     int
	parent   *node
	children []*node
}

func (n *node) isFile() bool {
	return len(n.children) == 0
}

func (n *node) sumUp() int {
	if n.isFile() {
		return n.size
	} else {
		var sum int
		for _, child := range n.children {
			sum += child.sumUp()
		}
		n.size = sum
		return sum
	}
}

type Tree struct {
	root    *node
	current *node
}

func NewTree() *Tree {
	n := &node{
		name:     "/",
		size:     9,
		parent:   nil,
		children: []*node{},
	}

	t := &Tree{
		root:    n,
		current: n,
	}

	return t
}

func (t *Tree) up() bool {
	if t.current.parent != nil {
		t.current = t.current.parent
		return true
	}
	return false
}

func (t *Tree) down(name string) bool {
	for _, n := range t.current.children {
		if n.name == name {
			t.current = n
			return true
		}
	}
	return false
}

func (t *Tree) insert(name string, size int) *node {
	n := &node{
		name:     name,
		size:     size,
		parent:   t.current,
		children: []*node{},
	}
	t.current.children = append(t.current.children, n)
	return n
}

// updateSizes fills in the sizes of the dirs
func (t *Tree) updateSizes() {
	t.root.sumUp()
}

// walk is a depth-first traversal that uses iteration with a supporting stack
// it returns a node pointer which is changed each time the returned 'next()' function is called
func (t *Tree) walk() (*node, func() bool) {
	var current node

	stakk := stack.New[*node]()
	stakk.Push(t.root)

	return &current, func() bool {
		if stakk.IsEmpty() {
			return false
		} else {
			c, _ := stakk.Pop()
			current = *c
			for _, child := range current.children {
				stakk.Push(child)
			}

			return true
		}
	}
}

type D struct {
	tree *Tree
}

func (d *D) Name() string {
	return "--- Day 7: No Space Left On Device ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")
	d.tree = NewTree()
	lines = lines[1:] // skip $ cd /
	for _, line := range lines {
		if line == "$ ls" {
			continue
		} else if line == "$ cd .." {
			ok := d.tree.up()
			if !ok {
				return "", fmt.Errorf("can not go up")
			}
		} else if strings.HasPrefix(line, "$ cd") {
			name := strings.Split(line, " ")[2]
			ok := d.tree.down(name)
			if !ok {
				return "", fmt.Errorf("can not go down to dir %s", name)
			}
		} else if strings.HasPrefix(line, "dir") {
			d.tree.insert(strings.Split(line, " ")[1], 0)
		} else if unicode.IsDigit(rune(line[0])) {
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			d.tree.insert(parts[1], size)
		} else {
			return "", fmt.Errorf("unknown line type")
		}
	}

	d.tree.updateSizes()

	currentNode, next := d.tree.walk()
	sizeLimit := 100000
	var sum int
	for next() {
		if currentNode.isFile() {
			continue
		}
		if currentNode.size < sizeLimit {
			sum += currentNode.size
		}
	}

	return fmt.Sprintf("%d", sum), err
}

func (d *D) Part2() (string, error) {
	currentNode, next := d.tree.walk()
	totalCapacity := 70000000
	unusedSpace := totalCapacity - d.tree.root.size
	spaceNeeded := 30000000
	freeAtLeast := spaceNeeded - unusedSpace

	var candidates []int
	for next() {
		if currentNode.isFile() {
			continue
		}
		if currentNode.size >= freeAtLeast {
			candidates = append(candidates, currentNode.size)
		}
	}

	sort.Ints(candidates)
	result := candidates[0]

	return fmt.Sprintf("%d", result), nil
}
