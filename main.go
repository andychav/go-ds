package main

import (
	"fmt"

	"github.com/andychav/go-ds/stack"
)

func main() {
	s := new(stack.Stack[int])

	fmt.Println(s.Pop())

}
