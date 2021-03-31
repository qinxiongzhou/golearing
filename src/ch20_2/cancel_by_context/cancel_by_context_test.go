package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	ctx,cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int,ctx context.Context) {
			for  {
				if isCancelled(ctx){
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i,"Done")
		}(i,ctx)
	}
	cancel()
	time.Sleep(time.Microsecond * 1000)
}


func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}