package main

import "fmt"

func sort_two_num(a, b int) (c, d int) {
	if a >= b {
		c, d = b, a
	} else {
		c, d = a, b
	}
	return c, d
}

func any_arg(args ...int) {
	for _, n := range args {
		fmt.Println("任意个数的参数n： ", n)
	}
}

func plus2(a int) (b int) {
	return a + 2
}

func plusX(a, x int) (b int) {
	return a + x
}

func add(x, y int) int {
	return x + y
}

//func calc(x, y int, op func(int, int) int) int {
//	return op(x, y)
//}

// 程序1
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 程序2
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

type T int

func (t T) M(n int) T {
	print(n)
	return t
}

type Add func(int, int) int

func main() {
	var f Add
	defer f(1, 2)
	fmt.Println("end")

	//var t T
	//defer t.M(1).M(2)
	//t.M(3).M(4)

	//f1, f2 := calc(10)
	//fmt.Println(f1(1), f2(2)) // 此时也是闭包的调用，所以base的起始值为10且不会被释放掉，11 9
	//fmt.Println(f1(3), f2(4)) // 12 8
	//fmt.Println(f1(5), f2(6)) //13 9

	//var f = adder()
	//fmt.Println(f(10)) //此时是闭包，内部函数对外部函数的变量的调用，所以x不会被释放掉，10
	//fmt.Println(f(20)) // 30
	//fmt.Println(f(30)) //60
	//
	//f1 := adder()
	//fmt.Println(f1(40)) // 此时是新建了一个闭包所以，这里x又从0开始了,40
	//fmt.Println(f1(50)) // 90

	//c, d := sort_two_num(7, 2)
	//fmt.Println(c, d)

	//p := plus2
	//fmt.Println("%v", p(4))
	//fmt.Println("%v\n", plusX(4, 5))

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v\n", i)
	//}
	//fmt.Printf("%v\n", i)
}
