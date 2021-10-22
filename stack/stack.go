package stack

import (
	"errors"
	"fmt"
)

var (
	ErrorPop = errors.New("error: trying to pop an empty stack")
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	stack := new(Stack)
	return stack
}

func (stack *Stack) Push(val int) {
	stack.items = append(stack.items, val)

	//fmt.Printf("Stack (push) -> %+v\n", stack.items)
}

func (stack *Stack) Peek() int {
	return stack.items[len(stack.items)-1]
}

func (stack *Stack) Pop() (int, error) {
	if len(stack.items) == 0 {
		return 0, ErrorPop
	}

	index := len(stack.items) - 1 // Get the index of the top most element.
	item := stack.items[index]    // Index into the slice and obtain the element.
	stack.items = stack.items[:index]

	//fmt.Printf("Stack (push) -> %+v\n", stack.items)
	return item, nil
}

func (stack *Stack) Size() int {
	return len(stack.items)
}

func (stack *Stack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Cycle() {
	var last = stack.items[len(stack.items)-1]
	//fmt.Printf("Last: %d\n", last)
	for i := len(stack.items) - 1; i > 0; i-- {
		stack.items[i] = stack.items[i-1]
		//fmt.Printf("%d) %+v\n", i, stack.items)
	}
	stack.items[0] = last
}

func (stack *Stack) RCycle() {
	var last = stack.items[0]
	//fmt.Printf("Last: %d\n", last)
	for i := 0; i < len(stack.items)-1; i++ {
		stack.items[i] = stack.items[i+1]
		//fmt.Printf("%d) %+v\n", i, stack.items)
	}
	stack.items[len(stack.items)-1] = last
}

func insertBottom(x int, s *Stack) {
	if s.IsEmpty() {
		s.Push(x)
	} else {
		y := s.Peek()
		s.Pop()
		insertBottom(x, s)
		s.Push(y)
	}
}

func (stack *Stack) Reverse() {
	if stack.Size() > 0 {
		x := stack.Peek()
		stack.Pop()
		stack.Reverse()
		insertBottom(x, stack)
	}
}

func (stack *Stack) Output() {
	for i := len(stack.items) - 1; i >= 0; i-- {
		fmt.Printf("%d", stack.items[i])
	}
}
