package main

import (
	"fmt"
	"math"
)

func main() {
	//	练习题
	//	1.判断 100300 - 100500之间有多少个素数，并输出所有素数。
	//for i := 100300; i <= 100500; i++ {
	//	flag_ := 1
	//	for j := 2; j < i; j++ {
	//		if i%j == 0 {
	//			flag_ = 0
	//		}
	//	}
	//	if flag_ == 1 {
	//		fmt.Println("素数为： ", i)
	//	}
	//}
	fmt.Println()
	fmt.Println()
	//	2.打印出100-1000000000 中所有的”自幂数”。所谓“自幂数”在一个固定的进制中，一个n位自然数等于自身各个数位上数字的n次幂之和
	//	则称此数为自幂数。 例如: 1^3 + 5^3+ 3^3 = 153。并按位数输出展示例如
	//三位位的水仙花数共有4个: 153，378，371，487;
	//四位的四叶玫瑰数共有3个:1634，828，9474;
	//五位的五角星数共有3个: 54748，92727，93884;六位的六合数只有1个:548834;
	//七位的北斗七星数共有4个: 1741725，4210818，9888817，9926315;
	//八位的八仙数共有3个: 24678850，24678851，88593477;
	//九位的九九重阳数共有4个: 146511208，472335975，534494836，912985153;

	//for i := 100; i < 10000000; i++ {
	//	flag_ := 0
	//	flag_ = Is_zimishu(i)
	//	if flag_ == 1 {
	//		fmt.Println("自幂数为：  ", i)
	//	}
	//}

	//3. 写一个计算n的阶乘之和程序并输出
	//var v int
	//fmt.Println("请输入一个整数: ")
	//fmt.Scanf("%d", &v)
	//sum := 1
	//for i := 1; i <= v; i++ {
	//	sum *= i
	//}
	//fmt.Println("阶乘的和为： ", sum)

	//4. 任意输入三条边，并判断能否组成三角形，若能，判断三角形的类型和计算三角形的面积。
	//两条较短边的平方和大于最长边的平方，此三角形就是锐角三角形；
	//两条较短边的平方和小于最长边的平方，此三角形就是钝角三角形；
	//两条边短边的平方和等于最长边的平方，此三角形就是直角三角形.
	var a, b, c int
	fmt.Println("请输入三条边: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	c = a
	fmt.Println(a, b, c)
	max, mid, min := a, a, a
	if b >= max {
		max = b
	} else if max >= b && b >= min {
		mid = b
		b = mid
	} else {
		min = b
	}

	if c >= max {
		max = c
	} else if max >= c && c >= min {
		mid = c
	} else {
		min = c
	}

	fmt.Println(min, mid, max)

	if (a+b) > c && (a+c) > b && (c+b) > a {
		fmt.Println("是三角形")

	} else {
		fmt.Println("不能构成三角形")
	}

}

// 判断是否是自幂数
func Is_zimishu(num int) int {
	x := 0
	if num >= 1000000 {
		gewei := num % 10
		shiwei := num / 10 % 10
		baiwei := num / 100 % 10
		qianwei := num / 1000 % 10
		wanwei := num / 10000 % 10
		shiwanwei := num / 100000 % 10
		baiwanwei := num / 1000000
		if (int(math.Pow(float64(qianwei), 7)) + int(math.Pow(float64(baiwei), 7)) + int(math.Pow(float64(shiwei), 7)) + int(math.Pow(float64(gewei), 7)) +
			int(math.Pow(float64(wanwei), 7)) + int(math.Pow(float64(shiwanwei), 7)) + int(math.Pow(float64(baiwanwei), 7))) == num {
			x = 1
		}
	} else if num >= 100000 {
		gewei := num % 10
		shiwei := num / 10 % 10
		baiwei := num / 100 % 10
		qianwei := num / 1000 % 10
		wanwei := num / 10000 % 10
		shiwanwei := num / 100000
		if (int(math.Pow(float64(qianwei), 6)) + int(math.Pow(float64(baiwei), 6)) + int(math.Pow(float64(shiwei), 6)) + int(math.Pow(float64(gewei), 6)) +
			int(math.Pow(float64(wanwei), 6)) + int(math.Pow(float64(shiwanwei), 6))) == num {
			x = 1
		}
	} else if num >= 10000 {
		gewei := num % 10
		shiwei := num / 10 % 10
		baiwei := num / 100 % 10
		qianwei := num / 1000 % 10
		wanwei := num / 10000
		if (int(math.Pow(float64(qianwei), 5)) + int(math.Pow(float64(baiwei), 5)) + int(math.Pow(float64(shiwei), 5)) + int(math.Pow(float64(gewei), 5)) + int(math.Pow(float64(wanwei), 5))) == num {
			x = 1
		}
	} else if num >= 1000 {
		gewei := num % 10
		shiwei := num / 10 % 10
		baiwei := num / 100 % 10
		qianwei := num / 1000
		if (int(math.Pow(float64(qianwei), 4)) + int(math.Pow(float64(baiwei), 4)) + int(math.Pow(float64(shiwei), 4)) + int(math.Pow(float64(gewei), 4))) == num {
			x = 1
		}
	} else if num >= 100 {
		gewei := num % 10
		shiwei := num / 10 % 10
		baiwei := num / 100
		if (baiwei*baiwei*baiwei + shiwei*shiwei*shiwei + gewei*gewei*gewei) == num {
			x = 1
		}
	}
	return x

}
