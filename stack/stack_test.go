package stack

import (
	"testing"
)

func TestPop(t *testing.T) {
	//return top of stack
	s := new(Stack[int])
	s.Push(1)
	s.Push(2)
	s.Push(3)
	x, _ := s.Pop()
	top, _ := s.Top()

	if x != 3 && top == 2 {
		t.Fatalf("Pop expected 2 but got %#v", x)
	}
}

func TestPush(t *testing.T) {
	s := new(Stack[int])
	s.Push(1)
	top, _ := s.Top()
	if top != 1 {
		t.Fatalf("Push expcted 1 but got %#v", top)
	}
}
