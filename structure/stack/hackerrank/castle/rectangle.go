package castle

import "fmt"
import Stack "algorithms-in-go/structure/stack"

type Castle struct {
	size  int
	walls []int
}

func (rect *Castle) Read() {
	fmt.Scan(&rect.size)
	for i := 0; i < rect.size; i++ {
		var n int
		fmt.Scan(&n)
		rect.walls = append(rect.walls, n)
	}
}

func (rect *Castle) GetLargestArea() int {
	/*
	   Record the position of a wall.
	   A wall position will be popped out of the stack whenever it is greater than the next wall
	   It will loop that progress until it figures out the wall higher
	*/
	stack := Stack.NewStack()
	result := -1

	i := 0
	for i < rect.size {
		wall := rect.walls[i]
		pos, _ := stack.Top().(int)

		if stack.IsEmpty() || rect.walls[pos] <= wall {
			stack.Push(i)
			i++
		} else {
			stack.Pop()
			newPos, _ := stack.Top().(int)
			area := rect.walls[pos] * (map[bool]int{true: i, false: i - newPos - 1}[stack.IsEmpty()])
			if area > result {
				result = area
			}
		}
	}

	for !stack.IsEmpty() {
		pos, _ := stack.Top().(int)
		stack.Pop()

		newPos, _ := stack.Top().(int)
		area := rect.walls[pos] * (map[bool]int{true: i, false: i - newPos - 1}[stack.IsEmpty()])
		if area > result {
			result = area
		}
	}

	return result
}
