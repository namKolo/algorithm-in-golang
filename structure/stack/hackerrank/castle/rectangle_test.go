package castle

import (
	"testing"
)

func TestBalancedBracket(t *testing.T) {
	t1 := Castle{
		size:  5,
		walls: []int{1, 2, 3, 4, 5},
	}

	if t1.GetLargestArea() != 9 {
		t.Error("Nooooooooo it should be 9")
	}

	t2 := Castle{
		size:  7,
		walls: []int{6, 2, 5, 4, 5, 1, 6},
	}

	if t2.GetLargestArea() != 12 {
		t.Error("Nooooooooo it should be 12")
	}

}
