package main

type Day interface {
	Name() string
	Part1() (int, error)
	Part2() (int, error)
}
