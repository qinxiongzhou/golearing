package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 90; i++ {
			ch <- i
			fmt.Printf("dataProducer insert data %d \n",i)
		}
		close(ch)
		wg.Done()
		fmt.Println("channel close!!!")
	}()
}

func dataReceiver(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			}else {
				fmt.Println("channel close")
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int,100)
	wg.Add(1)
	dataProducer(ch,&wg)

	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	
	wg.Wait()
}