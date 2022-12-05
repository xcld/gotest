package main

import (
	"fmt"
)

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
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	//点评：
	//样式1
	//for i := 1; i < 101; i++ {
	//	if i%3 == 0 && i%5 == 0 {
	//		fmt.Printf("FizzBuzz,")
	//	} else if i%3 == 0 {
	//		fmt.Printf("Fizz,")
	//	} else if i%5 == 0 {
	//		fmt.Printf("Buzz,")
	//	} else {
	//		fmt.Printf("%d ,", i)
	//	}
	//}
	// 1 ,2 ,Fizz,4 ,Buzz,Fizz,7 ,8 ,Fizz,Buzz,11 ,Fizz,13 ,14 ,FizzBuzz,16 ,17 ,Fizz,19 ,Buzz,Fizz,22 ,23 ,Fizz,Buzz,26 ,Fizz,28 ,29 ,FizzBuzz,31 ,32 ,Fizz,34 ,Buzz,Fizz,37 ,38 ,Fizz,Buzz,41 ,Fizz,43 ,44 ,FizzBuzz,46 ,47 ,Fizz,49 ,Buzz,Fizz,52 ,53 ,Fizz,Buzz,56 ,Fizz,58 ,59 ,FizzBuzz,61 ,62 ,Fizz,64 ,Buzz,Fizz,67 ,68 ,Fizz,Buzz,71 ,Fizz,73 ,74 ,FizzBuzz,76 ,77 ,Fizz,79 ,Buzz,Fizz,82 ,83 ,Fizz,Buzz,86 ,Fizz,88 ,89 ,FizzBuzz,91 ,92 ,Fizz,94 ,Buzz,Fizz,97 ,98 ,Fizz,Buzz,
	//方法2：
	//const (
	//	FIZZ = 3
	//	BUZZ = 5
	//) //定义两常量增加代码可读性
	//var Bu bool
	//for i := 1; i < 100; i++ {
	//	Bu = false
	//	if i%FIZZ == 0 {
	//		fmt.Printf("Fizz")
	//		Bu = true
	//	}
	//	if i%BUZZ == 0 {
	//		fmt.Printf("Buzz")
	//		Bu = true
	//	}
	//	if !Bu { //上都不符合原值输出
	//		fmt.Printf("%v", i)
	//	}
	//	fmt.Println() //换行
	//}

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
	//for i := 1; i < 101; i++ {
	//	for j := 1; j <= i; j++ {
	//		st := "A"
	//		fmt.Print(st)
	//	}
	//	fmt.Println("", i)
	//	//fmt.Println("这里是i的次数：", i)
	//}

	// 法二：
	var sum = 0 // 想清楚是从零开始还是从1开始
	for i := 1; i < 101; i++ {
		for j := 1; j <= i; j++ {
			st := "A"
			fmt.Print(st)
			sum++
			if sum == 100 {
				goto loop2
			}
		}
		fmt.Println("  ", i)

	}
loop2:
	{
	}

}

//点评：
//上面代码输出少了个A 一共才99个
//A 2
//AA 4
//AAA 7
//AAAA 11
//AAAAA 16
//AAAAAA 22
//AAAAAAA 29
//AAAAAAAA 37
//AAAAAAAAA 46
//AAAAAAAAAA 56
//AAAAAAAAAAA 67
//AAAAAAAAAAAA 79
//AAAAAAAAAAAAA 92
//AAAAAAAA
// 测试提交
