# golearing
#### Go 语言学习的测试案例，从入门开始
#### 实实在在操一边，才能真正掌握。

### Thread vs. Groutine
#### 1. 创建时默认的stack的大小
* JDK5以后Java Thread stack默认为1M
* Groutine 的stack初始化大小为2K
#### 2.和KSE（Kernel Space Entity）的对应关系
* Java Thread 是1：1
* Groutine 是M：N

```java
for i := 0; i < 10; i++ {
    go func (j int){
        fmt.Println(j)
    }(i)
}
```
详情请见:src/ch16/groutine/groutine_test.go