package queue

import (
	"fmt"
	"github.com/drdenzy/containers"
)

// node represents a node in the linked list used by LinkedQueue.
type node[T any] struct {
	data T        // Data stored in the node
	next *node[T] // Pointer to the next node in the linked list
}

// LinkedQueue represents a queue implemented using a linked list.
type LinkedQueue[T any] struct {
	frontPtr *node[T] // Pointer to the front (head) of the queue
	rearPtr  *node[T] // Pointer to the rear (tail) of the queue
	count    int      // Number of elements in the queue
}

// NewLinkedQueue creates a new empty LinkedQueue.
func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

// Size returns the number of elements in the queue.
func (q *LinkedQueue[T]) Size() int {
	return q.count
}

// Clear clears the queue, removing all elements.
func (q *LinkedQueue[T]) Clear() {
	q.frontPtr, q.rearPtr = nil, nil
	q.count = 0
}

// IsEmpty checks if the queue is empty.
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.count == 0
}

// Front returns the element at the front of the queue without removing it.
func (q *LinkedQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}
	return q.frontPtr.data, nil
}

// Back returns the element at the back of the queue without removing it.
func (q *LinkedQueue[T]) Back() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}
	return q.rearPtr.data, nil
}

// Enqueue adds an element to the rear of the queue.
func (q *LinkedQueue[T]) Enqueue(item T) {
	newNode := &node[T]{data: item, next: nil}

	if q.IsEmpty() {
		q.frontPtr = newNode
		q.rearPtr = q.frontPtr
	} else {
		q.rearPtr.next = newNode
		q.rearPtr = q.rearPtr.next
	}
	q.count++
}

// Dequeue removes and returns the element at the front of the queue.
func (q *LinkedQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}
	dequeuedItem := q.frontPtr.data
	q.frontPtr = q.frontPtr.next

	// Check if it was the last node on the queue that was just dequeued and set the rear pointer to nil as well.
	if q.frontPtr == nil {
		q.rearPtr = nil
	}

	q.count--

	return dequeuedItem, nil
}

// String returns a string representation of the queue.
func (q *LinkedQueue[T]) String() string {
	var result = fmt.Sprintf("LinkedQueue:\nsize: %d\n", q.count)
	for n := q.frontPtr; n != nil; n = n.next {
		result += fmt.Sprintf("contents: %v", n.data)
	}
	return result + "\n"
}
