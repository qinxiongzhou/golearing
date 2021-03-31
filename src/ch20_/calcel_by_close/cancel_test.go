package calcel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	cancenChan := make(chan struct{},0)
	for i := 0; i < 5; i++ {
		go func(i int,cancelCh chan struct{}) {
			for  {
				if isCancelled(cancenChan){
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i,"Done")
		}(i,cancenChan)
	}
	/*cancel_1(cancenChan)
	cancel_1(cancenChan)*/
	cancel_all(cancenChan)
	time.Sleep(time.Microsecond * 1000)
}

//cancel one groutine
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

//cancel all groutine
func cancel_all(cancelChan chan struct{})  {
	close(cancelChan)
}

func isCancelled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}
