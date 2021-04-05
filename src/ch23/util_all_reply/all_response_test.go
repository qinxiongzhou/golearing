package util_all_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	//use buffer channel。do not block goroutine
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func() {
			ret := runTask(i)
			ch <- ret
		}()
	}

	resultSet := ""
	for i := 0; i < numOfRunner; i++ {
		resultSet += <-ch + "\n"
	}
	return resultSet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After:", runtime.NumGoroutine())
}

/*result1: use channel
first_response_test.go:28: Before: 2
first_response_test.go:29: The result is from 9
first_response_test.go:30: After: 11

result2: use buffer channel
first_response_test.go:29: Before: 2
first_response_test.go:30: The result is from 10
first_response_test.go:31: After: 2*/
