package day8

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type forest struct {
	width   int
	height  int
	field   [][]int
	visible [][]bool
}

func newForest(width, height int) *forest {
	f := &forest{
		width:  width,
		height: height,
		field:  nil,
	}
	f.field = make([][]int, height)
	f.visible = make([][]bool, height)
	for i := 0; i < height; i++ {
		f.field[i] = make([]int, width)
		f.visible[i] = make([]bool, width)
	}
	return f
}

func (f *forest) print() {
	for h := 0; h < f.height; h++ {
		for w := 0; w < f.width; w++ {
			if f.visible[h][w] {
				fmt.Printf("\033[1m%d\033[0m", f.field[h][w])
			} else {
				fmt.Printf("%d", f.field[h][w])
			}
		}
		fmt.Println()
	}
}

func (f *forest) scenicScore(targetH, targetW int) int {
	self := f.field[targetH][targetW]

	// forward
	var forwards int
	for w := targetW + 1; w < f.width; w++ {
		forwards++
		if f.field[targetH][w] >= self {
			break
		}
	}

	// backwards
	var backwards int
	for w := targetW - 1; w > -1; w-- {
		backwards++
		if f.field[targetH][w] >= self {
			break
		}
	}

	// downwards
	var downwards int
	for h := targetH + 1; h < f.height; h++ {
		downwards++
		if f.field[h][targetW] >= self {
			break
		}
	}

	// upwards
	var upwards int
	for h := targetH - 1; h > -1; h-- {
		upwards++
		if f.field[h][targetW] >= self {
			break
		}
	}
	return forwards * backwards * upwards * downwards
}

type D struct {
	forest *forest
}

func (d *D) Name() string {
	return "--- Day 8: Treetop Tree House ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")

	d.forest = newForest(len(lines), len(lines[0]))

	for h, line := range lines {
		for w, char := range line {
			digit, _ := strconv.Atoi(string(char))
			d.forest.field[h][w] = digit
		}
	}

	// rows
	for h := 0; h < d.forest.height; h++ {
		// forward
		var maxForward int
		for w := 0; w < d.forest.width; w++ {
			if w == 0 {
				d.forest.visible[h][w] = true
			} else if maxForward < d.forest.field[h][w] {
				d.forest.visible[h][w] = true
			}
			if d.forest.visible[h][w] && d.forest.field[h][w] > maxForward {
				maxForward = d.forest.field[h][w]
			}

		}
		// backward
		var maxBackward int
		for w := d.forest.width - 1; w > -1; w-- {
			if w == d.forest.width-1 {
				d.forest.visible[h][w] = true
			} else if maxBackward < d.forest.field[h][w] {
				d.forest.visible[h][w] = true
			}
			if d.forest.visible[h][w] && d.forest.field[h][w] > maxBackward {
				maxBackward = d.forest.field[h][w]
			}
		}
	}

	// columns
	for w := 0; w < d.forest.width; w++ {
		// downwards
		var maxDownwards int
		for h := 0; h < d.forest.height; h++ {
			if h == 0 {
				d.forest.visible[h][w] = true
			} else if maxDownwards < d.forest.field[h][w] {
				d.forest.visible[h][w] = true
			}
			if d.forest.visible[h][w] && d.forest.field[h][w] > maxDownwards {
				maxDownwards = d.forest.field[h][w]
			}
		}
		// upwards
		var maxUpwards int
		for h := d.forest.height - 1; h > -1; h-- {
			if h == d.forest.height-1 {
				d.forest.visible[h][w] = true
			} else if maxUpwards < d.forest.field[h][w] {
				d.forest.visible[h][w] = true
			}
			if d.forest.visible[h][w] && d.forest.field[h][w] > maxUpwards {
				maxUpwards = d.forest.field[h][w]
			}
		}
	}

	// count visible
	var sum int
	for h := 0; h < d.forest.height; h++ {
		for w := 0; w < d.forest.width; w++ {
			if d.forest.visible[h][w] {
				sum++
			}
		}
	}

	return fmt.Sprintf("%d", sum), err
}

func (d *D) Part2() (string, error) {
	var maxScenicScore int

	for h := 1; h < d.forest.height-1; h++ {
		for w := 1; w < d.forest.width-1; w++ {
			scenicScore := d.forest.scenicScore(h, w)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return fmt.Sprintf("%d", maxScenicScore), nil
}
