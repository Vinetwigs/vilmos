package stack

import "errors"

type IntStack struct {
	items []int
}

func NewIntStack() *IntStack {
	stack := new(IntStack)
	return stack
}

func (stack *IntStack) Push(val int) {
	stack.items = append(stack.items, val)
}

func (stack *IntStack) Pop() (int, error) {
	if len(stack.items) == 0 {
		return -1, errors.New("error: stack is empty, cannot pop")
	}

	item := stack.items[len(stack.items)-1]

	return item, nil
}

func (stack *IntStack) Size() int {
	return len(stack.items)
}

func (stack *IntStack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	} else {
		return false
	}
}
