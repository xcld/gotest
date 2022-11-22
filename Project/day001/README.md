# 今日知识点
### 今日作业
#### 课后思考题
- Switch和select有什么区别
- print和printf和println有什么区别
- 对比Python和Go的关键字、数据类型、常变量、以及流程控制有什么区别参考[Goguide](https://github.com/coderit666/GoGuide)
#### Go Quiz: 从Go面试题看数值类型的自动推导

##  背景

Google工程师Valentin Deleplace出了几道关于Go数值类型的计算题，很有迷惑性，整体正确率不到50%，拿出来和大家分享下。

### 题目1

```go
var y = 5.2
const z = 2
fmt.Println(y / z)
```

* A: `2.6`

* B: `2.5 `

* C: `2`

* D: 编译错误

### 题目2

```go
const a = 7.0
var b = 2
fmt.Println(a / b)
```

* A: `3.5`
* B: `3 `
* C: 编译错误

### 题目3

```go
a := 40
f := float64(a/100.0)
fmt.Println(f)
```

* A: `0`

* B: `0.4`

* C: 编译错误

####  思考题

题目1：

``` go
var (
    a = 1.0
    b = 2
)
fmt.Println(a / b)
```

题目2：

```go
const (
    x = 5.0
    y = 4
)
fmt.Println(x / y)
```

题目3：

```go
const t = 4.8
var u = 2
fmt.Println(t / u)
```

题目4：

```go
d := 9.0
const f int = 2
fmt.Println(d / f)
```



  
