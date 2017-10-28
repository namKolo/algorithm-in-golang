package stack

import "testing"

func TestLinkedList(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.IsEmpty() {
		t.Error("Stack should not be empty")
	}
	a := stack.Pop()

	if a != 3 {
		t.Errorf("First element should be %d", 3)
	}

	if stack.Size() != 2 {
		t.Error("Size should be 2")
	}

	anotherStack := NewStack()
	anotherStack.Push(4)
	anotherStack.Push(5)

	stack.MergeStack(anotherStack)
	if stack.Size() != 4 {
		t.Error("Stack size should be 5")
	}

}
