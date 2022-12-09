package main

import (
	"fmt"
	"math"
)

func main() {
	//	练习题
	//	1.判断 100300 - 100500之间有多少个素数，并输出所有素数。
	for i := 100300; i <= 100500; i++ {
		flag_ := 1
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag_ = 0
			}
		}
		if flag_ == 1 {
			fmt.Println("素数为： ", i)
		}
	}
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
	for i := 100; i < 10000000; i++ {
		flag_ := 0
		flag_ = Is_zimishu(i)
		if flag_ == 1 {
			fmt.Println("自幂数为：  ", i)
		}
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
