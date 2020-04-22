package stack

import "errors"

type item struct {
	value string
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Push(value string) {
	newItem := item{
		value: value,
		next:  stack.top,
	}
	stack.top = &newItem
	stack.size++
}

func (stack *Stack) Pop() (string, error) {
	if stack.Size() < 1 {
		return "", errors.New("No more items")
	}

	value := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return value, nil
}
