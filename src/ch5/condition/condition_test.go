package condition

import (
	"fmt"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {

	}

	if v, err := someFunc(); err == "" {
		t.Log("1==1", v)
	} else {
		t.Log("1==1")
	}
}

func someFunc() (int, string) {
	v := 1
	err := ""
	return v, err
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(fmt.Sprint(i) + " is Even")
		case i%2 == 1:
			t.Log(fmt.Sprint(i) + " is Odd")
		default:
			t.Log(fmt.Sprint(i) + " is nuknow")

		}
	}
}
