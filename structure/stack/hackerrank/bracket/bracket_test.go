package bracket

import "testing"

func TestBalancedBracket(t *testing.T) {
	t1 := "([][{}])"
	t2 := "{[][()}]"

	b1 := bracket{t1}
	if !b1.isBalanced() {
		t.Error("t1 should be balanced")
	}

	b1 = bracket{t2}
	if b1.isBalanced() {
		t.Error("t2 should not be balanced")
	}
}
