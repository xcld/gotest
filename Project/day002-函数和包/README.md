# day002 函数和包

----

# 知识点回顾：

- 5.1 定义格式
- 5.2 自定义函数
  - 5.2.1 无参无返回值
  - 5.2.2 有参无返回值
    - 5.2.2.1 普通参数列表
    - 5.2.2.2 不定参数列表
      1) 不定参数类型
      2) 不定参数的传递
  - 5.2.3 无参有返回值
    - 5.2.3.1 一个返回值
    - 5.2.3.2 多个返回值
  - 5.2.4 有参有返回值
- 5.3 递归函数
- 5.4 函数类型
- 5.5 作用域
  - 5.5.1 局部变量
  - 5.5.2 全局变量
  - 5.5.3 不同作用域同名变量
- 5.6 匿名函数与闭包
   `闭包=函数+引用环境` 匿名函数都是闭包
```go
//闭包案例示范，文件后缀名检测程序
func makeSuffixFunc(suffix string) func(string) string {
 //声明一个makeSuffixFunc函数，接收一个suffix参数 ，返回一个func（//接收string参数）返回string类型的函数
    return func(name string) string {  //return 一个和上面定义结构一样的匿名函数
        if !strings.HasSuffix(name, suffix) { //判断字符串后缀的string包，如果!(name以suffix为后缀返回true)
			// name 是匿名函数定义变量，这个suffix是外层函数的变量  闭包=函数+外层变量的引用
             return name + suffix
        }
        return name
    }
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
    txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
    fmt.Println(txtFunc("test")) //test.txt
}
```
- 5.7 延迟调用defer
   - 5.7.1 defer作用
   - 5.7.2 多个defer执行顺序
   - 5.7.3 defer和匿名函数结合使用
- 5.8 获取命令行参数
- 6 工程管理
  - 6.1 工作区
    - 6.1.1 工作区介绍
    - 6.1.2 GOPATH设置
  - 6.2 包
    - 6.2.1 自定义包
    - 6.2.2 main包
    - 6.2.3 main函数和init函数
    - 6.2.4 导入包
      - 6.2.4.1 点操作
      - 6.2.4.2 别名操作
      - 6.2.4.3 _操作
  - 6.3 测试案例
    - 6.3.1 测试代码
    - 6.3.2 GOPATH设置
    - 6.3.3 编译运行程序
    - 6.3.4 go install的使用
  - [6.4 包管理工具](https://learnku.com/go/t/33859)
## defer知识点总结
###  defer六大原则
总结下`defer`语义要注意的六大关键点：

1. defer后面跟的必须是函数或者方法调用，defer后面的表达式不能加括号。

   ```go
   defer (fmt.Println(1)) // 编译报错，因为defer后面跟的表达式不能加括号
   ```

2. 被defer的函数或方法的参数的值在执行到defer语句的时候就被确定下来了。

   ```go
   func a() {
       i := 0
       defer fmt.Println(i) // 最终打印0
       i++
       return
   }
   ```

   上例中，被defer的函数fmt.Println的参数`i`在执行到defer这一行的时候，`i`的值是0，fmt.Println的参数就被确定下来是0了，因此最终打印的结果是0，而不是1。

3. 被`defer`的函数或者方法如果存在多级调用，只有最后一个函数或方法会被`defer`到函数return或者panic之前执行，参见上面的说明。

4. 被defer的函数执行顺序满足LIFO原则，后defer的先执行。

 ```go
func b() {
   for i := 0; i < 4; i++ {
       defer fmt.Print(i)
   }
}
```

   上例中，输出的结果是3210，后defer的先执行。

5. 被defer的函数可以对defer语句所在的函数的命名返回值(named return value)做读取和修改操作。

   ```go
   // f returns 42
   func f() (result int) {
   	defer func() {
   		// result is accessed after it was set to 6 by the return statement
   		result *= 7
   	}()
   	return 6
   }
   ```

   上例中，被defer的函数func对defer语句所在的函数**f**的命名返回值result做了修改操作。

   调用函数`f`，返回的结果是42。

   执行顺序是函数`f`先把要返回的值6赋值给result，然后执行被defer的函数func，result被修改为42，然后函数`f`返回result给调用方，也就返回了42。

6. 即使`defer`语句执行了，被`defer`的函数不一定会执行。对这句话不理解的可以参考下面的思考题。

* 被defer的函数值(function value)在执行到defer语句的时候就被确定下来了。

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	defer func() {
  		r := recover()
  		fmt.Println(r)
  	}()
  	var f func(int) // f没有初始化赋值，默认值是nil
  	defer f(1) // 函数变量f的值已经确定下来是nil了
  	f = func(a int) {}
  }
  ```

* 如果被`defer`的函数或方法的值是nil，在执行`defer`这条语句的时候不会报错，但是最后调用nil函数或方法的时候就引发`panic: runtime error: invalid memory address or nil pointer dereference`。



# 操作题
1. 根据课件将上面知识点回顾都手写一遍代码，并备注有疑问待深入理解的地方。
2. 整数顺序 ，编写函数，返回其（两个）参数正确的（自然）数字顺序：
   - f(7,2) → 2,7
   - f(2,7) → 2,7
```go
func sort_two_num(a, b int) (int) {
	if a >= b {
		return b, a
	}
	return a, b
}
```
3. 变参，编写函数接受整数类型变参，并且每行打印一个数字。
```go
func any_arg(args ...int) {
	for _, n := range args {
		fmt.Println("任意个数的参数n： ", n)
	}
}
```
~~4. -斐波那契数列，斐波那契数列以：1, 1, 2, 3, 5, 8, 13, . . . 开始。或者用数学形式表达：x1 = 1; x2 =
      1; xn = xn−1 + xn−2 ∀n > 2。
      编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列~~ （到数组继续）

5. 函数返回一个函数：
   - 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名
        称叫做 plusTwo。然后可以像下面这样使用：
```go
p := plusTwo()
fmt.Printf("%v\n", p(2))
```
   - 使上面函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整
      数加上 x。 
```go
func plus2(a int) (b int) {
	return a + 2
}

