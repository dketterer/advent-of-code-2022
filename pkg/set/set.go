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

type S[I comparable] map[I]bool

func New[I comparable]() Set[I] {
	m := S[I]{}
	return &m
}

func (s *S[I]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *S[I]) Len() int {
	return len(*s)
}

func (s *S[I]) Push(i I) {
	(*s)[i] = true
}

func (s *S[I]) Pop() (I, bool) {
	for key := range *s {
		delete(*s, key)
		return key, true
	}
	var result I
	return result, false
}

func (s *S[I]) Contains(i I) bool {
	return (*s)[i]
}

func (s *S[I]) Print() {
	for i := range *s {
		fmt.Printf("%s", i)
	}
	fmt.Println()
}

func (s *S[I]) Clear() {
	for {
		_, ok := s.Pop()
		if !ok {
			break
		}
	}
}
