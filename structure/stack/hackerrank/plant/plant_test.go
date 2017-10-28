package plant

import (
	"testing"
)

func TestPlant(t *testing.T) {
	t1 := Garden{
		pesticides: []int{1, 2, 3, 4, 5},
	}

	if t1.GetAliveDay() != 1 {
		t.Error("Nooooooooo it should be 1 but", t1.GetAliveDay())
	}

	t2 := Garden{
		pesticides: []int{6, 5, 8, 4, 7, 10, 9},
	}

	if t2.GetAliveDay() != 2 {
		t.Error("Nooooooooo it should be 2 but get", t2.GetAliveDay())
	}

}