func plusX(a, x int) (b int) {
	return a + x
}

// 编写一个函数返回另外一个函数，用到闭包+匿名函数的知识，外层函数传入参数，内层匿名函数引用外层函数参数。
//训练闭包
//func plus2() func(int) {
//	return func(x int) int { return x+2} //在返回语句中定义了一个 +2 的函数
//}

//同理可得还是闭包，
//func PlusX(x int) func(int) int{
//	return func(y int) int {return x+y}
//}
```
6. 程序改错题，下面的程序有什么错误？
```go
func main() {
for i := 0; i < 10; i++ {
fmt.Printf("%v\n", i)
}
fmt.Printf("%v\n", i)
}
```
    i是局部变量，在for循环结束后就会被回收掉，所有后面的i会找不到
点评：说得没错，找不到直接定义一个全局变量i，这题就搞完了。
# 思考题
1. ret2的值是什么？
```go
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) 
}

// 结果是30
```
## 闭包
2. 以下程序输出的值为？为什么
```go
// 程序1
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //此时是闭包，内部函数对外部函数的变量的调用，所以x不会被释放掉，10
    fmt.Println(f(20))  // 30
    fmt.Println(f(30))  //60
    
    f1 := adder()
    fmt.Println(f1(40))  // 此时是新建了一个闭包所以，这里x又从0开始了,40
    fmt.Println(f1(50))  // 90
}
```
```go
//程序2
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))  // 此时也是闭包的调用，所以base的起始值为10且不会被释放掉，11 9
    fmt.Println(f1(3), f2(4))  // 12 8
    fmt.Println(f1(5), f2(6)) //13 9
}
```
## defer
题目1：程序运行结果是什么？

```go
package main
type T int
func (t T) M(n int) T {
	print(n)
	return t
}
func main() {
	var t T
	defer t.M(1).M(2)
	t.M(3).M(4)
}
```
- 函数的用法没看懂，defer的顺序不理解，为为什么是1342，不是3412
### 上题解析
这道题主要考察以下知识点：

* 被`defer`的函数或方法什么时候执行？

* 被`defer`的函数或方法的参数的值是什么时候确定的？

* 被`defer`的函数或方法如果存在多级调用是什么机制？比如本题的`t.M(n).M(n)`就存在二级调用，先调用了`M(n)`函数，再调用了`M()`方法。

   

## 解析

我们再看看官方文档怎么说的：

>Each time a "defer" statement executes, the function value and parameters to
>the call are evaluated as usual and saved anew but the actual function is not 
>invoked. 
>
>Instead, deferred functions are invoked immediately before the 
>surrounding function returns, in the reverse order they were deferred. 
>
>That is, if the surrounding function returns through an explicit return statement, 
>deferred functions are executed after any result parameters are set by that 
>return statement but before the function returns to its caller. 
>
>If a deferred function value evaluates to nil, execution panics when the function is 
>invoked, not when the "defer" statement is executed.
官方文档的前两句话对我们求解本题至关重要，用中文来表述就是：

**假设我们在函数`A`的函数体内运行了`defer B(params)`，那被`defer`的函数`B`的参数会像普通函数调用一样被即时计算，但是被`defer`的函数`B`的调用会延迟到函数`A` return或者panic之前执行。**

如果`defer`后面跟的是多级函数的调用，只有最后一个函数会被延迟执行。

比如上例里的`defer t.M(n)`中，`M(n)`是会被即时执行的，不会延后，只有最后一个方法`M(n)`的调用会被延后执行。

同时，函数的传参也是被即时计算的，也就是`defer t.M(n).M(n)`里涉及的参数`n`的值也是被即时计算保存好的，延后执行的时候就用事先计算好的值。

题目2："end"会被打印么？`f(1, 2)`会不会编译报错？

```go
package main
import "fmt"
type Add func(int, int) int
func main() {
	var f Add
	defer f(1, 2)
	fmt.Println("end")
}
// 会被打印， 不清楚原因
```
题目3：
```go
package main
func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()
	
	var f func()
	defer f()
	f = func() {
		r += 2
	}
	return 1
}
func main() {
	println(bar())
}
```

- A: 1
- B: 7
- C: 12
- D: 13
- E: 15
### 这题看不懂
主要考察以下知识点：

* 被`defer`的函数的值是什么时候确定的？**注意**：这里说的是函数值，不是函数的参数。
* 如果函数的值是`nil`，那`defer`一个nil函数是什么结果？
* 多个被`defer`的函数的执行先后顺序，遵循LIFO原则。
* `defer` 和`recover`结合可以捕获`panic`。
* `defer`如果对函数的命名返回值参数做修改，有什么影响？

## 解析

### 函数值

上面提到了函数值的概念，对应的英文是`function value`。可能刚学习Go的同学还不太了解，我先讲解下这个概念。

首先，函数和`struct`，`int`等一样，也是一个类型(type)。

我们可以先定义一个函数类型，再声明一个该函数类型的变量，并且给该变量赋值。该函数类型变量的值我们就可以叫做函数值。看下面的代码示例就一目了然了。

```go
package main
import "fmt"
// 函数类型FuncType
type FuncType func(int) int
// 定义变量f1, 类型是FuncType
var f1 FuncType = func(a int) int { return a }
// 定义变量f, 类型是一个函数类型，函数签名是func(int) int
// 在main函数里给f赋值，零值是nil
var f func(int) int
func main() {
	fmt.Println(f1(1))
	// 定义变量f2，类型是FuncType
	var f2 FuncType
	f2 = func(a int) int {
		a++
		return a
	}
	fmt.Println(f2(1))
	// 给函数类型变量f赋值
	f = func(a int) int {
		return 10
	}
	fmt.Println(f(1))
}
```

我们平时实现函数的时候，通常是把函数的类型定义，函数变量和变量赋值一起做了。

函数类型的变量如果只声明不初始化，零值是nil，也就是默认值是nil。

通过上面的例子我们知道可以先定义函数类型，后面再定义函数类型的变量。



### 原理解析

我们再回顾下官方文档怎么说的：

>Each time a "defer" statement executes, the function value and parameters to
>the call are evaluated as usual and saved anew but the actual function is not 
>invoked. 
>
>Instead, deferred functions are invoked immediately before the 
>surrounding function returns, in the reverse order they were deferred. 
>
>That is, if the surrounding function returns through an explicit return statement, 
>deferred functions are executed after any result parameters are set by that 
>return statement but before the function returns to its caller. 
>
>If a deferred function value evaluates to nil, execution panics when the function is 
>invoked, not when the "defer" statement is executed.
本题的代码里的`bar()`函数在return之前按照如下顺序执行

| 执行                                | 执行结果                                                     |
| ----------------------------------- | ------------------------------------------------------------ |
| 执行return 1                        | 把1赋值给函数返回值参数`r`                                   |
| 执行f()                             | 因为`f`的值在`defer f()`这句执行的时候就已经确定下来是nil了，所以会引发panic |
| 执行`bar`函数里第1个被`defer`的函数 | r先加4，值变为5，然后recover捕获上一步的panic，r的值再加8，结果就是13 |
| `bar()`返回`r`的值                  | r的值是13，返回13。`main`里打印13                            |

因此本题的运行结果是13，答案是`D`。



##  总结

在`defer`第1篇文章里我们列出了`defer`六大原则，参考[Go defer语义6大原则](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=1521159887&lang=zh_CN#rd)。

本文再总结下本题目涉及的`defer`语义另外2大注意事项：

* 被defer的函数值(function value)在执行到defer语句的时候就被确定下来了。

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	defer func() {
  		r := recover()
  		fmt.Println(r)
  	}()
  	var f func(int) // f没有初始化赋值，默认值是nil
  	defer f(1) // 函数变量f的值已经确定下来是nil了
  	f = func(a int) {}
  }
  ```

