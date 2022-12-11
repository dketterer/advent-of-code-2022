package day9

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dketterer/advent-of-code-2022/pkg/set"
)

//go:embed input.txt
var input string

type direction rune

type move struct {
	direction direction
	steps     int
}

type point struct {
	x float64
	y float64
}

type rope struct {
	knots []point
}

func newRope(numKnots int) *rope {
	knots := make([]point, numKnots)

	return &rope{
		knots: knots,
	}
}

func (r *rope) simulate(moves []move) set.Set[point] {
	visited := set.New[point]()

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			r.moveHead(m.direction)
			r.moveKnots()
			visited.Push(r.knots[len(r.knots)-1])
		}
	}
	return visited
}

func (r *rope) moveHead(m direction) {
	switch m {
	case 'R':
		r.knots[0].x++
	case 'L':
		r.knots[0].x--
	case 'U':
		r.knots[0].y++
	case 'D':
		r.knots[0].y--
	}
}

func (r *rope) moveKnots() {
	for i := 1; i < len(r.knots); i++ {
		prev := &r.knots[i-1]
		next := &r.knots[i]
		difX := prev.x - next.x
		difXAbs := math.Abs(difX)
		difY := prev.y - next.y
		difYAbs := math.Abs(difY)

		switch {
		case difXAbs+difYAbs > 2:
			next.x += math.Copysign(1, difX)
			next.y += math.Copysign(1, difY)
		case difXAbs > 1:
			next.x += math.Copysign(1, difX)
		case difYAbs > 1:
			next.y += math.Copysign(1, difY)
		}
	}
}

func printVisited(visited set.Set[point]) {
	minX := math.Inf(1)
	maxX := math.Inf(-1)
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	v := visited.(*set.S[point])

	for p := range *v {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := maxY; y > minY-1; y-- {
		for x := minX; x < maxX+1; x++ {
			if visited.Contains(point{x: x, y: y}) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

}

type D struct {
	moves []move
	rope  *rope
}

func (d *D) Name() string {
	return "--- Day 9: Rope Bridge ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")

	d.moves = make([]move, len(lines))

	expr := regexp.MustCompile("(.)\\s(\\d+)")
	for i, line := range lines {
		match := expr.FindStringSubmatch(line)
		d.moves[i].direction = direction(rune(match[1][0]))
		d.moves[i].steps, _ = strconv.Atoi(match[2])
	}

	d.rope = newRope(2)

	visited := d.rope.simulate(d.moves)

	//printVisited(visited)

	return fmt.Sprintf("%d", visited.Len()), err
}

func (d *D) Part2() (string, error) {
	d.rope = newRope(10)

	visited := d.rope.simulate(d.moves)

	//printVisited(visited)

	return fmt.Sprintf("%d", visited.Len()), nil
}
