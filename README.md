# golearing
Go 语言学习的测试案例，从入门开始
实实在在操一边，才能真正掌握。
# 基础

## 安装


# 并发

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
```

详情请见:src/ch19/select/select_test.go

## channel的关闭

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
详情请见:src/ch20/channel_close/channel_close_test.go

## channel 广播
利用channel的close，来向每个groutine广播消息

```go
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
```
详情请见：src/ch20_/calcel_by_close/cancel_test.go

## Context

* 根Context：通过context.Backgroud()创建
* 子Context：context.WithCancel(parentContext)创建
    * ctx,cancel := context.WithCancel(context.Background())
    
* 当前Context被取消时，基于他的子context都会被取消
* 接收取消通知<-ctx.Done()

```go
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
```
详情请见：src/ch20_2/cancel_by_context/cancel_by_context_test.go

## singleton 单例模式

```go
type Singleton struct {
}
var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}
```
详情请见：src/ch21/singleton/once_test.go

## 仅需要任意子任务完成即可结束整个任务

工作经常遇到启动多个任务，只要有一个任务完成，即可结束整个任务。

例如：向多个搜索引擎（google、百度等）发起请求，只要有一个完成，即可返回用户

```go
func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d",id)
}

func FirstResponse() string {
	numOfRunner := 10
	//use buffer channel。do not block goroutine
	ch := make(chan string,numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func() {
			ret := runTask(i)
			ch <-ret
		}()
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:",runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("After:",runtime.NumGoroutine())
}
```
详情请见：src/ch22/util_anyone_reply/first_response_test.go


## 所有子任务完成才可结束整个任务

可以使用sync.waitGroup来实现，这里用csp来实现
```go
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
```

详情请见：src/ch23/util_all_reply/all_response_test.go

## 对象池
当对象创建需要消耗大量资源，比如DB连接，可以使用channel，实现对象池来缓存对象
```go

type ResuableObj struct {
}

type ObjPool struct {
	bufChan chan *ResuableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResuableObj, numOfObj)

	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ResuableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret,nil
	case <-time.After(timeout):
		return nil,errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ResuableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")

	}
}
```
```go
func TestObjPool(t *testing.T) {
	pool := obj_pool.NewObjPool(10)

	for i := 0; i < 12; i++ {
		if v,err := pool.GetObj(time.Second * 1);err != nil{
			t.Error(err)
		}else {
			fmt.Printf("%T\n",v)
			if err := pool.ReleaseObj(v);err != nil{
				t.Error(err)
			}
		}
	}

}

```
详情请见：
src/ch32/obj_pool/obj_pool.go
src/ch32/obj_pool/obj_pool_test.go

## sync.pool 对象缓存

![sync.Pool](/images/syncPool.png)

**获取**
* 尝试从私有对象获取
* 私有对象不存在，尝试从当前Processor的共享池获取
* 如果当前Processor共享池也是空的，那么就尝试去其他Processor的共享池获取
* 如果所有子池都是空的，最后就用用户指定的New函数产生一个新的对象返回

**放回**
* 如果私有对象不存在则保存为私有对象
* 如果私有对象存在，放入当前Processor子池的共享池中

```go
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{New: func() interface{} {
		fmt.Println("create a new object.")
		return 100
	}}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC()//GC 会清楚sync。pool中缓存的对象
	v1,_ := pool.Get().(int)
	fmt.Println(v1)
}
```

详情请见：src/ch33/obj_cache/sync_pool_test.go

# 测试

## 基本测试
package "testing"，多个go文件都可以使用这个package，文件间的方法可以相互调用

**内置单元测试框架**

- Fail，Error：该测试失败，该测试继续，其他测试继续执行
- FailNow，Fatal：该测试失败，该测试终止，其他测试继续执行

```go
func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("error")
	fmt.Println("end")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("error")
	fmt.Println("end")
}

//结果
=== RUN   TestErrorInCode
start
functions_test.go:23: error
end
--- FAIL: TestErrorInCode (0.00s)
=== RUN   TestFatalInCode
start
functions_test.go:29: error
--- FAIL: TestFatalInCode (0.00s)
```

- 代码覆盖率

```shell
go test -v -cover
##结果
coverage: 100.0% of statements
```

- 断言
  
https://github.com/stretchr/testify

安装
```shell
go get -u github.com/stretchr/testify/assert
```
使用
```go
func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}
	for i:=0;i< len(inputs);i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i],ret)
	}
}
```
详情请见：src/ch34/unit_test/functions_test.go

## Benchmark

**基本结构**
```go
func BenchmarkName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}
```

**执行命令**
```shell
go test -bench=. -benchmem -count=1
```
- -bench=<相关benchmark测试>。Windows下使用go test命令行时，-bench=. 应该写成-bench="."
- -benchmem性能测试的时候显示测试函数的内存分配大小，内存分配次数的统计信息
- -count n,运行测试和性能多少此，默认一次

```shell
# 查看相关参数细节
go help testflag
```

**测试结果**
```shell
BenchmarkConcatStringByAdd-4             6087056               192.7 ns/op            16 B/op          4 allocs/op
BenchmarkConcatStringByBytesBuffer-4    10425752               117.0 ns/op            64 B/op          1 allocs/op
```
- BenchmarkConcatStringByAdd-4 表示测试的函数名，-4表示GOMAXPROCS(线程数)的值为4
- 6087056 表示一共执行了6087056次，即b.N的值
- 192.7 ns/op 表示平均每次操作花费了192.7纳秒
- 16 B/op 表示每次操作申请了16Byte的内存申请
- 4 allocs/op 表示每次操作申请了4次内存

详情请见：src/ch35/benchmark/concat_string_test.go

## BDD Behavior Driven Development
行为驱动开发

**项目网站**

https://github.com/smartystreets/goconvey

**安装**

```shell
go get -u github.com/smartystreets/goconvey/convey
```

详情请见：src/ch36/bdd/bdd_spec_test.go

# 反射

## reflect

- reflect.TypeOf返回类型（reflect.Type）:reflect.TypeOf(f)
- reflect.ValueOf返回值（reflect.Value）:reflect.ValueOf(f)
- 可以从reflect.ValueOf获得类型:reflect.ValueOf(f).Type()
- 通过kind来判断类型

```go
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown",t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}
```

```go
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f),reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}
```

```go

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{
		EmployeeID: "1",
		Name:       "Mike",
		Age:        30,
	}

	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get Name field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}

```

## reflect.DeepEqual
对象深入比较

```go
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	//fmt.Println(a == b) //invalid operation: a == b (map can only be compared to nil)
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?",reflect.DeepEqual(s1,s2))
	t.Log("s2 == s3?",reflect.DeepEqual(s2,s3))
}
```