* 如果被`defer`的函数或方法的值是nil，在执行`defer`这条语句的时候不会报错，但是最后调用nil函数或方法的时候就引发`panic: runtime error: invalid memory address or nil pointer dereference`。


题目4：程序运行结果？
```go
package main
import "fmt"
func test1() int {
	var i = 0
	defer func() {
		i = 10
	}()
	return i
}
func test2() (result int) {
	defer func() {
		result *= 10
	}()
	return 6
}
func test3() (result int) {
	result = 8
	defer func() {
		result *= 10
	}()
	return
}
func main() {
	result1 := test1()
	result2 := test2()
	result3 := test3()
	fmt.Println(result1, result2, result3)
}
```

- A: 0 6 8 
- B: 0 60 80
- C: 10 6 80
- D: 10 60 8
- E: 编译报错
#### 60和80是闭包的结果，0不清楚
###点评：
这道题主要考察以下知识点：

* 在被defer的函数里对返回值做修改在什么情况下会生效？

   

### 详解

我们来看下官方文档里的规定：

> Each time a "defer" statement executes, the function value and parameters to
> the call are evaluated as usual and saved anew but the actual function is not 
> invoked. Instead, deferred functions are invoked immediately before the 
> surrounding function returns, in the reverse order they were deferred. That
> is, if the surrounding function returns through an explicit return statement, 
> deferred functions are executed after any result parameters are set by that 
> return statement but before the function returns to its caller. If a deferred
> function value evaluates to nil, execution panics when the function is 
> invoked, not when the "defer" statement is executed.
重点是这句：

