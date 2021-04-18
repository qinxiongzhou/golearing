# golearing
Go 语言学习的测试案例，从入门开始
实实在在操一边，才能真正掌握。
# 1 基础

## 1.1 安装

**下载安装Go语言安装包**

https://golang.org/doc/install

https://golang.google.cn/dl

**IDE**

[Golang IDEA](https://www.jetbrains.com/go/promo/?gclid=Cj0KCQjwse-DBhC7ARIsAI8YcWLG31A2cZS1pnU1zpniXghAs1SEZux39ENFLroFpCUDszuJuIrTtVIaAn2FEALw_wcB)

**环境变量**

GOPATH=%USERPROFILE%\go;D:\workspace\golearing

这里的D:\workspace\golearing是工程的开发目录。基于Windows操作系统上的设置。

## 1.2 HelloWorld

```go
func main(){
	fmt.Println("Hello ")
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	os.Exit(0)
}
```

**应用程序的入口**
* 1、必须是main包：package main
* 2、必须是main方法：func main(){}
* 3、文件名不一定是main.go

**退出返回值**
* 1、Go中main函数不支持任何返回值
* 2、通过os.Exit来返回状态

**获取命令行参数**
* 1、main函数不支持传入参数
func main(arg []string)
* 2、在程序中直接通过os.Args获取命令行参数

详情请见：src/ch1/hello/hello_world.go
[hello_world](/src/ch1/hello/hello_world.go)

## 1.3 go build linux on windows
在windows系统上编译成linux系统上运行的文件

**在Windows系统上操作**

```shell
set GOOS=linux
go build hello.go
```
![build_windows_1.png](/images/build_windows_1.png)

![build_windows_2.png](/images/build_windows_2.png)

**在linux系统上操作**

```shell
[root@worker ~]# ll 
total 1892
-rw-r--r-- 1 root root 1937308 Mar 19 15:28 hello
[root@worker ~]# chmod 645 hello 
[root@worker ~]# ll
total 1892
-rw-r--r-x 1 root root 1937308 Mar 19 15:28 hello
[root@worker ~]# ./hello 
Hello World
```
![build_linux_1.png](/images/build_linux_1.png)

## 1.4 基本数据类型
类型 |
---- |
bool |
string|
int  int8  int16  int32  int64|
unit  unit8  unit16  unit32  unit64|
byte //alias for unit8|
rune // alias for int32,represent a Unicode code point|
float32  float64|
complex64  complex102|

## 1.5 类型转换
* 1、Go语言不允许隐式类型转换
* 2、别名和原有类型也不能进行隐式类型转换

```go
package type_test
import "testing"
type MyInt int64
func TestImplicit(t *testing.T) {
   var a int = 1
   var b int64
   //b = a //Cannot use 'a' (type int) as type int64
   b = int64(a)

   var c MyInt
   //c = b //Cannot use 'b' (type int64) as type MyInt
   c = MyInt(b)
   
   t.Log(a,b,c)
}
```

详情请见：src/ch3/type_test/type_test.go

## 1.6 指针类型
* 1、不支持指针运算（有很多C++程序员会用到指针，然后访问后续空间。这种操作是不支持的）
* 2、string是值类型，其默认的初始化值为空字符串，而不是nil

```go
func TestPoint(t *testing.T) {
   a := 1
   aPrt := &a
   //aPrt = aPrt +1 //Invalid operation: aPrt +1 (mismatched types *int and untyped int)
   t.Log(a,aPrt)
   t.Log("%T %T",a,aPrt)
}
```

```go
func TestString(t *testing.T) {
   var s string
   t.Log("*" + s + "*")
   t.Log(len(s))

   if s == "" { //注意这里是“”，而不是nil

   }
}
```

详情请见：src/ch3/type_test/type_test.go

## 1.7 用==比较数组
* 1、相同维数且含有相同个数元素的数组才可以比较
* 2、每个元素都相同的才相等

```go
func TestCompareArray(t *testing.T) {
   a := [...]int{1,2,3,4}
   b := [...]int{1,2,3,5}
   //c := [...]int{1,2,3,4,5} //c declared but not used
   d := [...]int{1,2,3,4}

   t.Log(a == b)
   //t.Log(a == c) //Invalid operation: a == c (mismatched types [4]int and [5]int)
   //t.Log(a == d) //Invalid operation: a == c (mismatched types [4]int and [5]int)
   t.Log(a == d)
}
```

详情请见：src/ch4/operator_test/operator_test.go

## 1.8 位运算符
* &^ 按位置零（自称：右1零）
* 右边运算值为0，则结果==左边值
* 右边运算值为1，则结果==0

```go
1 &^ 0 --1
1 &^ 1 --0
0 &^ 0 --0
0 &^ 1 --0
```

```go
const (
   Readable = 1 << iota
   Writable
   Executable
)

func TestBitClear(t *testing.T) {
   a := 7 //0111
   a = a &^ Readable
   t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
```
详情请见：src/ch4/operator_test/operator_test.go

## 1.9 if条件
* 1、condition表达式结果必须为布尔值
* 2、支持变量赋值
```go
if var declaration; condition{
    //code to be executed if condition is true
}
```
```go
if v,err := someFunc(); err==nil{
   t.Log("1==1",v)
}else {
   t.Log("1==1")
}
```
```go
func TestSwitchCaseCondition(t *testing.T) {
   for i := 0; i < 5; i++ {
      switch {
      case i%2 == 0:
         t.Log(fmt.Sprint(i) + " is Even")
      case i%2 == 1:
         t.Log(fmt.Sprint(i) + " is Odd")
      default:
         t.Log(fmt.Sprint(i) + " is nuknow")

      }
   }
}
```
详情请见：src/ch5/condition/condition_test.go

## 1.10 数组init和travel
**init**
```go
func TestArrayInit(t *testing.T) {
   var arr [3]int
   arr1 := [4]int{1,2,3,4}
   arr3 := [...]int{1,3,4,5} //不关注具体个数，可以用[...]表示
   t.Log(arr[1],arr[2])
   t.Log(arr1,arr3)
}
```
**travel**
```go
func TestArrayTravel(t *testing.T) {
   arr3 := [...]int{1,3,4,5}
   for i := 0;i<len(arr3);i++ {
      t.Log(arr3[i])
   }
   // idx：数组下标；e；数组值
   for idx, e := range arr3 {
      t.Log(idx,e)
   }
   //当不关注下标值时，可以用_代替
   for _, i := range arr3 {
      t.Log(i)
   }
}
```

**section**
```go
func TestArraySection(t *testing.T) {
arr3 := [...]int{1,3,4,5}
arr3_sec := arr3[1:2]
t.Log(arr3_sec) //[3]
arr3_sec = arr3[1:3]
t.Log(arr3_sec) //[3 4]
arr3_sec = arr3[1:len(arr3)]
t.Log(arr3_sec) //[3 4 5]
arr3_sec = arr3[1:]
t.Log(arr3_sec) //[3 4 5]
arr3_sec = arr3[:3]
t.Log(arr3_sec) //[1 3 4]
}
```
详情请见：src/ch6/array/array_test.go

## 1.11 切片内部结构 slice

切片：动态数组

声明的方式和数组很类似，只是中括号中无内容：例如；var s0 []int

![slice.png](/images/slice.png)

```go
func TestSliceInit(t *testing.T) {
   var s0 []int
   t.Log(len(s0), cap(s0))

   s0 = append(s0, 1)
   t.Log(len(s0), cap(s0))

   s1 := []int{1, 2, 3, 4}
   t.Log(len(s1), cap(s1))

   s2 := make([]int, 3, 5) //声明切片，len=3，cap=5
   t.Log(len(s2), cap(s2))
   t.Log(s2[0], s2[1], s2[2])
   s2 = append(s2,1)
   t.Log(s2[0], s2[1], s2[2],s2[3])
   t.Log(len(s2), cap(s2))
}
```
详情请见：src/ch6/slice/slice_test.go

# 2 并发

## 2.1 Thread vs. Groutine 协程

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

## 2.2 Mutex RWMutex 锁
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

## 2.3 WaitGroup 等待组
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

## 2.4 csp并发机制，利用channel做协程间异步通讯

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

## 2.5 select 多路选择和超时

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

## 2.6 channel的关闭

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

## 2.7 channel 广播
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

## 2.8 Context

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

## 2.9 singleton 单例模式

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

## 2.10 仅需要任意子任务完成即可结束整个任务

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


## 2.11 所有子任务完成才可结束整个任务

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

## 2.12 对象池
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

## 2.13 sync.pool 对象缓存

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

# 3 测试

## 3.1 基本测试
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

## 3.2 Benchmark

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

## 3.3 BDD Behavior Driven Development
行为驱动开发

**项目网站**

https://github.com/smartystreets/goconvey

**安装**

```shell
go get -u github.com/smartystreets/goconvey/convey
```

详情请见：src/ch36/bdd/bdd_spec_test.go

# 4 反射

## 4.1 reflect

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

  ## 4.2 reflect.DeepEqual
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
详情请见：src/ch38/flexible_reflect_test.go

## 4.3 不安全编程

```go
type MyInt int
//类型别名，可以使用
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}
```
**并发读写共享缓存，达到读和写的安全性。写的时候是块新的内存，写完成后，把读的指针通过atomic。StorePointer的原子操作重新指向一下**
```go

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data,i)
		}
		atomic.StorePointer(&shareBufPtr,unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data,*(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

详情请见：src/ch39/unsafe_pragramming/unsafe_test.go