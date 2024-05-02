package queue

import (
	"github.com/drdenzy/containers"
	"testing"
)

func TestLinkedQueue(t *testing.T) {

	t.Run("Linked Queue with Integer Data type", func(t *testing.T) {

		t.Run("Linked Queue: Enqueue", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()
			testLQ.Enqueue(7)
			testLQ.Enqueue(14)
			testLQ.Enqueue(21)

			if testLQ.Size() != 3 {
				t.Errorf("LinkedQueue size is wrong. Expected 3, got %d", testLQ.Size())
			}

			f, _ := testLQ.Front()
			if f != 7 {
				t.Errorf("LinkedQueue front is wrong. Expected 7, got %d", f)
			}

			b, _ := testLQ.Back()
			if b != 21 {
				t.Errorf("LinkedQueue back is wrong. Expected 21, got %d", b)
			}
		})

		t.Run("Linked Queue: Dequeue", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()
			testLQ.Enqueue(7)
			testLQ.Enqueue(14)
			testLQ.Enqueue(21)

			item, _ := testLQ.Dequeue()

			if item != 7 {
				t.Errorf("LinkedQueue dequeue is wrong. Expected 7, got %d", item)
			}

			_, _ = testLQ.Dequeue()

			_, _ = testLQ.Dequeue()

			if testLQ.Size() != 0 {
				t.Errorf("LinkedQueue size is wrong. Expected 0, got %d", testLQ.Size())
			}

			expectedErrMsg := "queue is empty"

			_, err := testLQ.Dequeue()
			if !containers.ErrorContains(err, expectedErrMsg) {
				t.Errorf("LinkedQueue dequeue is wrong. Expected %s, got %s", expectedErrMsg, err)
			}
		})

		t.Run("Linked Queue: Front", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			testLQ.Enqueue(7)
			testLQ.Enqueue(14)
			_, _ = testLQ.Dequeue()

			f, _ := testLQ.Front()
			if f != 14 {
				t.Errorf("LinkedQueue front is wrong. Expected 14, got %d", f)
			}

		})

		t.Run("Linked Queue: Back", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			testLQ.Enqueue(7)
			testLQ.Enqueue(14)
			_, _ = testLQ.Dequeue()

			f, _ := testLQ.Back()
			if f != 14 {
				t.Errorf("LinkedQueue back is wrong. Expected 14, got %d", f)
			}

		})

		t.Run("Linked Queue: IsEmpty", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			testLQ.Enqueue(7)
			testLQ.Enqueue(14)

			if testLQ.IsEmpty() == true {
				t.Errorf("LinkedQueue isEmpty is wrong. Expected false, got: %t", testLQ.IsEmpty())
			}

			_, _ = testLQ.Dequeue()
			_, _ = testLQ.Dequeue()

			if !testLQ.IsEmpty() {
				t.Errorf("LinkedQueue isEmpty is wrong. Expected true, got: %t", testLQ.IsEmpty())
			}
		})

		t.Run("Linked Queue: Size", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			testLQ.Enqueue(7)
			testLQ.Enqueue(14)

			if testLQ.Size() != 2 {
				t.Errorf("LinkedQueue size is wrong. Expected 2, got %d", testLQ.Size())
			}
		})

		t.Run("Linked Queue: Clear", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			testLQ.Enqueue(7)
			testLQ.Enqueue(14)

			testLQ.Clear()

			if testLQ.Size() != 0 {
				t.Errorf("LinkedQueue size is wrong. Expected 0, got %d", testLQ.Size())
			}

		})

		t.Run("Linked Queue: custom Error", func(t *testing.T) {
			testLQ := NewLinkedQueue[int]()

			expectedErrorMsg := "queue is empty"

			_, err := testLQ.Front()

			if !containers.ErrorContains(err, expectedErrorMsg) {
				t.Errorf("LinkedQueue front is wrong. Expected %s, got %s", expectedErrorMsg, err)
			}

			_, err = testLQ.Back()

			if !containers.ErrorContains(err, expectedErrorMsg) {
				t.Errorf("LinkedQueue back is wrong. Expected %s, got %s", expectedErrorMsg, err)
			}

			_, err = testLQ.Dequeue()

			if !containers.ErrorContains(err, expectedErrorMsg) {
				t.Errorf("LinkedQueue dequeue is wrong. Expected %s, got %s", expectedErrorMsg, err)
			}

		})

	})

}
