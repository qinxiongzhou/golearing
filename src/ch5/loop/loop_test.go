package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	// while(n<5)
	for n < 5 {
		t.Log(n)
		n++
	}
}
