package stack

import (
	"reflect"
	"testing"
)

func TestContiguousStackInt(t *testing.T) {
	testStack := NewContiguousStack[int]()

	t.Run("Contiguous Stack with elements of type int", func(t *testing.T) {
		testStack.Push(1)
		testStack.Push(2)

		t.Run("Stack Push", func(t *testing.T) {
			var stackElements []int
			for i := 0; i < 2; i++ {
				poppedElement, _ := testStack.Pop()
				stackElements = append(stackElements, poppedElement)
			}

			expectedStackElements := []int{2, 1}

			if !reflect.DeepEqual(expectedStackElements, stackElements) {
				t.Errorf("Expected: %v, Got: %v", expectedStackElements, stackElements)
			}

			// add items to stack for the next test
			testStack.Push(1)
			testStack.Push(2)
		})

		t.Run("Stack Pop", func(t *testing.T) {

			popErrorMsg := ""

			poppedItem, err := testStack.Pop()

			if !ErrorContains(err, popErrorMsg) {
				t.Errorf("Error should contain %s, but got %s", popErrorMsg, err)
			}

			if poppedItem != 2 {
				t.Errorf("Expected popped item to be 2; Got: %v", poppedItem)
			}

			if testStack.Size() != 1 {
				t.Errorf("Expected stack size to be 1; Got: %v", testStack.Size())
			}

		})

		// push more items on to the test stack for the next test
		testStack.Push(3)
		testStack.Push(4)

		t.Run("Stack Peek", func(t *testing.T) {
			peekErrorMsg := ""
			topItem, err := testStack.Peek()

			if !ErrorContains(err, peekErrorMsg) {
				t.Errorf("Error should contain %s, but got %s", peekErrorMsg, err)
			}

			if topItem != 4 {
				t.Errorf("Expected peeked item to be 4; Got: %v", topItem)
			}

		})

		t.Run("Stack Size", func(t *testing.T) {
			if testStack.Size() != 3 {
				t.Errorf("Expected stack size to be 3; Got: %v", testStack.Size())
			}
		})

		t.Run("Stack IsEmpty", func(t *testing.T) {
			if testStack.IsEmpty() == true {
				t.Errorf("Expected stack to NOT be empty: %v; Got: %v", testStack.Size(), testStack.IsEmpty())
			}
			stackSize := testStack.Size()

			for i := 0; i < stackSize; i++ {
				_, _ = testStack.Pop()
			}

			if testStack.IsEmpty() == false {
				t.Errorf("Expected stack to be empty: %v; Got: %v", testStack.Size(), testStack.IsEmpty())
			}
		})

		t.Run("Stack Clear", func(t *testing.T) {
			testStack.Push(10)
			testStack.Push(20)

			testStack.Clear()

			expectedStackElements := []int{}

			if !reflect.DeepEqual(expectedStackElements, testStack.elements) {
				t.Errorf("Expected: %v, Got: %v", expectedStackElements, testStack.elements)
			}

		})
		// clear out the test stack for the next test
		testStack.Clear()

		t.Run("Stack Custom Error", func(t *testing.T) {
			ExpectedErrorMsg := "Stack is empty. Zero value of the elements type returned"

			_, err := testStack.Pop()
			if !ErrorContains(err, ExpectedErrorMsg) {
				t.Errorf("Error should contain %s, but got %s", ExpectedErrorMsg, err)
			}

			_, err = testStack.Peek()
			if !ErrorContains(err, ExpectedErrorMsg) {
				t.Errorf("Error should contain %s, but got %s", ExpectedErrorMsg, err)
			}
		})

	})
}
