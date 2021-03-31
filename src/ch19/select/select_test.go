package _select__test_

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 5000)
	return "Done"
}

func asyncService() chan string	{
	retCh := make(chan string,1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	// select 阻塞等待，等待的时间由case <-time.After(time.Millisecond * 500)决定，若无，则一直阻塞等待
	select {
	case ret:= <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 500):
		t.Error("timeout")
	}
	fmt.Println("final Done")
}
