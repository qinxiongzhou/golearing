package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Microsecond * 1000)
	//counter = 4880,结果不正确
	t.Logf("counter = %d",counter)
}

//线程安全
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			//相当于java的finally
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Microsecond * 1000)
	t.Logf("counter = %d",counter)
}

//等待组，可以不用写时间等待函数
func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			//相当于java的finally
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d",counter)
}
