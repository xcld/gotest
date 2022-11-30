# 今日知识点
### 今日作业
#### 课后思考题
- Switch和select有什么区别
- print和printf和println有什么区别
- 对比Python和Go的关键字、数据类型、常变量、以及流程控制有什么区别参考[Goguide](https://github.com/coderit666/GoGuide)
```
    关键字都是预留好的，有指定作用的；
    常变量：python要全部大写，golang不要求全部大写；相同都是不能对常量进行修改
```

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
```
    选D，因为变量y，z没有声明类型，所以编译会不通过,错误
    应该选A，默认会转换为精度更高的类型
```

### 题目2

```go
const a = 7.0
var b = 2
fmt.Println(a / b)
```

* A: `3.5`
* B: `3 `
* C: 编译错误
```
    输出3，选B，解释同第一题
```
### 题目3

```go
a := 40
f := float64(a/100.0)
fmt.Println(f)
```

* A: `0`

* B: `0.4`

* C: 编译错误
```
    选错了，选了A；运行结果是B，原因不清楚
```
####  思考题

题目1：

``` go
var (
    a = 1.0
    b = 2
)
fmt.Println(a / b)

//输出：0.5；答案：编译错误，不同的数据类型不能直接运算，要转换成相同格式的
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



  
