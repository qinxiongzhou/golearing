# golearing
Go 语言学习的测试案例，从入门开始
实实在在操一边，才能真正掌握。

## Thread vs. Groutine 协程

1. 创建时默认的stack的大小
* JDK5以后Java Thread stack默认为1M
* Groutine 的stack初始化大小为2K

2. 和KSE（Kernel Space Entity）的对应关系
* Java Thread 是1：1
* Groutine 是M：N

![kernel_switch_entity](/images/kernel_switch_entity.png)

![gmp模型](/images/gmp.png)

```go
for i := 0; i < 10; i++ {
    go func (j int){
        fmt.Println(j)
    }(i)
}
```
详情请见:src/ch16/groutine/groutine_test.go

## Mutex RWMutex 锁
* 在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
* 使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
* 在 Lock() 之前使用 Unlock() 会导致 panic 异常。
* 已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
* 在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
* 适用于读写不确定，并且只有一个读或者写的场景。

```go
func TestCounterWaitGroup(t *testing.T) {
    var mut sync.Mutex
    counter := 0
    for i := 0; i < 5000; i++ {
        go func () {
        //相当于java的finally
            defer func () {
                mut.Unlock()
            }()
            mut.Lock()
            counter++
        }()
    }
    time.Sleep(time.Microsecond * 1000)
    t.Logf("counter = %d", counter)
}
```
详情请见:src/ch17/share_mem/share_mem_test.go

## WaitGroup 等待组
主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep() 来睡眠一段时间，等待其他线程充分运行。对于简单的代码，100个for循环可以在1秒之内运行完毕，time.Sleep() 也可以达到想要的效果。

但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。这时候就不能使用time.Sleep() 来完成等待操作了。WaitGroup就可以登上舞台了。

主线程可以利用sync.WaitGroup等待goroutine都运行完毕后在执行接下来的代码。

```go
func TestCounterWaitGroup(t *testing.T) {
    var wg sync.WaitGroup
    counter := 0
        for i := 0; i < 5000; i++ {
            wg.Add(1)
            go func() {
                //相当于java的finally
                defer func() {
                    wg.Done()
                }()
            }()
        }
    wg.Wait()
}
```
详情请见:src/ch17/share_mem/share_mem_test.go

## csp并发机制，利用channel做协程间异步通讯

Communicating sequential processes

![Channel模型](/images/channel.png)

```go
func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Task is done.")
}

func asyncService() chan string	{
	//创建1个buffer的channel。这样retCh <- ret时，不用阻塞等待别的协程取走数据后再往下执行代码
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
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
```

详情请见:src/ch18/csp/async_service_test.go

## select 多路选择和超时

![select多路选择和超时](/images/select.png)

```go
func service() string {
	time.Sleep(time.Microsecond * 500)
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
	select {
	case ret:= <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 50):
		t.Error("timeout")
	}
}
```

## channel的关闭和广播

* 向关闭的channel发送数据，会导致panic
* v,ok<-ch;ok为bool值，true表示正常接受，false表示通道关闭
* 所有的channel消费者都会在channel关闭时，立刻从阻塞等待中返回上述OK值为false。
这个广播机制常被利用，进行向多个订阅者同时发送信号。如：退出信号
* 所有的channel消费者都会在channel关闭时，如果是buffer channel，channel中还有数据，返回的OK值为true；如果没有数据，返回的OK值为false


```go

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 90; i++ {
			ch <- i
			fmt.Printf("dataProducer insert data %d \n",i)
		}
		close(ch) //关闭channel
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
```













