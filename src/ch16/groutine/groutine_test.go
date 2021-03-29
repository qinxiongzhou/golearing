package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func (j int){
			fmt.Println(j)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
