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

参考代码：[hello_world](/src/ch1/hello/hello_world.go)

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
参考代码：[type_test](/src/ch3/type_test/type_test.go)

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

参考代码：[type_test](/src/ch3/type_test/type_test.go)

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

参考代码：[operator_test](/src/ch4/operator_test/operator_test.go)

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
参考代码：[operator_test](/src/ch4/operator_test/operator_test.go)

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
参考代码：[condition_test](/src/ch5/condition/condition_test.go)

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
参考代码：[array_test](/src/ch6/array/array_test.go)

## 1.11 切片

**基本介绍**

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

**共享储存结构**

![slice_share_mem.png](/images/slice_share_mem.png)

```go
func TestSliceShareMemory(t *testing.T) {
   year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul",
      "Aug","Sep","Oct","Nov","Dec"}
   q2 := year[3:6]
   t.Log(q2,len(q2),cap(q2)) //[Apr May Jun] 3 9

   sumer := year[5:8]
   t.Log(sumer,len(sumer),cap(sumer)) //[Jun Jul Aug] 3 7

   sumer[0] = "Unknow"
   t.Log(q2) //[Apr May Unknow]
   t.Log(year) //[Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]

}
```
参考代码：[slice_test](/src/ch6/slice/slice_test.go)

## 1.13 数组 VS 切片
* 1、容量是否可伸缩
  - 数组不行，切片可以
* 2、是否可以进行比较
  - 数组可以，切片不行
  
## 1.14 map
**init**
```go
func TestName(t *testing.T) {
   m1 := map[int]int{1: 1, 2: 4, 3: 6}
   t.Log(m1[2]) //4
   t.Logf("len m1=%d", len(m1)) //len m1=3
   m2 := map[int]int{}
   t.Logf("len m2=%d", len(m2))//len m2=0
   m3 := make(map[int]int, 10) //用make创建map结构，10指cap，不是len
   t.Logf("len m3=%d", len(m3))//len m3=0
}
```
**access Not Exist string key**
```go
func TestAccessNotExiststringKey(t *testing.T) {
   m1 := map[int]int{}
   t.Log(m1[1])
   m1[2] = 0
   t.Log(m1[2])

   //value 不存在，返回值也是0，用如下方式区分值为0还是不存在
   if v,ok := m1[3]; ok {
      t.Logf("Key 3 value is %d" ,v)
   }else {
      t.Logf("key 3 is not existing")
   }

   m1[3] = 0
   //value 不存在，返回值也是0，用如下方式区分值为0还是不存在
   if v,ok := m1[3]; ok {
      t.Logf("Key 3 value is %d" ,v)
   }else {
      t.Logf("key 3 is not existing")
   }
}
```
**travel map  遍历map**
```go
func TestTravelMap(t *testing.T) {
   m1 := map[int]int{1: 1, 2: 4, 3: 6}
   for k, v := range m1 {
      t.Log(k, v)
   }
}
```
**value is a function，map的值是一个方法**
```go
func TestMapWithFuncValue(t *testing.T) {
   //map value is a function
   m := map[int]func(op int) int{}
   m[1] = func(op int) int {
      return op
   }
   m[2] = func(op int) int {
      return op * op
   }
   m[3] = func(op int) int {
      return op * op * op
   }
   t.Log(m[1](2), m[2](2), m[3](2))

}
```
参考代码：[mapext_test](/src/ch8/map/mapext_test.go)


## 1.15 set

Go 的内置集合中没有Set的实现，可以map[type]bool
* 1、元素的唯一性
* 2、基本操作
  - 添加元素
  - 判断元素是否存在
  - 删除元素
  - 元素个数

```go
func TestMapForSet(t *testing.T) {
   mySet := map[string]bool{}
   mySet["Go"]=true
   key := "Go"

   if mySet[key]{
      t.Log(key+" is exist")
   }else {
      t.Log(key + " is not exist")
   }
}
```
参考代码：[mapext_test](/src/ch8/map/mapext_test.go)

## 1.16 函数
* 1、可以有多个返回值
* 2、所有参数都是值传递：slice、map、channel会有传引用的错觉
* 3、函数可以作为变量的值
* 4、函数可以作为参数和返回值

```go
// 方法作为参数传递
func timeSpent(myInner func(op int) int) func(op int) int {
   return func(n int) int {
      start :=time.Now()
      ret :=myInner(n)
      fmt.Println("time spent:",time.Since(start).Seconds())
      return  ret
   }
}

func slowFun(op int) int {
   time.Sleep(time.Second*1)
   fmt.Println("slowFun")
   return op
}

func TestFn02(t *testing.T) {
   a,_ :=returnMultiValues()
   t.Log(a)
   tsSF := timeSpent(slowFun)
   t.Log(tsSF(10))
}
```
参考代码：[func_test](/src/ch10/func/func_test.go)

**可变参数**
```go
func sum(ops ...int) int {
   ret := 0
   for _, op := range ops {
      ret += op
   }
   return ret
}

func TestVarParam(t *testing.T) {
   t.Log(sum(1,2,3,5))
   t.Log(sum(1,5,3,12))
}
```
参考代码：[funcVarParam_test](/src/ch10/func/funcVarParam_test.go)

## 1.17 defer函数
被调用的方法延迟执行，一般在方法推出之前执行，可用于做资源释放

```go
func clear()  {
   fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
   defer clear()
   fmt.Println("Start")
   panic("Fatal error") //defer仍会执行
}
```
参考代码：[funcDefer_test](/src/ch10/func/funcDefer_test.go)

## 1.18 结构体

**定义**

```go
type Employee struct {
   Id string
   Name string
   Age int
}

func TestCreateEmployeeObj(t *testing.T) {
   e := Employee{"0", "Bob", 20}
   e1 := Employee{
      Name: "Mike",
      Age:  30,
   }
   e2 := new(Employee)
   e2.Id = "2"
   e2.Age = 22
   e2.Name = "Rose"

   t.Log(e)
   t.Log(e1)
   t.Log(e1.Id)
   t.Log(e2)
   t.Logf("e is %T",e) //e is encapsulation.Employee
   t.Logf("e2 is %T",e2) //e2 is *encapsulation.Employee
}
```

