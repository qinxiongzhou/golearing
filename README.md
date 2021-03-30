# golearing
Go 语言学习的测试案例，从入门开始
实实在在操一边，才能真正掌握。

## Thread vs. Groutine 协程
#### 1. 创建时默认的stack的大小
* JDK5以后Java Thread stack默认为1M
* Groutine 的stack初始化大小为2K
#### 2.和KSE（Kernel Space Entity）的对应关系
* Java Thread 是1：1
* Groutine 是M：N

```go
for i := 0; i < 10; i++ {
    go func (j int){
        fmt.Println(j)
    }(i)
}
```
详情请见:src/ch16/groutine/groutine_test.go

## Mutex RWMutex 锁
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

## WaitGroup
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