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
	return stack.Size() == 0
}

func (stack *Stack) Cycle() {
	var s = stack.items
	var lastPos = len(s) - 1
	var last = s[lastPos]
	copy(s[1:], s[:lastPos])
	s[0] = last
}

func (stack *Stack) RCycle() {
	var s = stack.items
	var lastPos = len(s) - 1
	var last = s[0]
	copy(s[:lastPos], s[1:])
	s[lastPos] = last
}

func (stack *Stack) Reverse() {
	var s = stack.items
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (stack *Stack) Output() {
	for i := len(stack.items) - 1; i >= 0; i-- {
		fmt.Printf("%d", stack.items[i])
	}
}
