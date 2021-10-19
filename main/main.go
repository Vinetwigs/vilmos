package main

import (
	"fmt"

	s "github.com/vinetwigs/vilmos"
)

func main() {
	s := s.IntStack{}
	s.NewIntStack()

	s.Push(1)
	fmt.Printf("%d", s.Pop())

}