> That is, if the surrounding function returns through an explicit return statement, 
> deferred functions are executed after any result parameters are set by that 
> return statement but before the function returns to its caller.
Go的规定是：如果在函数A里执行了 defer B(xx)，函数A显式地通过return语句来返回时，会先把返回值赋值给A的返回值参数，然后执行被defer的函数B，最后才真正地返回给函数A的调用者。

对于`test1`函数，执行`return i`时，先把`i`的值赋值给`test1`的返回值，defer语句里对`i`的赋值并不会改变函数`test1`的返回值，`test1`函数返回0。

对于`test2`函数，执行`return i`时，先把`i`的值赋值给`test2`的命名返回值result，defer语句里对`result`的修改会改变函数`test2`的返回值，`test2`函数返回60。

对于`test3`函数，虽然`return`后面没有具体的值，但是编译器不会报错，执行`return`时，先执行被defer的函数，在被defer的函数里对result做了修改，result的结果变为80，最后`test3`函数return返回的时候返回80。

所以答案是B。

**所以想要对return返回的值做修改，必须使用命名返回值(Named Return Value)**。

## 递归
Redhat的首席工程师、Prometheus开源项目Maintainer [Bartłomiej Płotka](https://twitter.com/bwplotka) 在Twitter上出了一道Go编程题，结果超过80%的人都回答错了。

题目如下所示，回答下面这段程序的输出结果。

```go
// named_return.go
package main
import "fmt"
func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}
func bbb() (done func(), _ error) {
	done, err := aaa()
	return func() { print("bbb: surprise!"); done() }, err
}
func main() {
	done, _ := bbb()
	done()
}
```

* A: `bbb: surprise!`
* B: `bbb: surprise!aaa: done`
* C: 编译报错
* D: 递归栈溢出

#### 猜的递归溢出，不知道原因
## 解析

在函数`bbb`最后执行return语句，会对返回值变量`done`进行赋值，

```go
done := func() { print("bbb: surprise!"); done() }
```

**注意**：闭包`func() { print("bbb: surprise!"); done() }`里的`done`并不会被替换成`done, err := aaa()`里的`done`的值。

因此函数`bbb`执行完之后，返回值之一的`done`实际上成为了一个递归函数，先是打印`"bbb: surprise!"`，然后再调用自己，这样就会陷入无限递归，直到栈溢出。因此本题的答案是`D`。

那为什么函数`bbb`最后return的闭包`func() { print("bbb: surprise!"); done() }`里的`done`并不会被替换成`done, err := aaa()`里的`done`的值呢？如果替换了，那本题的答案就是`B`了。



这个时候就要搬出一句老话了：

> This is a feature, not a bug


我们可以看下面这个更为简单的例子，来帮助我们理解：

```go
// named_return1.go
package main
import "fmt"
func test() (done func()) {
	return func() { fmt.Println("test"); done() }
}
func main() {
	done := test()
	// 下面的函数调用会进入死循环，不断打印test
	done()
}
```

正如上面代码里的注释说明，这段程序同样会进入无限递归直到栈溢出。

如果函数`test`最后return的闭包`func() { fmt.Println("test"); done() }`里的`done`是被提前解析了的话，因为`done`是一个函数类型的变量，`done`的零值是`nil`，那闭包里的`done`的值就会是`nil`，执行`nil`函数是会引发panic的。

**但实际上Go设计是允许上面的代码正常执行的(通过这种方式可以返回一个递归函数)**，因此函数`test`最后return的闭包里的`done`的值并不会提前解析，`test`函数执行完之后，实际上产生了下面的效果，返回的是一个递归函数，和本文开始的题目一样。

```go
done := func() { fmt.Println("test"); done() }
```

因此也会进入无限递归，直到栈溢出。



##  总结

这个题目其实很tricky，在实际编程中，要避免对命名返回值采用这种写法，非常容易出错。

想了解国外Go开发者对这个题目的讨论详情可以参考[Go Named Return Parameters Discussion](https://twitter.com/bwplotka/status/1494362886738780165)。

另外题目作者也给了如下所示的解释，原文地址可以参考[详细解释](https://go.dev/play/p/ELPEi2AK0DP)：

```go
package main
func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}
func bbb() (done func(), _ error) {
	// NOTE(bwplotka): Here is the problem. We already defined special "return argument" variable called "done".
	// By using `:=` and not `=` we define a totally new variable with the same name in
	// new, local function scope.
	done, err := aaa()
	// NOTE(bwplotka): In this closure (anonymous function), we might think we use `done` from the local scope,
	// but we don't! This is because Go "return" as a side effect ASSIGNS returned values to
	// our special "return arguments". If they are named, this means that after return we can refer
	// to those values with those names during any execution after the main body of function finishes
	// (e.g in defer or closures we created).
	//
	// What is happening here is that no matter what we do in the local "done" variable, the special "return named"
	// variable `done` will get assigned with whatever was returned. Which in bbb case is this closure with
	// "bbb:surprise" print. This means that anyone who runs this closure AFTER `return` did the assignment
	// will start infinite recursive execution.
	//
	// Note that it's a feature, not a bug. We use this often to capture
	// errors (e.g https://github.com/efficientgo/tools/blob/main/core/pkg/errcapture/doc.go)
	//
	// Go compiler actually detects that `done` variable defined above is NOT USED. But we also have `err`
	// variable which is actually used. This makes compiler to satisfy that unused variable check,
	// which is wrong in this context..
	return func() { print("bbb: surprise!"); done() }, err
}
func main() {
	done, _ := bbb()
	done()
}
```

不过这个解释是有瑕疵的，主要是这句描述：

> By using `:=` and not `=` we define a totally new variable with the same name in
>  new, local function scope.
对于`done, err := aaa()`，返回变量`done`并不是一个新的变量，而是和函数`bbb`的返回变量`done`是同一个变量。
