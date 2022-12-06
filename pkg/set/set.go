package set

import "fmt"

type Set[I comparable] interface {
	IsEmpty() bool
	Len() int
	Push(I)
	Pop() (I, bool)
	Contains(I) bool
	Print()
	Clear()
}

type set[I comparable] map[I]bool

func New[I comparable]() Set[I] {
	m := set[I]{}
	return &m
}

func (s *set[I]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *set[I]) Len() int {
	return len(*s)
}

func (s *set[I]) Push(i I) {
	(*s)[i] = true
}

func (s *set[I]) Pop() (I, bool) {
	for key, _ := range *s {
		delete(*s, key)
		return key, true
	}
	var result I
	return result, false
}

func (s *set[I]) Contains(i I) bool {
	return (*s)[i]
}

func (s *set[I]) Print() {
	for i, _ := range *s {
		fmt.Printf("%s", i)
	}
	fmt.Println()
}

func (s *set[I]) Clear() {
	for {
		_, ok := s.Pop()
		if !ok {
			break
		}
	}
}