**行为（方法）定义**
```go
type Employeee struct {
   Id string
   Name string
   Age int
}
//第一种定义方式在实例对应方法被调用是，实例的成员会进行值复制
//func (e Employeee) String() string {
// //Address is c0000745b0
// //Address is c0000745e0
// fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
// return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
//}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employeee) String() string {
   //Address is c0000745b0
   //Address is c0000745b0
   fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
   return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

func TestStructOperations(t *testing.T) {
   e := Employeee{"0", "Bob", 20}
   fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
   t.Log(e.String())
}
```
参考代码：[struct_operation_test](/src/ch11/encapsulation/struct_operation_test.go)

## 1.19 接口与依赖
Java的接口定义

![interface_java](/images/interface_java.png)

Go的接口定义

![interface_go](/images/interface_go.png)

```go
import "testing"

type Programmer interface {
   WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() string  {
   return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
   var p Programmer
   p = new(GoProgrammer)
   t.Log(p.WriteHelloWorld())
}
```
与其他主要编程语言的差异
* 1、接口为非入侵性，实现不依赖于接口定义
* 2、所以接口的定义可以包含在接口使用者包内

参考代码：[interface_test](/src/ch11/interface/interface_test.go)

**接口变量**

![interface_val](/images/interface_val.png)

## 1.20 自定义类型
```go
type IntConv func(op int) int
type MyPoint int
```

## 1.21 继承（扩展和复用）
能继承，但不能重写方法

```go
type Pet struct {
}

func (p *Pet) Speak() {
   fmt.Print("Pet Speak")
}

func (p *Pet) SpeakTo(host string) {
   p.Speak()
   fmt.Println("Pet SpeakTo ", host)
}

type Dog struct {
   Pet
}

func (d *Dog) SpeakTo(host string) {
   fmt.Println("Dog SpeakTo", host)
}

func TestDog(t *testing.T) {
   dog := new(Dog)
   dog.Pet.SpeakTo("Hello World!")
   dog.SpeakTo("Hello World!")
}
```
参考代码：[extension_test](/src/ch12/extension/extension_test.go)

## 1.22 多态
利用接口进行多态设计
```go
type Code string
type Programmer interface {
   WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer)WriteHelloWorld() Code {
   return "fmt.Println(\"Go Programmer say : Hello World!\")"
}

type JavaProgrammer struct {

}

func (java * JavaProgrammer)WriteHelloWorld() Code {
   return "fmt.Println(\"Java Programmer say : Hello World!\")"
}

func TestPolymorphism(t *testing.T) {
   var p1 Programmer = new(GoProgrammer)
   var p2 Programmer = new(JavaProgrammer)
   p1.WriteHelloWorld()
   p2.WriteHelloWorld()
}
```
参考代码：[polymorphism_test](/src/ch13/polymorphism/polymorphism_test.go)

## 1.23 空接口与断言
* 1、空接口可以表示任何类型
* 2、通过断言来将空接口转换为指定类型
  - v,ok:=p.(int)
```go
func DoSomething(p interface{})  {
   if i,ok:=p.(int);ok {
      fmt.Println("Integer",i)
      return
   }
   if i,ok:=p.(string);ok {
      fmt.Println("string",i)
      return
   }
   fmt.Println("Unknow Type")
/*
   switch v:=p.(type) {
   case int:
      fmt.Println("Integer",v)
   case string:
      fmt.Println("string",v)
   default:
      fmt.Println("Unknow Type")
   }*/
}

func TestEmptyInterfaceAssertion(t *testing.T) {
   DoSomething(10)
   DoSomething("10")
}
```
参考代码：[empty_interface_test](/src/ch13/empty_interface/empty_interface_test.go)

## 1.24 Go的错误机制
* 1、没有异常机制
* 2、error类型实现了error接口
* 3、可以通过errors.New来快速创建错误实例
```go
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 1000")

func GetFibonacci(n int) ([]int, error) {
   if n < 2 {
      return nil,LessThanTwoError
   }
   if n > 100 {
      return nil,LargerThanHundredError
   }
   fibList := []int{1,1}
   for i := 2;i < n; i++ {
      fibList = append(fibList,fibList[i-2]+fibList[i-1])
   }
   return fibList,nil
}

func TestGetFibonacci(t *testing.T) {
   if v,err := GetFibonacci(1); err != nil {
      t.Error(err)
   } else {
      t.Log(v)
   }
}
```
参考代码：[error_test](/src/ch14/error/error_test.go)

## 1.25 panic和os.Exit 退出程序，recover恢复
* panic用于不可以恢复的错误
* panic退出前会执行defer指定的内容

**panic vs os.Exit**

* os.Exit退出时不会调用defer指定的函数
* os.Exi退出时不输出当前调用栈信息

**recover**

* 用于捕捉panic抛出的异常，做恢复错误

```go
func TestPanicVxExit(t *testing.T) {

   defer func() {
      if err := recover(); err != nil {
         fmt.Println("recovered from ", err)
      }
   }()

   fmt.Println("start")
   panic(errors.New("Something wrong!"))
   //os.Exit(-1)
}
```
参考代码：[panic_recover_test](/src/ch14/panic_recover/panic_recover_test.go)

## 1.26 package
* 1、基本服用模块单元
  - 以首字母大写来表明可被包外代码访问
* 2、代码的package可以和所在的目录不一致
* 3、同一目录里的Go代码的package要保持一致
* 4、通过go get 来获取远程依赖
  - go get -u强制从网络更新远程依赖
* 5、注意代码在GitHub上的组织形式，以适应go get
  - 直接以代码路径开始，不要有src，自己放package给别人用时，特别注意
