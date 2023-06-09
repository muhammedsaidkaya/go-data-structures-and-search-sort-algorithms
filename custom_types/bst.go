package custom_types

import (
	"bytes"
	"fmt"
	"searchAlgorithms"
)

type Node[T searchAlgorithms.Comparable[T]] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (node *Node[T]) put(x *Node[T], value T) *Node[T] {
	if x == nil {
		return &Node[T]{Value: value}
	}
	cmp := value.CompareTo(x.Value)
	if cmp < 0 {
		x.Left = node.put(x.Left, value)
	} else if cmp > 0 {
		x.Right = node.put(x.Right, value)
	} else {
		x.Value = value
	}
	return x
}

func hibbardDeletion[T searchAlgorithms.Comparable[T]](x *Node[T], value T) *Node[T] {
	if x == nil {
		return nil
	}
	cmp := value.CompareTo(x.Value)
	if cmp < 0 {
		x.Left = hibbardDeletion(x.Left, value)
	} else if cmp > 0 {
		x.Right = hibbardDeletion(x.Right, value)
	} else {
		if x.Right == nil {
			return x.Left
		} else {
			t := x
			x = minNode(t.Right)
			x.Right = deleteMin(t.Right)
			x.Left = t.Left
		}
	}
	return x
}

func minNode[T searchAlgorithms.Comparable[T]](x *Node[T]) *Node[T] {
	if x == nil {
		return nil
	} else if x.Left == nil {
		return x
	} else {
		return minNode(x.Left)
	}
}

func deleteMin[T searchAlgorithms.Comparable[T]](x *Node[T]) *Node[T] {
	if x == nil {
		return nil
	} else if x.Left == nil {
		return x.Right
	} else {
		x.Left = deleteMin(x.Left)
	}
	return x
}

type BST[T searchAlgorithms.Comparable[T]] struct {
	Root *Node[T]
}

func (bst *BST[T]) Put(value T) {
	if bst.Root == nil {
		bst.Root = &Node[T]{Value: value}
	} else {
		bst.Root = bst.Root.put(bst.Root, value)
	}
}

func (bst *BST[T]) Get(value T) *Node[T] {
	x := bst.Root
	for x != nil {
		cmp := value.CompareTo(x.Value)
		if cmp < 0 {
			x = x.Left
		} else if cmp > 0 {
			x = x.Right
		} else {
			return x
		}
	}
	return nil
}

func (bst *BST[T]) Delete(value T) {
	bst.Root = hibbardDeletion(bst.Root, value)
}

func (bst BST[T]) InOrder(x *Node[T], list []T) []T {
	if x == nil {
		return list
	}
	list = bst.InOrder(x.Left, list)
	list = append(list, x.Value)
	list = bst.InOrder(x.Right, list)
	return list
}

func (bst BST[T]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	var list []T
	list = bst.InOrder(bst.Root, list)
	for index, value := range list {
		buffer.WriteString(fmt.Sprintf("%v", value))
		if index != len(list)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}
