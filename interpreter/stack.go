package interpreter

import (
	"errors"
	"fmt"
)

var (
	ErrorPop               = errors.New("error: trying to pop an empty stack")
	ErrorFullStack         = errors.New("error: trying to push in a full stack")
	ErrorInvalidStackIndex = errors.New("error: invalid stack index")
	ErrorInvalidMaxSize    = errors.New("error: invalid max stack size")
)

type Stack struct {
	items   []int32
	maxSize int
}

func NewStack(maxSize int) (*Stack, error) {
	if maxSize < -1 {
		return nil, ErrorInvalidMaxSize
	}
	stack := new(Stack)
	stack.maxSize = maxSize
	return stack, nil
}

func (stack *Stack) Push(val int32) error {
	if stack.maxSize == -1 {
		stack.items = append(stack.items, val)
		return nil
	} else if len(stack.items) < stack.maxSize {
		stack.items = append(stack.items, val)
		return nil
	}
	return ErrorFullStack
	//fmt.Printf("Stack (push) -> %+v\n", stack.items)
}

func (stack *Stack) Peek() int32 {
	if stack.IsEmpty() {
		return 0
	}
	return stack.items[len(stack.items)-1]
}

func (stack *Stack) Pop() (int32, error) {
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

func (stack *Stack) Cycle() *Stack {
	var s = stack.items
	var lastPos = len(s) - 1
	var last = s[lastPos]
	copy(s[1:], s[:lastPos])
	s[0] = last
	return stack
}

func (stack *Stack) RCycle() *Stack {
	var s = stack.items
	var lastPos = len(s) - 1
	var last = s[0]
	copy(s[:lastPos], s[1:])
	s[lastPos] = last
	return stack
}

func (stack *Stack) Reverse() *Stack {
	var s = stack.items
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return stack
}

func (stack *Stack) Output() bool {
	for i := len(stack.items) - 1; i >= 0; i-- {
		fmt.Printf("%d", stack.items[i])
	}
	return true
}

func (stack *Stack) GetItemAt(i int) (int32, error) {
	if i < 0 || i > len(stack.items)-1 {
		return 0, ErrorInvalidStackIndex
	}
	return stack.items[i], nil
}

func (stack *Stack) Clear() *Stack {
	stack.items = nil
	return stack
}