```go
package series

func GetFibonacciSerie(n int) []int {
   ret := []int{1, 1}
   for i := 2; i < n; i++ {
      ret = append(ret, ret[i-2]+ret[i-1])
   }
   return ret
}
```
```go
package client

import (
   "ch15/series"
   "testing"
)

func TestPackage(t *testing.T) {
   t.Log(series.GetFibonacciSerie(10))
}
```
参考代码：[my_series](/src/ch15/series/my_series.go)

参考代码：[package_test](/src/ch15/client/package_test.go)

**如果要引用自己写的package，需要工程的路径写到GOPATH的环境变量中**

如果是idea IDE，可以如下设置：

![gopath_ide](/images/gopath_ide.png)

## 1.27 init方法
* 在main被执行前，所有依赖的package的init方法都会被执行
* 不同包的init函数按照包导入的依赖关系决定执行顺序
* 每个包可以有多个init方法
* 包的每个源文件也可以有多个init函数，这点比较特殊

```go
func init() {
   fmt.Println("init1")
}

func init() {
   fmt.Println("init2")
}

func Square(n int) int {
   return n ^ 2
}
```
参考代码：[my_series](/src/ch15/series/my_series.go)

## 1.28 获取远程package
我的GOPATH=C:\Users\Ryan\go;D:\workspace\golearing

idea上导入：
go get -t github.com/easierway/concurrent_map

![package_get_from_remote](/images/package_get_from_remote.png)

下载的文件存在在如下目录

![package_localstore](/images/package_localstore.png)

```go
package remote

import (
   "github.com/easierway/concurrent_map"
   "testing"
)
func TestConcurrentMap(t *testing.T) {
   m := concurrent_map.CreateConcurrentMap(10)
   m.Set(concurrent_map.StrKey("key"),10)
   t.Log(m.Get(concurrent_map.StrKey("key")))
}
```
参考代码：[remote_package_test](/src/ch15/remote_package/remote_package_test.go)

## 1.29 vendor路径，包管理工具dep

* 1、当前包下的vendor目录
* 2、向上级目录查找，直到找到src下的vendor目录
* 3、在GOPATH下面查找依赖包
* 4、在GOROOT目录下查找

1、下载最新的windows的exe包

dep工具：https://github.com/golang/dep/releases

![dep_1](/images/dep_1.png)

2、放到GO的bin路径中，并改名为dep.exe

![dep_2](/images/dep_2.png)

3、在idea的setting中设置dep

![dep_3](/images/dep_3.png)

4、dep的使用，进入到代码所在的文件夹中，例如：
ch15/remote_package
```go
#dep初始化，初始化配置文件Gopkg.toml
dep init
#dep加载依赖包，自动归档到vendor目录
dep ensure
# 最终会生成vendor目录，Gopkg.toml和Gopkg.lock的文件
```

5、工程目录如下图
![dep_5](/images/dep_5.png)

## 1.30 包管理工具 Go Modules
Go Modules于1.11版本初步引入，在1.12版本中正式支持，它是Go语言官方提供的包管理解决方法。

Go Modules和传统的GOPATH不同，不需要包含固定的三个子目录，一个源代码目录，甚至空目录都可以作为Module，只要其中包含go.mod文件。

新建一个Module：
```shell
go mod init [module name]
```
将会在当前目录生成一个go.mod文件，内容为：
```go
module ch45config/local
go 1.16
```
通过require关键字引入相关依赖
```go

require (
github.com/spf13/viper v1.7.1
)
```
最后通过download命令手动下载依赖关系
```go
go mod download
```

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
参考代码：[groutine_test](/src/ch16/groutine/groutine_test.go)

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
参考代码：[share_mem_test](/src/ch17/share_mem/share_mem_test.go)

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
参考代码：[share_mem_test](/src/ch17/share_mem/share_mem_test.go)

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
参考代码：[async_service_test](/src/ch18/csp/async_service_test.go)

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
参考代码：[select_test](/src/ch19/select/select_test.go)

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
参考代码：[channel_close_test](/src/ch20/channel_close/channel_close_test.go)

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
参考代码：[cancel_test](/src/ch20_/calcel_by_close/cancel_test.go)

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
参考代码：[cancel_by_context_test](/src/ch20_2/cancel_by_context/cancel_by_context_test.go)

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
参考代码：[once_test](/src/ch21/singleton/once_test.go)

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
参考代码：[first_response_test](/src/ch22/util_anyone_reply/first_response_test.go)

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
参考代码：[all_response_test](/src/ch23/util_all_reply/all_response_test.go)

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
参考代码：[obj_pool](/src/ch32/obj_pool/obj_pool.go)

参考代码：[obj_pool_test](/src/ch32/obj_pool/obj_pool_test.go)

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
参考代码：[sync_pool_test](/src/ch33/obj_cache/sync_pool_test.go)

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
参考代码：[functions_test](/src/ch34/unit_test/functions_test.go)

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

参考代码：[concat_string_test](/src/ch35/benchmark/concat_string_test.go)

## 3.3 BDD Behavior Driven Development
行为驱动开发

**项目网站**

https://github.com/smartystreets/goconvey

**安装**

```shell
go get -u github.com/smartystreets/goconvey/convey
```

参考代码：[bdd_spec_test](/src/ch36/bdd/bdd_spec_test.go)

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
参考代码：[flexible_reflect_test](/src/ch38/flexible_reflect_test.go)

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
参考代码：[unsafe_test](/src/ch39/unsafe_pragramming/unsafe_test.go)

# 5 WEB框架

## 5.1 helloWorld
```go
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Hello World")
	})
	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}",t)
		writer.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080",nil)
}
```
URL分两种：
- 末尾是/：表示一个子树，后面可以跟其他子路径；
- 末尾不是/：表示一个叶子，固定的路径；

采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理

参考代码：[hello_http](/src/ch44/hello_http/hello_http.go)

