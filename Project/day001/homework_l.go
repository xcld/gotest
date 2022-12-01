package main

import "fmt"

func main() {

	//Q1. (0) For-loop--第26页
	//1.创建一个基于 for 的简单的循环。使其循环 10 次，并且使用fmt 包打印出计数器的值。
	for i := 1; i <= 10; i++ {
		fmt.Println("i为：", i)
	}
	//2.用goto 改写1的循环。关键字 for 不可使用。
	fmt.Println("goto test")
	var x = 0
loop:
	if x < 10 {
		fmt.Println("x的goto循环: ", x)
		x++
		goto loop
	}
	//3.再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println("arr为：", arr)

	//02.(0) FizzBuzz
	//1.解决这个叫做 Fizz-Buzz[23] 的问题:
	//编写一个程序，打印从 1 到 100 的数字。当是三个倍数就打印“Fizz代智数字，当是五的倍数就打印“Buzz”。 当数字同时是三和五的倍数时，打印“FizzBuzz”。
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz, 这是：", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz, 这是：", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz, 这是：", i)
		}
	}

	//Q3.(1) 字符串
	//1.建立一个 Go 程序打印下面的内容(到 100 个字符):
	//A
	//AA
	//AAA
	//AAAA
	//PAAPA
	//AaAaAa
	//AAAAAAA
	// 法一：low：
	for i := 1; i < 101; i++ {
		for j := 1; j <= i; j++ {
			st := "A"
			fmt.Print(st)
		}
		fmt.Println("", i)
		//fmt.Println("这里是i的次数：", i)
	}

	//for i := 1; i < 101; i++ {
	//	st := "A" * i
	//
	//	fmt.Println(st)
	//}
}

// 测试提交



