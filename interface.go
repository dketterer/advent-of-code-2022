package main

type Day interface {
	Name() string
	Part1() (string, error)
	Part2() (string, error)
}
