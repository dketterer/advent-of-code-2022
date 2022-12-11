package main

import (
	"fmt"

	"github.com/dketterer/advent-of-code-2022/day1"
	"github.com/dketterer/advent-of-code-2022/day2"
	"github.com/dketterer/advent-of-code-2022/day3"
	"github.com/dketterer/advent-of-code-2022/day4"
	"github.com/dketterer/advent-of-code-2022/day5"
	"github.com/dketterer/advent-of-code-2022/day6"
	"github.com/dketterer/advent-of-code-2022/day7"
	"github.com/dketterer/advent-of-code-2022/day8"
	"github.com/dketterer/advent-of-code-2022/day9"
)

func main() {
	days := []Day{
		&day1.D{},
		&day2.D{},
		&day3.D{},
		&day4.D{},
		&day5.D{},
		&day6.D{},
		&day7.D{},
		&day8.D{},
		&day9.D{},
	}
	for _, d := range days {
		fmt.Println(d.Name())
		part1, err := d.Part1()
		if err != nil {
			fmt.Printf("Part 1: ERROR: %s\n", err.Error())
		}
		fmt.Printf("Part 1: %s\n", part1)
		part2, err := d.Part2()
		if err != nil {
			fmt.Printf("Part 2: ERROR: %s\n", err.Error())
		}
		fmt.Printf("Part 2: %s\n", part2)
	}
}
