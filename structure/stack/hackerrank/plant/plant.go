package plant

import "fmt"
import Stack "algorithms-in-go/structure/stack"

type Garden struct {
	pesticides []int
}

type Plant struct {
	position  int
	daysToDie int
}

func (garden *Garden) Size() int {
	return len(garden.pesticides)
}

func (garden *Garden) Read() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var height int
		fmt.Scan(&height)
		garden.pesticides = append(garden.pesticides, height)
	}
}

func (garden *Garden) GetAliveDay() int {
	plants := garden.pesticides
	stack := Stack.NewStack()

	maxDaysToDie := 0

	for i, plant := range plants {
		daysToDie := 1
		stop := false

		for !stack.IsEmpty() && !stop {
			top, _ := stack.Top().(Plant)
			if plants[top.position] >= plant {
				daysToDie = map[bool]int{true: daysToDie, false: top.daysToDie + 1}[daysToDie > top.daysToDie+1]
				stack.Pop()
			} else {
				stop = true
			}
		}

		if stack.IsEmpty() {
			daysToDie = 0
		}
		if maxDaysToDie < daysToDie {
			maxDaysToDie = daysToDie
		}

		stack.Push(Plant{position: i, daysToDie: daysToDie})
	}

	return maxDaysToDie
}

func main() {
	garden := Garden{}
	garden.Read()
	fmt.Println(garden.GetAliveDay())
}
