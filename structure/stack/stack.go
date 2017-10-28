package stack

/*
  Stack is Interface for abstract data structure
*/
type Stack interface {
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
	Top() interface{}
	Iterate() <-chan interface{}
	MergeStack(stack Stack)
	GetCurrentItem() *node
}

type node struct {
	item interface{}
	next *node
}

type linkedListStack struct {
	current *node
	depth   int
}

/*
	New returns new Stack that implement interface Stacr
*/
func NewStack() Stack {
	var stack *linkedListStack = new(linkedListStack)
	stack.depth = 0
	return &linkedListStack{}
}

func (stack *linkedListStack) Push(obj interface{}) {
	stack.current = &node{item: obj, next: stack.current}
	stack.depth++
}

func (stack *linkedListStack) Pop() interface{} {
	if stack.depth > 0 {
		item := stack.current.item
		stack.current = stack.current.next
		stack.depth--
		return item
	}
	return nil
}

func (stack *linkedListStack) IsEmpty() bool {
	return stack.depth == 0
}

func (stack *linkedListStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.current.item
}

func (stack *linkedListStack) GetCurrentItem() *node {
	if stack.IsEmpty() {
		return nil
	}

	return stack.current
}

func (stack *linkedListStack) Size() int {
	return stack.depth
}

func (stack *linkedListStack) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for {
			if stack.IsEmpty() {
				break
			}
			ch <- stack.Pop()
		}
		close(ch)

	}()
	return ch
}

func (stack *linkedListStack) MergeStack(anotherStack Stack) {
	if !stack.IsEmpty() && !anotherStack.IsEmpty() {
		stack.current.next = anotherStack.GetCurrentItem()
		stack.depth = stack.depth + anotherStack.Size()
	}
}
