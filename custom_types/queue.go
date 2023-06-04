package custom_types

import (
	"bytes"
	"errors"
	"fmt"
	"searchAlgorithms"
)

type Queue[T searchAlgorithms.Comparable[T]] struct {
	Head *Item[T]
	Tail *Item[T]
}

func (queue Queue[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	item := queue.Head
	for item != nil {
		buffer.WriteString(fmt.Sprintf("%v", item.Value))
		item = item.Next
		if item != nil {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

type Item[T searchAlgorithms.Comparable[T]] struct {
	Next  *Item[T]
	Value T
}

func NewQueue[T searchAlgorithms.Comparable[T]]() *Queue[T] {
	return &Queue[T]{
		Head: nil,
		Tail: nil,
	}
}

func (queue *Queue[T]) Enqueue(value T) {
	newItem := &Item[T]{Next: nil, Value: value}
	if queue.Head == nil {
		queue.Head = newItem
		queue.Tail = newItem
		return
	}
	queue.Tail.Next = newItem
}

func (queue *Queue[T]) Dequeue() (any, error) {
	if queue.Head == nil {
		return nil, errors.New("Queue is empty")
	}
	item := *queue.Head
	queue.Head = queue.Head.Next
	return item.Value, nil
}
