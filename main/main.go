package main

import (
	"fmt"
	stack "vilmos/stack"
)

func main() {
	s := stack.NewIntStack()

	s.Push(1)
	item, _ := s.Pop()
	fmt.Printf("%d", item)

}
