package main

import (
	"fmt"

	"github.com/dketterer/advent-of-code-2022/day1"
)

func main() {
	days := []Day{
		&day1.D{},
	}
	for _, d := range days {
		fmt.Println(d.Name())
		part1, err := d.Part1()
		if err != nil {
			fmt.Printf("Part 1: ERROR: %s\n", err.Error())
		}
		fmt.Printf("Part 1: %d\n", part1)
		part2, err := d.Part2()
		if err != nil {
			fmt.Printf("Part 2: ERROR: %s\n", err.Error())
		}
		fmt.Printf("Part 2: %d\n", part2)
	}
}