## 5.2 http_router
```go
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

```
参考代码：[http_router](/src/ch45/http_router/http_router.go)

## 5.3 http_gin
```go
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

参考代码：[http_gin](/src/ch45/http_gin/http_gin.go)

## 5.4 中间件（类似Java的拦截器）

```go
func main() {
	//打印日志到控制台
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	r := gin.Default()
	//全局拦截器
	r.Use(GlobalMiddleware)
	//匹配路径的拦截器
	r.GET("/path",AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"data":"OK"})
		fmt.Println("path")
	})
	r.Run(":8000")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		authorized := check(token)
		if authorized {
			context.Next()
			return
		}
		context.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
		context.Abort()
		return
	}
}

func check(token string) bool {
	if token == "ginAuth" {
		return true
	}else {
		return false
	}
}

//全局中间件 允许跨域
func GlobalMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	fmt.Println("global")
	c.Next()
}
```

参考代码：[MiddleWareDemo](/src/ch45/http_gin_middleware/MiddleWareDemo.go)

## 5.5 memory存储用户登录信息

```go
type User struct {
	Id int
	Name string
	Password string
}

var UserById = make(map[int]*User)
var UserByName = make(map[string][]*User)

func main() {
	http.HandleFunc("/login",loginMemory)
	http.HandleFunc("/info",userInfo)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for _,user := range UserByName[r.Form.Get("username")]{
		fmt.Fprintf(w," %v",user)
	}
}

func loginMemory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.tpl")
		log.Println(t.Execute(w,nil))
	} else {
		_ = r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password",r.Form["password"])
		user1 := User{1,r.Form.Get("username"),r.Form.Get("password")}

		store(user1)

		if pwd := r.Form.Get("password");pwd == "123456" {
			fmt.Fprintf(w,"欢迎登录，Hello %s",r.Form.Get("username"))
		} else {
			fmt.Fprintf(w,"密码错误，请重新输入")
		}
	}

}

func store(user User) {
	UserById[user.Id] = &user
	UserByName[user.Name] = append(UserByName[user.Name],&user)
}
```
参考代码：[memory](/src/ch45/http_gin_memory/memory.go)

## 5.6 mysql

```go

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id int
	Name string
	Habits string
	CreateTime string
}

var tpl = `<html>
<head>
<title></title>
</head>
<body>
<form action="/info" method="post">
	用户名:<input type="text" name="username">
	兴趣爱好:<input type="text" name="habits">
	<input type="submit" value="提交">
</form>
</body>
</html>`

var db *sql.DB

var err error

func init() {
	db,err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/user?charset=utf8")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/form", submitForm)
	http.HandleFunc("/info", userInfo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method == "POST" {
		user1 := User{Name: r.Form.Get("username"),Habits: r.Form.Get("habits")}
		store(user1)
		fmt.Fprintf(w," %v",queryByName(r.Form.Get("username")))
	}

}

func queryByName(name string) User {
	user := User{}
	stmt,err := db.Prepare("select * from user where name=?")
	checkErr(err)

	rows,_ :=stmt.Query(name)

	for rows.Next(){
		var id int
		var name string
		var habits string
		var createdTime string
		err = rows.Scan(&id, &name, &habits, &createdTime)
		checkErr(err)
		fmt.Printf("[%d, %s, %s, %s]\n", id, name, habits, createdTime)
		user = User{id,name,habits,createdTime}
		break
	}

	return user
}

