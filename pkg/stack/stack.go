package stack

import "fmt"

type Stack[I any] interface {
	IsEmpty() bool
	Push(i I)
	Pop() (I, bool)
	Print()
}

type stack[I any] []I

func New[I any]() Stack[I] {
	return &stack[I]{}
}

func (s *stack[I]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack[I]) Push(i I) {
	*s = append(*s, i)
}

func (s *stack[I]) Pop() (I, bool) {
	if s.IsEmpty() {
		var result I
		return result, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *stack[I]) Print() {
	for _, char := range *s {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}
