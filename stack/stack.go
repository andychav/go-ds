package stack

import (
	"errors"
)

type Stack[T any] struct {
	stack []T
	size  int
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		return *new(T), errors.New("stack is empty - popping default value")
	}

	t, _ := s.Top()
	s.stack = s.stack[:s.size-1]
	s.size--
	return t, nil

}

func (s *Stack[T]) Push(t T) {
	s.stack = append(s.stack, t)
	s.size++
}

func (s *Stack[T]) Top() (T, error) {
	if s.size == 0 {
		return *new(T), errors.New("stack is empty - returning default value for top of stack ")
	}

	return s.stack[s.size-1], nil
}

func (s *Stack[T]) Size() int {
	return s.size
}