func store(user User) {
	//插入数据
	stmt,err := db.Prepare("INSERT INTO user SET name=?,habits=?,created_time=?")
	t := time.Now().UTC().Format("2006-01-02")
	res,err := stmt.Exec(user.Name,user.Habits,t)

	checkErr(err)

	id,err := res.LastInsertId()
	checkErr(err)

	fmt.Println("last insert id is: %d \n",id)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:",r.Method)//获取请求的方法
	var t *template.Template
	t = template.New("Products")//创建一个模板
	t,_ = t.Parse(tpl)
	log.Println(t.Execute(w,nil))
}
```
参考代码：[mysql](/src/ch45/http_gin_mysql/mysql.go)

## 5.7 beego ORM框架

```go
package main
import (
	"fmt"
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type Person struct {
	PersonId int    `orm:"pk"`
	Name   string `orm:"size(100)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8")

	// register model
	orm.RegisterModel(new(Person))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	person := Person{Name: "aoho"}

	// insert
	id, err := o.Insert(&person)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	person.Name = "boho"
	num, err := o.Update(&person)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Person{PersonId: person.PersonId}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	var maps []orm.Params

	res, err := o.Raw("SELECT * FROM person").Values(&maps)
	fmt.Printf("NUM: %d, ERR: %v\n", res, err)
	for _, term := range maps {
		fmt.Println(term["person_id"], ":", term["name"])
	}
	// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```

参考代码：[beego](/src/ch45/http_gin_beego/beego.go)

# 6 注册中心 consul
注册中心的类型非常多，这里只拿consul来实现
## 6.1 consul的运行

[consul 下载](https://www.consul.io/downloads)

windows环境是一个.exe结尾的可执行文件

运行一个单实例的consul服务

```shell
consul.exe agent -dev
```
运行结果：在本地8500端口启动consul服务

```shell
D:\softwork\consul_1.9.5_windows_amd64>consul.exe agent -dev
==> Starting Consul agent...
           Version: '1.9.5'
           Node ID: '2afc9ed7-d9a5-88c6-50db-e07aa64cc943'
         Node name: 'qinxiongzhou'
        Datacenter: 'dc1' (Segment: '<all>')
            Server: true (Bootstrap: false)
       Client Addr: [127.0.0.1] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
      Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
           Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false, Auto-Encrypt-TLS: false

==> Log data will now stream in as it occurs:

    2021-05-06T17:29:58.732+0800 [INFO]  agent.server.raft: initial configuration: index=1 servers="[{Suffrage:Voter ID:2afc9ed7-d9a5-88c6-50db-e07aa6
4cc943 Address:127.0.0.1:8300}]"
    2021-05-06T17:29:58.751+0800 [INFO]  agent.server.raft: entering follower state: follower="Node at 127.0.0.1:8300 [Follower]" leader=
    2021-05-06T17:29:58.752+0800 [INFO]  agent.server.serf.wan: serf: EventMemberJoin: qinxiongzhou.dc1 127.0.0.1
    2021-05-06T17:29:58.753+0800 [INFO]  agent.server.serf.lan: serf: EventMemberJoin: qinxiongzhou 127.0.0.1
    2021-05-06T17:29:58.754+0800 [INFO]  agent.router: Initializing LAN area manager
    2021-05-06T17:29:58.755+0800 [INFO]  agent.server: Adding LAN server: server="qinxiongzhou (Addr: tcp/127.0.0.1:8300) (DC: dc1)"
    2021-05-06T17:29:58.755+0800 [INFO]  agent.server: Handled event for server in area: event=member-join server=qinxiongzhou.dc1 area=wan
    2021-05-06T17:29:58.756+0800 [INFO]  agent: Started DNS server: address=127.0.0.1:8600 network=udp
    2021-05-06T17:29:58.758+0800 [INFO]  agent: Started DNS server: address=127.0.0.1:8600 network=tcp
    2021-05-06T17:29:58.759+0800 [INFO]  agent: Starting server: address=127.0.0.1:8500 network=tcp protocol=http
    2021-05-06T17:29:58.760+0800 [WARN]  agent: DEPRECATED Backwards compatibility with pre-1.9 metrics enabled. These metrics will be removed in a fu
ture version of Consul. Set `telemetry { disable_compat_1.9 = true }` to disable them.
    2021-05-06T17:29:58.762+0800 [INFO]  agent: started state syncer
    2021-05-06T17:29:58.760+0800 [INFO]  agent: Started gRPC server: address=127.0.0.1:8502 network=tcp
==> Consul agent running!
    2021-05-06T17:29:58.818+0800 [WARN]  agent.server.raft: heartbeat timeout reached, starting election: last-leader=
    2021-05-06T17:29:58.819+0800 [INFO]  agent.server.raft: entering candidate state: node="Node at 127.0.0.1:8300 [Candidate]" term=2
    2021-05-06T17:29:58.820+0800 [DEBUG] agent.server.raft: votes: needed=1
    2021-05-06T17:29:58.820+0800 [DEBUG] agent.server.raft: vote granted: from=2afc9ed7-d9a5-88c6-50db-e07aa64cc943 term=2 tally=1
    2021-05-06T17:29:58.821+0800 [INFO]  agent.server.raft: election won: tally=1
    2021-05-06T17:29:58.822+0800 [INFO]  agent.server.raft: entering leader state: leader="Node at 127.0.0.1:8300 [Leader]"
    2021-05-06T17:29:58.823+0800 [INFO]  agent.server: cluster leadership acquired
    2021-05-06T17:29:58.823+0800 [INFO]  agent.server: New leader elected: payload=qinxiongzhou
    2021-05-06T17:29:58.823+0800 [DEBUG] agent.server: Cannot upgrade to new ACLs: leaderMode=0 mode=0 found=true leader=127.0.0.1:8300
    2021-05-06T17:29:58.826+0800 [DEBUG] agent.server.autopilot: autopilot is now running
    2021-05-06T17:29:58.827+0800 [DEBUG] agent.server.autopilot: state update routine is now running
    2021-05-06T17:29:58.827+0800 [INFO]  agent.leader: started routine: routine="federation state anti-entropy"
    2021-05-06T17:29:58.829+0800 [INFO]  agent.leader: started routine: routine="federation state pruning"
    2021-05-06T17:29:58.830+0800 [DEBUG] connect.ca.consul: consul CA provider configured: id=07:80:c8:de:f6:41:86:29:8f:9c:b8:17:d6:48:c2:d5:c5:5c:7f
:0c:03:f7:cf:97:5a:a7:c1:68:aa:23:ae:81 is_primary=true
    2021-05-06T17:29:58.843+0800 [INFO]  agent.server.connect: initialized primary datacenter CA with provider: provider=consul
    2021-05-06T17:29:58.844+0800 [INFO]  agent.leader: started routine: routine="intermediate cert renew watch"
    2021-05-06T17:29:58.845+0800 [INFO]  agent.leader: started routine: routine="CA root pruning"
    2021-05-06T17:29:58.845+0800 [DEBUG] agent.server: successfully established leadership: duration=22.0042ms
    2021-05-06T17:29:58.846+0800 [INFO]  agent.server: member joined, marking health alive: member=qinxiongzhou
    2021-05-06T17:29:59.090+0800 [INFO]  agent.server: federation state anti-entropy synced
    2021-05-06T17:29:59.151+0800 [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
    2021-05-06T17:29:59.154+0800 [INFO]  agent: Synced node info
    2021-05-06T17:29:59.266+0800 [DEBUG] agent: Skipping remote check since it is managed automatically: check=serfHealth
    2021-05-06T17:29:59.267+0800 [DEBUG] agent: Node info in sync
    2021-05-06T17:29:59.269+0800 [DEBUG] agent: Node info in sync
```
结果：

![consul.png](/images/consul.png)

## 6.2 DisconverClient定义

定义：

```go
type DiscoveryClient interface {
	/**
	 * 服务注册接口
	 * @param serviceName 服务名
	 * @param instanceId 服务实例Id
	 * @param instancePort 服务实例端口
	 * @param healthCheckUrl 健康检查地址
	 * @param instanceHost 服务实例地址
	 * @param meta 服务实例元数据
	 */
	Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool
	/**
	 * 服务注销接口
	 * @param instanceId 服务实例Id
	 */
	DeRegister(instanceId string, logger *log.Logger) bool
	/**
	 * 发现服务实例接口
	 * @param serviceName 服务名
	 */
	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}
```
参考代码：[disconver_client](/src/ch45discover/discover/discover_client.go)

## 6.3 DisconverClient实现，利用go-kit

```go
package discover

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"log"
	"strconv"
	"sync"
)

type KitDiscoverClient struct {
	Host   string // Consul Host
	Port   int    // Consul Port
	client consul.Client
	// 连接 consul 的配置
	config *api.Config
	mutex sync.Mutex
	// 服务实例缓存字段
	instancesMap sync.Map
}

func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoveryClient, error) {
	// 通过 Consul Host 和 Consul Port 创建一个 consul.Client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		Host:   consulHost,
		Port:   consulPort,
		config:consulConfig,
		client: client,
	}, err
}

func (consulClient *KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool {

	// 1. 构建服务实例元数据
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			HTTP:                           "http://" + instanceHost + ":" + strconv.Itoa(instancePort) + healthCheckUrl,
			Interval:                       "15s",
		},
	}

	// 2. 发送服务注册到 Consul 中
	err := consulClient.client.Register(serviceRegistration)

	if err != nil {
		log.Println("Register Service Error!")
		return false
	}
	log.Println("Register Service Success!")
	return true
}

func (consulClient *KitDiscoverClient) DeRegister(instanceId string, logger *log.Logger) bool {

	// 构建包含服务实例 ID 的元数据结构体
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	// 发送服务注销请求
	err := consulClient.client.Deregister(serviceRegistration)

	if err != nil {
		logger.Println("Deregister Service Error!")
		return false
	}
	log.Println("Deregister Service Success!")

	return true
}

func (consulClient *KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {

	//  该服务已监控并缓存
	instanceList, ok := consulClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	// 申请锁
	consulClient.mutex.Lock()
	defer consulClient.mutex.Unlock()
	// 再次检查是否监控
	instanceList, ok = consulClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	} else {
		// 注册监控
		go func() {
			// 使用 consul 服务实例监控来监控某个服务名的服务实例列表变化
			params := make(map[string]interface{})
			params["type"] = "service"
			params["service"] = serviceName
			plan, _ := watch.Parse(params)
			plan.Handler = func(u uint64, i interface{}) {
				if i == nil {
					return
				}
				v, ok := i.([]*api.ServiceEntry)
				if !ok {
					return // 数据异常，忽略
				}
				// 没有服务实例在线
				if len(v) == 0 {
					consulClient.instancesMap.Store(serviceName, []interface{}{})
				}
				var healthServices []interface{}
				for _, service := range v {
					if service.Checks.AggregatedStatus() == api.HealthPassing {
						healthServices = append(healthServices, service.Service)
					}
				}
				consulClient.instancesMap.Store(serviceName, healthServices)
			}
			defer plan.Stop()
			plan.Run(consulClient.config.Address)
		}()
	}

	// 根据服务名请求服务实例列表
	entries, _, err := consulClient.client.Service(serviceName, "", false, nil)
	if err != nil {
		consulClient.instancesMap.Store(serviceName, []interface{}{})
		logger.Println("Discover Service Error!")
		return nil
	}
	instances := make([]interface{}, len(entries))
	for i := 0; i < len(instances); i++ {
		instances[i] = entries[i].Service
	}
	consulClient.instancesMap.Store(serviceName, instances)
	return instances
}

```
参考代码：[kit_discover_client](/src/ch45discover/discover/kit_discover_client.go)

## 6.4 main运行

```go
func main() {
	// 从命令行中读取相关参数，没有时使用默认值
	var (
		// 服务地址和服务名
		servicePort = flag.Int("service.port", 10086, "service port")
		serviceHost = flag.String("service.host", "127.0.0.1", "service host")
		serviceName = flag.String("service.name", "SayHello", "service name")
		// consul 地址
		consulPort = flag.Int("consul.port", 8500, "consul port")
		consulHost = flag.String("consul.host", "127.0.0.1", "consul host")
	)
	flag.Parse()

	// 声明服务发现客户端
	var discoveryClient discover.DiscoveryClient
	discoveryClient, err := discover.NewKitDiscoverClient(*consulHost, *consulPort)
	// 声明并初始化 Service
	var svc = service.NewDiscoveryServiceImpl(discoveryClient)

	// 省略...
	
	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, config.KitLogger)
	// 定义服务实例ID
	instanceId := *serviceName + "-" + uuid.Must(uuid.NewV4()).String()
	// 启动 http server
	go func() {
		config.Logger.Println("Http Server start at port:" + strconv.Itoa(*servicePort))
		//启动前执行注册
		if !discoveryClient.Register(*serviceName, instanceId, "/health", *serviceHost,  *servicePort, nil, config.Logger){
			config.Logger.Printf("string-service for service %s failed.", serviceName)
			// 注册失败，服务启动失败
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"  + strconv.Itoa(*servicePort), handler)
	}()

	error := <-errChan
	//服务退出取消注册
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}
```

参考代码：[main](/src/ch45discover/main.go)

# 7 RPC 远程过程调用

Remote Procedure Call。它是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。

## 7.1 简易的Go语言原生RPC

- 方法是可输出的，例如：方法名的首字母必须大写
- 方法必须有两个参数，必须是输出类型或者是内建类型
- 方法的第二个参数是指针类型
- 方法返回类型为error

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```
server 端
```go
type StringRequest struct {
	A string
	B string
}

type Service interface {
	//Concat a and b
	Concat(req StringRequest,ret *string) error
	//a,b pkg string value
	Diff(req StringRequest,ret *string) error
}

type StringService struct {
}

func (s StringService) Concat(req StringRequest, ret *string) error {
	if len(req.A) + len(req.B) > StrMaxSize{
		*ret = ""
		return ErrMaxSize
	}
	*ret = req.A + req.B
	return nil
}

func (s StringService) Diff(req StringRequest, ret *string) error {
	if len(req.A) < 1  || len(req.B) < 1{
		*ret = ""
		return nil
	}
	res := ""
	if len(req.A) >= len(req.B){
		for _,char := range req.B{
			if strings.Contains(req.A,string(char)){
				res = res + string(char)
			}
		}
	} else {
		for _,char := range req.A{
			if strings.Contains(req.B,string(char)){
				res = res + string(char)
			}

		}
	}
	*ret = res
	return nil
}
```

参考代码：[server](/src/ch45rpc/basic/string-service/server.go)

server端启动

```go
func main() {
	stringService := new(service.StringService)
	registerError := rpc.Register(stringService)

	if registerError != nil {
		log.Fatal("Register error: ",registerError)
	}

	rpc.HandleHTTP()
	l,e := net.Listen("tcp","127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error: ",e)
	}
	http.Serve(l,nil)
}
```
参考代码：[server](/src/ch45rpc/basic/server.go)

client 端

```go
func TestClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//Synchronous call
	stringReq := &service.StringRequest{"A","B"}
	var reply string
	err = client.Call("StringService.Concat",stringReq,&reply)
	if err != nil {
		log.Fatal("StringService error:", err)
	}
	fmt.Printf("StringService Concat : %s concat %s = %s\n",stringReq.A,stringReq.B,reply)


	//async call
	stringReq = &service.StringRequest{"ACD","BDF"}
	call := client.Go("StringService.Diff",stringReq,&reply,nil)
	_ = <-call.Done
	fmt.Printf("StringService Diff : %s diff %s = %s\n",stringReq.A,stringReq.B,reply)
}
```

参考代码：[client_test](/src/ch45rpc/basic/client_test.go)

## 7.2 gRPC

### 7.2.1 安装

grpc的安装过程，随着时间推移，可能有差别。本节文字写于2021-5-11，确保可用。如有问题，请自己参考官网教程。

[grpc官网](https://grpc.io/docs/languages/go/quickstart/)

- Go的安装，安装最新的三个主要版本。安装说明，请参考[Go’s Getting Started guide](https://golang.org/doc/install)
- Protocol Buffer安装，即protoc，要求是3.0版本。[Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
  windows环境安装，请下载protoc-3.16.0-win64.zip。[Protocol_离线win版本](https://github.com/protocolbuffers/protobuf/releases/tag/v3.16.0/) 并把路径配置到path环境变量中
- Go plugins安装
  * 安装protoc-gen-go和protoc-gen-go-grpc插件
  ```go
  go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
  ```
  * 更新到path环境变量中，让protoc可以找到这两个插件。下载后的路径一般在GOPATH路径的bin目录中，例如：C:\Users\ryan\go\bin

### 7.2.2 编译运行

写两个接口的proto文件，两个消息体
```go
syntax = "proto3";
option go_package = "ch45rpc/grpcdemo/pb";
package pb;
service StringService{
    rpc Concat(StringRequest) returns (StringResponse) {}
    rpc Diff(StringRequest) returns (StringResponse) {}
}
message StringRequest {
    string A = 1;
    string B = 2;
}
message StringResponse {
    string Ret = 1;
    string err = 2;
}
```
参考代码：[string.proto](/src/ch45rpc/grpcdemo/pb/string.proto)

**编译**
```shell
protoc --go_out=. --go_opt=paths=source_relative \ 
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \ 
       pb\string.proto
```

生成string.pb.go、string_grpc.pb.go两个文件

### 7.2.3 接口实现

**接口实现**
```go
type StringService struct{
	pb.UnsafeStringServiceServer
}
func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := pb.StringResponse{Ret: req.A + req.B}
	return &response, nil
}
func (s *StringService) Diff(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A) < 1 || len(req.B) < 1 {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	response := pb.StringResponse{Ret: res}
	return &response, nil
}
```
参考代码：[service.go](/src/ch45rpc/grpcdemo/string-service/service.go)

**server端**
```go
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
```
参考代码：[server.go](/src/ch45rpc/grpcdemo/server.go)

**client端**

```go
func main() {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	bookClient := pb.NewStringServiceClient(conn)
	stringReq := &pb.StringRequest{A: "A", B: "B"}
	reply, _ := bookClient.Concat(context.Background(), stringReq)
	fmt.Printf("StringService Concat : %s concat %s = %s\n", stringReq.A, stringReq.B, reply.Ret)
}
```
参考代码：[client.go](/src/ch45rpc/grpcdemo/client.go)

## 7.3 gRPC流式编程

gRPC可以定义4中类型的服务接口，分别是一元RPC、服务器流RPC、客户端流式RPC和双向流RPC

- 一元RPC是指客户端向服务器发送请求并获得响应，就像正常的函数调用一样。
  ```go
  rpc Concat(StringRequest) returns (StringResponse){}
  ```
- 服务器流RPC是指客户端发送一个对象，服务器端返回一个Stream（流式消息）
  ```go
  rpc LotsOfServerStream(StringRequest) returns (stream StringResponse){}
  ```
- 客户端流式RPC，客户端发送一个stream(流式消息)服务端返回一个对象
  ```go
  rpc LotsOfClientStream(Stream StringRequest) returns (StringResponse){}
  ```
- 双向流RPC，两个流独立运行。类似WebSocket(长链接)，客户端可以向服务端请求消息，服务端也可以向客户端请求消息
  ```go
  rpc LotsOfServerAndClientStream(stream StringRequest) returns (stream StringResponse){}
  ```

参考代码：[string.proto](/src/ch45rpc/grpcstreamdemo/pb/string.proto)

**编译**
```shell
protoc --go_out=. --go_opt=paths=source_relative \ 
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \ 
       pb\string.proto
