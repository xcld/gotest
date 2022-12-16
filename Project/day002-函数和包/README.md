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
3. 变参，编写函数接受整数类型变参，并且每行打印一个数字。

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
6. 程序改错题，下面的程序有什么错误？
```go
func main() {
for i := 0; i < 10; i++ {
fmt.Printf("%v\n", i)
}
fmt.Printf("%v\n", i)
}
```
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
	fmt.Println(f(10)) 
	fmt.Println(f(20)) 
	fmt.Println(f(30)) 

	f1 := adder()
	fmt.Println(f1(40)) 
	fmt.Println(f1(50)) 
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
	fmt.Println(f1(1), f2(2)) 
	fmt.Println(f1(3), f2(4)) 
	fmt.Println(f1(5), f2(6)) 
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
