package custom_types

import (
	"bytes"
	"errors"
	"fmt"
	"searchAlgorithms"
)

type Stack[T searchAlgorithms.Comparable[T]] struct {
	Capacity uint64
	Tail     uint64
	List     []T
}

func NewStack[T searchAlgorithms.Comparable[T]](capacity uint64) *Stack[T] {
	return &Stack[T]{
		Capacity: capacity,
		Tail:     0,
		List:     make([]T, capacity),
	}
}

func (stack Stack[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	for index, value := range stack.List {
		buffer.WriteString(fmt.Sprintf("%v", value))
		if index != len(stack.List)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func (stack *Stack[T]) Push(value T) error {
	if stack.Tail >= stack.Capacity {
		return errors.New("Stack is full...")
	}
	stack.List[stack.Tail] = value
	stack.Tail++
	return nil
}

func (stack *Stack[T]) Pop() (any, error) {
	if stack.Tail == 0 {
		return nil, errors.New("Stack is empty...")
	}
	stack.Tail--
	value := stack.List[stack.Tail]
	return value, nil
}