```

**双向流RPC 服务端**
```go
func (s *StringService) LotsOfServerAndClientStream(server pb.StringService_LotsOfServerAndClientStreamServer) error {
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv %v", err)
			return err
		}
		server.Send(&pb.StringResponse{Ret: in.A + in.B})
	}
	return nil
}
```
参考代码：[service.go](/src/ch45rpc/grpcstreamdemo/string-service/service.go)

**双向流RPC 客户端**
```go
func sendClientAndServerStreamRequest(client pb.StringServiceClient) {
	fmt.Printf("test sendClientAndServerStreamRequest \n")
	var err error
	stream, err := client.LotsOfServerAndClientStream(context.Background())
	if err != nil {
		log.Printf("failed to call: %v", err)
		return
	}
	var i int
	for {
		err1 := stream.Send(&pb.StringRequest{A: strconv.Itoa(i), B: strconv.Itoa(i + 1)})
		if err1 != nil {
			log.Printf("failed to send: %v", err)
			break
		}
		reply, err2 := stream.Recv()
		if err2 != nil {
			log.Printf("failed to recv: %v", err)
			break
		}
		log.Printf("sendClientAndServerStreamRequest Ret is : %s", reply.Ret)
		i++
	}
}
```
参考代码：[client.go](/src/ch45rpc/grpcstreamdemo/client.go)


# 8 配置管理

## 8.1 本地化配置viper

配置文件
```yaml
RegisterTime: "2019-6-18 10:00:00"
Address: "Shanghai"
ResumeInformation:
  Name: "aoho"
  Sex: "male"
  Age: 20
  Habits:
    - "Basketball"
    - "Running"

