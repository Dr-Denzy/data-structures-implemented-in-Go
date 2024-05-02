package queue

import (
	"github.com/drdenzy/containers"
	"reflect"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	t.Run("Circular Queue of Integers", func(t *testing.T) {
		testCQ := NewCircularQueue[int](5)
		_, _ = testCQ.Enqueue(7)
		_, _ = testCQ.Enqueue(14)

		t.Run("Circular Queue: Size", func(t *testing.T) {
			if testCQ.Size() != 2 {
				t.Errorf("Circular Queue size should be 2, but got %d", testCQ.Size())
			}
		})

		t.Run("Circular Queue: Enqueue", func(t *testing.T) {
			ok, _ := testCQ.Enqueue(100)
			if !ok {
				t.Errorf("Circular Queue enqueue error: expected true, got false")
			}
			_, _ = testCQ.Enqueue(200)

			ok, _ = testCQ.Enqueue(300)
			if ok {
				t.Errorf("Circular Queue enqueue error: expected false, got true")
			}

			expected := []int{7, 14, 100, 200, 0}
			if !reflect.DeepEqual(testCQ.buffer, expected) {
				t.Errorf("Circular Queue enqueue error: expected %v, got %v", expected, testCQ.buffer)
			}

		})

		t.Run("Circular Queue: Dequeue", func(t *testing.T) {
			item, _ := testCQ.Dequeue()
			if item != 7 {
				t.Errorf("Circular Queue dequeue error: expected 7, got %v", item)
			}

		})

		t.Run("Circular Queue: Front", func(t *testing.T) {
			f, _ := testCQ.Front()
			if f != 14 {
				t.Errorf("Circular Queue front error: expected 14, got %v", f)
			}
		})

		_, _ = testCQ.Enqueue(21)

		t.Run("Circular Queue: Back", func(t *testing.T) {
			b, _ := testCQ.Back()
			if b != 21 {
				t.Errorf("Circular Queue backerror: expected 21, got %v", b)
			}
		})

		_, _ = testCQ.Enqueue(28)

		t.Run("Circular Queue: IsFull", func(t *testing.T) {
			ok := testCQ.IsFull()

			if !ok {
				t.Errorf("Circular Queue isFull error: expected true, got false")
			}

			expectedErrMsg := "queue is full"
			_, err := testCQ.Enqueue(1000)
			if !containers.ErrorContains(err, expectedErrMsg) {
				t.Errorf("Circular Queue IsFull error message: expected %v, got %v", expectedErrMsg, err)
			}
		})

		for i := 0; i < testCQ.bufferSize; i++ {
			_, _ = testCQ.Dequeue()
		}

		t.Run("Circular Queue: IsEmpty", func(t *testing.T) {
			ok := testCQ.IsEmpty()
			if !ok {
				t.Errorf("Circular Queue isEmpty error: expected true, got false")
			}

			expectedErrMsg := "queue is empty"
			_, err := testCQ.Dequeue()
			if !containers.ErrorContains(err, expectedErrMsg) {
				t.Errorf("Circular Queue IsEmpty error message: expected %v, got %v", expectedErrMsg, err)
			}
		})

		t.Run("Circular Queue: Custom Error", func(t *testing.T) {
			expectedErrMsg := "queue is empty"
			_, err := testCQ.Front()
			if !containers.ErrorContains(err, expectedErrMsg) {
				t.Errorf("Circular Queue Front() error: expected %v, got %v", expectedErrMsg, err)
			}

			_, err = testCQ.Back()
			if !containers.ErrorContains(err, expectedErrMsg) {
				t.Errorf("Circular Queue Back() error: expected %v, got %v", expectedErrMsg, err)
			}

		})

		t.Run("Circular Queue: Clear Buffer", func(t *testing.T) {
			testCQ.Clear()
			if testCQ.count != 0 {
				t.Errorf("Circular Queue clear error: expected 0, got %d", testCQ.count)
			}
		})

	})

}
