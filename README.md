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

## WaitGroup
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