```

配置类
```go
func init() {
	viper.AutomaticEnv()
	initDefault()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	if err := sub("ResumeInformation", &Resume); err != nil {
		log.Fatal("Fail to parse config", err)
	}
}
func initDefault() {
	//设置读取的配置文件
	viper.SetConfigName("resume_config")
	//添加读取的配置文件路径
	viper.AddConfigPath("./ch45config/local/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	viper.AddConfigPath("E:/qinxiongzhou/go/learning/src/ch45config/local/")
	//设置配置文件类型
	viper.SetConfigType("yaml")
}
func main() {
	fmt.Printf(
		"姓名: %s\n爱好: %s\n性别: %s \n年龄: %d \n",
		Resume.Name,
		Resume.Habits,
		Resume.Sex,
		Resume.Age,
	)
	//反序列化
	parseYaml(viper.GetViper())
	fmt.Println(Contains("Basketball", Resume.Habits))
}
func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
type ResumeInformation struct {
	Name   string
	Sex    string
	Age    int
	Habits []interface{}
}
type ResumeSetting struct {
	RegisterTime      string
	Address           string
	ResumeInformation ResumeInformation
}
func parseYaml(v *viper.Viper) {
	var resumeConfig ResumeSetting
	if err := v.Unmarshal(&resumeConfig); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println("resume config:\n ", resumeConfig)
}
func sub(key string, value interface{}) error {
	log.Printf("配置文件的前缀为：%v", key)
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	return sub.Unmarshal(value)
}
```

运行，可能会出现如下异常
```shell
cannot find package "github.com/hashicorp/hcl/hcl/printer" volatiletech/sqlboiler#904
```
这个的原因是viper采用Modules的方式，引用的hcl包是v1.0.0版本的。我们的工程，也需要改成modules的方式，否则会使用hcl的最新版本，而最新版本中，并没有printer的包。

**解决方案：**
项目目录下，执行mod init，生成go.mod文件
```shell
>go mod init

go: creating new go.mod: module ch45config/local
go: to add module requirements and sums:
        go mod tidy
```

在生成的go.mod文件中，添加如下
```go
require (
	github.com/spf13/viper v1.7.1
)
```
设置idea：打开Go Modules模块。File>settings>Go>Go Modules
![goModules.png](/images/goModules.png)
结果：
![ExternalLibrariesGoModules.png](/images/ExternalLibrariesGoModules.png)

如果idea没有显示Go Modules模块，请把go.mod文件放到项目根目录上。

# 9 网关

## 9.1 自定义网关
**关键代码：** 主要是为http.ListenAndServe提供一个proxy的方法
```go

func main() {
	//创建反向代理
	proxy := NewReverseProxy(consulClient, logger)
	//开始监听
	go func() {
		logger.Log("transport", "HTTP", "addr", "9090")
		errc <- http.ListenAndServe(":9090", proxy)
	}()
}

// NewReverseProxy 创建反向代理处理方法
func NewReverseProxy(client *api.Client, logger log.Logger) *httputil.ReverseProxy {

	//创建Director
	director := func(req *http.Request) {

		//查询原始请求路径
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		//按照分隔符'/'对路径进行分解，获取服务名称serviceName
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]

		//调用consul api查询serviceName的服务实例列表
		result, _, err := client.Catalog().Service(serviceName, "", nil)
		if err != nil {
			logger.Log("ReverseProxy failed", "query service instance error", err.Error())
			return
		}

		if len(result) == 0 {
			logger.Log("ReverseProxy failed", "no such service instance", serviceName)
			return
		}

		//重新组织请求路径，去掉服务名称部分
		destPath := strings.Join(pathArray[2:], "/")

		//随机选择一个服务实例
		tgt := result[rand.Int()%len(result)]
		logger.Log("service id", tgt.ServiceID)

		//设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
		req.URL.Path = "/" + destPath
	}
	return &httputil.ReverseProxy{Director: director}
```
参考代码：[myGateway.go](/src/ch45-gateway/gateway/main.go)

**运行：**

1、运行consul。如有疑惑，请参考6.1章节
```shell
consul.exe agent -dev
```
查看consul，http://localhost:8500/ui/dc1/services

2、运行service服务。代码在/ch45discover/main.go。

服务运行在10086端口，有say-hello的接口，http://localhost:10086/say-hello ，返回如下
```json
{
  message: "Hello World!"
}
```

3、运行网关服务。代码在/ch45-gateway/gateway/main.go。

访问http://localhost:9090/SayHello/say-hello ，返回如下：
```json
{
  message: "Hello World!"
}
```

## 9.2 kong网关

