package queue

import (
	"fmt"
	"github.com/drdenzy/containers"
)

// CircularQueue represents a fixed-size circular buffer or queue.
type CircularQueue[T any] struct {
	buffer     []T // buffer to hold the elements on the queue
	frontIndex int // front index (inclusive, i.e. first element)
	rearIndex  int // rear index (exclusive, i.e. next after last element)
	back       int // holds the index location of the last item in the buffer (queue)
	bufferSize int // capacity of the buffer (queue)
	count      int // holds the number in the buffer (queue)
}

// NewCircularQueue creates a new CircularQueue with the specified capacity.
func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		buffer:     make([]T, capacity),
		frontIndex: 0,
		rearIndex:  0,
		bufferSize: capacity,
	}
}

// String returns a string representation of the CircularQueue.
func (q *CircularQueue[T]) String() string {
	if q.IsEmpty() {
		return fmt.Sprintf("empty fixed-circular queue: %v", q.buffer)
	}
	f, _ := q.Front()
	b, _ := q.Back()
	return fmt.Sprintf("Fixed Circular Queue: %+v\n "+
		"front Index Location: %v, rear Index location: %v\n "+
		"Item at the front of the queue: %v, Item at the back of the queue: %v\n "+
		"number of items in the queue: %v, buffer capacity: %v, actual buffer capacity: %v\n",
		q.buffer, q.frontIndex, q.rearIndex, f, b, q.count, q.bufferSize, q.bufferSize-1)
}

// Clear clears the CircularQueue.
func (q *CircularQueue[T]) Clear() {
	q.buffer = make([]T, 0)
	q.count = 0
}

// Size returns the number of elements in the CircularQueue.
func (q *CircularQueue[T]) Size() int {
	return q.count
}

// IsFull checks if the CircularQueue is full.
func (q *CircularQueue[T]) IsFull() bool {
	// verify whether the queue is full or empty
	return q.frontIndex == (q.rearIndex+1)%q.bufferSize
}

// IsEmpty checks if the CircularQueue is empty.
func (q *CircularQueue[T]) IsEmpty() bool {
	return q.frontIndex == q.rearIndex
}

// Enqueue adds an element to the CircularQueue.
func (q *CircularQueue[T]) Enqueue(value T) (bool, error) {
	if q.IsFull() {
		return false, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is full. %v not enqueued", value),
			Item: value,
		}
	}
	// add item to queue
	q.buffer[q.rearIndex] = value

	// track the position of last item on the queue
	q.back = q.rearIndex

	// increment the count of number of items on the queue
	q.count++

	// set the new position of the rear index
	q.rearIndex = (q.rearIndex + 1) % q.bufferSize

	return true, nil
}

// Dequeue removes and returns an element from the front of the CircularQueue.
func (q *CircularQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}

	// get the item from the front of queue
	dequeuedItem := q.buffer[q.frontIndex]

	// decrement the count of the number of items on the queue
	q.count--

	// set the new position of frontIndex
	q.frontIndex = (q.frontIndex + 1) % q.bufferSize

	return dequeuedItem, nil
}

// Front returns the element at the front of the CircularQueue without removing it.
func (q *CircularQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}
	return q.buffer[q.frontIndex], nil
}

// Back returns the element at the back of the CircularQueue without removing it.
func (q *CircularQueue[T]) Back() (T, error) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("queue is empty. zero-value of data type %T returned", zeroValue),
			Item: zeroValue,
		}
	}
	return q.buffer[q.back], nil
}
