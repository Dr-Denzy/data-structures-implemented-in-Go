package stack

import (
	"fmt"
)

type node[T any] struct {
	data T
	next *node[T]
}

type LinkedStack[T any] struct {
	top       *node[T]
	stackSize int
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) Size() int {
	return s.stackSize
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.stackSize == 0
}

func (s *LinkedStack[T]) Clear() {
	s.stackSize = 0
	s.top = nil
}

func (s *LinkedStack[T]) Push(data T) {
	s.top = &node[T]{data: data, next: s.top} // the linking happens here
	s.stackSize++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	poppedItem := s.top.data
	s.top = s.top.next
	s.stackSize--
	return poppedItem, nil
}

func (s *LinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	peekedItem := s.top.data
	return peekedItem, nil
}

func (s *LinkedStack[T]) Display() []T {
	dataFromNodes := make([]T, 0, s.stackSize)
	currentNode := s.top
	for currentNode != nil {
		dataFromNodes = append(dataFromNodes, currentNode.data)
		currentNode = currentNode.next
	}
	return dataFromNodes
}
