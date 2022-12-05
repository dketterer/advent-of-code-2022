package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type game struct {
	opponent int
	me       int
}

// play returns my points after playing the game
func (g *game) play() int {
	// draw
	if g.opponent == g.me {
		return g.me + 3
	}
	switch {
	case g.me == 1 && g.opponent == 3: // I win
		fallthrough
	case g.me == 2 && g.opponent == 1: // I win
		fallthrough
	case g.me == 3 && g.opponent == 2: // I win
		return g.me + 6
	default: // I loose
		return g.me
	}
}

func (g *game) print(score int) {
	fmt.Printf("%d vs %d => %d\n", g.opponent, g.me, score)
}

type D struct {
	games []*game
}

func (d *D) Name() string {
	return "--- Day 2: Rock Paper Scissors ---"
}

func (d *D) Part1() (_ string, err error) {
	lines := strings.Split(input, "\n")
	d.games = make([]*game, len(lines))

	for i, line := range lines {
		inputs := strings.Split(line, " ")
		if len(inputs) != 2 {
			return "", fmt.Errorf("got not exactly two inputs, but instead: %v\n", inputs)
		}

		me := int([]rune(inputs[1])[0])
		me -= 87
		opponent := int([]rune(inputs[0])[0])
		opponent -= 64
		d.games[i] = &game{
			opponent: opponent,
			me:       me,
		}
	}

	var totalScore int
	for _, g := range d.games {
		score := g.play()
		totalScore += score
	}

	return fmt.Sprintf("%d", totalScore), err
}

func (d *D) Part2() (string, error) {
	var totalScore int
	for _, g := range d.games {
		var myChoice int
		switch g.opponent {
		case 1: // Rock
			switch g.me {
			case 1: // Loose
				myChoice = 3
			case 2: // Draw
				myChoice = 1
			case 3: // Win
				myChoice = 2
			}
		case 2:
			switch g.me { // Paper
			case 1: // Loose
				myChoice = 1
			case 2: // Draw
				myChoice = 2
			case 3: // Win
				myChoice = 3
			}
		case 3: // Scissor
			switch g.me {
			case 1: // Loose
				myChoice = 2
			case 2: // Draw
				myChoice = 3
			case 3: // Win
				myChoice = 1
			}
		}
		g.me = myChoice
		score := g.play()
		totalScore += score
	}

	return fmt.Sprintf("%d", totalScore), nil
}
