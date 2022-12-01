// package main
//
// import (
//
//	"fmt"
//	"reflect"
//
// )
//
// var _ int = 18
// var Gender, Height, Weight int = 0, 178, 48
// var (
//
//	Bust  float32 = 90
//	Waist float32 = 60
//	Hip   float32 = 90
//
// )
//
// const Pi float64 = 3.14159
// const e = 2.71828
// const (
//
//	G float64 = 6.7e-11
//	C         = 2.99e8
//
// )
// const (
//
//	Zero = iota
//	Fir
//	Sec
//	Thi
//	Fou01, Fou02, Fou03 = iota, iota, iota
//	Fif                 = "fifth"
//	Six
//	Seventh = iota + 3
//
// )
//
// type OrderStatus int
//
// const (
//
//	Max                   = 5
//	Cancelled OrderStatus = Max - iota
//
// )
//
// const GoldenRatio, Bernstein, HafnerSarnakMcCurley = 0.61803, 0.70258, 4.66920
// const Feigenbaum, C2, M1 float64 = 2.50290, 0.66016, 0.26149
//
// var ZiFu string = "abc\n"
//
//	func main() {
//		NoPay := 1
//		//NoPay := 2
//		ZiFu01 := `abc\n`
//		Scale := float32(5 / 8)
//		fmt.Printf("国际时装周模特%d标准为:\n%dcm的身高\n，%dkg的体重\n，三围：\n胸围为：%fcm\n,腰围为：%fcm\n,臀围为：%fcm\n，并且上下身比例为：%f\n", Gender, Height, Weight, Bust, Waist, Hip, Scale)
//		//fmt.Printf("It's over %f\n", Power)
//		fmt.Println(reflect.TypeOf(Scale))
//		fmt.Printf("圆周率:%f\n自然对数的底:%f\n重力常数:%E\n真空光速:%E", Pi, e, G, C)
//		fmt.Printf("黄金比φ:%f\nEmbree-Trefethen 常数β:%f\n第一费根鲍姆常数δ:%f\n第二费根鲍姆常数α:%f\n", GoldenRatio, Bernstein, HafnerSarnakMcCurley, Feigenbaum)
//		fmt.Printf("孪生质数常数C2:%f\nMeissel-Mertens常数M1:%f\n", C2, M1)
//		fmt.Println(Zero, Fir, Sec, Thi, Fou01, Fou02, Fou03, Fif, Six, Seventh)
//		fmt.Println(Cancelled, Max, NoPay)
//		fmt.Println(ZiFu, ZiFu01)
//	}
//
// package main
//
// import "fmt"
//
// var Iput float64
//
//	func main() {
//		fmt.Println("请输入一个整数：")
//		fmt.Scanf("%f", &Iput)
//		var str = int(Iput)
//		fmt.Println("这个整数是： ", str)
//	}

// package main
//
// import "fmt"
//
//	func main() {
//		var ch1, ch2, ch3 byte
//		ch1 = 'a'
//		ch2 = 97
//		ch3 = '\n'
//		fmt.Printf("ch1= %c,ch2 = %c ,%c", ch1, ch2, ch3)
//		type gaibian int
//		var b int = int(ch1)
//		var Test gaibian = 4 << 2
//		Order := 2 ^ 3 + 1 | 1*2
//		println(b, Test, Order)
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		var _ int = 18
//		var Gender, Height, Weight int = 0, 178, 48
//		var (
//			Bust  float32 = 90
//			Waist float32 = 60
//			Hip   float32 = 90
//		)
//		switch Gender {
//		case 0:
//			fmt.Println("这个模特是女生！")
//		case 1:
//			fmt.Println("这个模特是男生！")
//		}
//		if Height >= 178 {
//			fmt.Println("模特身高合格！")
//		} else if Weight <= 50 {
//			fmt.Println("模特体重合格！")
//		} else if Bust == 90 {
//			fmt.Println("模特胸围合格")
//		} else if Waist == 60 {
//			fmt.Println("模特腰围合格！")
//		} else if Hip == 90 {
//			fmt.Println("模特臀围合格！")
//		} else {
//			fmt.Println("模特不符合要求！")
//		}
//	}
//
// package main
//
// import "fmt"
//
//	func main() {
//		var scord float64
//		fmt.Println("请输入你的成绩：")
//		fmt.Scanf("%f", &scord)
//		if scord >= 90 {
//			fmt.Println("你的成绩等级是A")
//		} else if scord >= 80 {
//			fmt.Println("你的成绩等级是B")
//		} else if scord >= 70 {
//			fmt.Println("你的成绩等级是C")
//		} else if scord >= 60 {
//			fmt.Println("你的成绩等级是D")
//		} else {
//			fmt.Println("你的成绩等级是F")
//		}
//	}
//
// package main
//
// import (
//
//	"fmt"
//
// )
//
//	func main() {
//		const (
//			Apple int = iota
//			Xiaomi
//			Huawei
//			Redmi
//			Meizu
//			Honor
//			Oppo
//		)
//		var InitList = [7]int{Apple, Xiaomi, Huawei, Redmi, Meizu, Honor, Oppo}
//		for index, value := range InitList {
//			fmt.Println("[InitList]index=", index, "value=", value)
//
//			switch value {
//			case Apple:
//				fmt.Println("红")
//			case Xiaomi:
//				fmt.Println("橙")
//			case Huawei:
//				fmt.Println("黄")
//			case Redmi:
//				fmt.Println("绿")
//			case Meizu:
//				fmt.Println("蓝")
//			case Honor:
//				fmt.Println("青")
//			case Oppo:
//				fmt.Println("紫")
//			}
//		}
//	}
//package main
//
//import "fmt"
//
//func main() {
//	var sum, Sum int
//Loop:
//	for i := 1; i <= 14; i++ {
//		for j := 1; j <= i; j++ {
//			if Sum == 7 {
//				break Loop
//			}
//			fmt.Printf("A")
//			Sum += j
//		}
//		sum += i
//		if Sum == 6 {
//			break
//		}
//		//fmt.Println("i=", i, "sum=", sum)
//		fmt.Println("\n", sum)
//	}
//}

// day001
// package main
//
// import (
//
//	"fmt"
//	"reflect"
//
// )
//
// var Day001 int = 001
//
//	func main() {
//		day001 := 1
//		day002 := 1.588888
//		var (
//			Height float64 = 180.6
//			Weight float64 = 65.5
//			Age    int     = 18
//		)
//		const Pi float64 = 3.1415926
//		const (
//			G float64 = 9.8
//			C float64 = 30000
//		)
//		const Abc, Def int = 6, 8
//		const (
//			X = iota + 10
//			Y
//			Z, A, B = iota, iota, iota
//		)
//		var _, Hip, Waigt int = 0, 60, 80
//		var ch1, ch2 byte
//		ch1 = 'a'
//		ch2 = 97
//		var ch3 string = "\n"
//		fmt.Println("hello day001!")
//		fmt.Println(Day001)
//		fmt.Println(day001)
//		fmt.Println(day002)
//		fmt.Println("type day001 is:", reflect.TypeOf(day001))
//		fmt.Println("type day002 is:", reflect.TypeOf(day002))
//		fmt.Println(Height, Weight, Age)
//		fmt.Println(Hip, Waigt)
//		fmt.Printf("圆周率是：%f\n重力加速度：%0.2f\n真空光速：%f\nAbc是：%d\n Def是：%d\n", Pi, G, C, Abc, Def)
//		fmt.Println(X, Y, Z, A, B)
//		fmt.Printf("ch1= %c,ch2=%c ,ch3=%s", ch1, ch2, ch3)
//
// {
//package main
//
//import "fmt"
//
//var Hip int
//
//type Tempc float64
//
//func main() {
//	var F Tempc = 36.5
//	fmt.Println("请输入一个整数：")
//	fmt.Scanf("%d", &Hip)
//	fmt.Println(Hip)
//	fmt.Println(F)
//}

// 协作贡献修改
// package main
//
// import "fmt"
//
//	func main() {
//		var Power, Name, Huawei = 2, 5, 6
//		var Sum, Mo int
//		var Liba, Dufu, Weiyunsuan int = 2, 4, 3
//		Sum = Power + Name - Huawei
//		Sum++
//		Mo = Huawei % Power
//		Guanxi := Power <= Huawei
//
//		//++Sum 无前自增自减
//		Weiyunsuan = Dufu & Liba
//		Weiyunsuan1 := Dufu | Liba
//		Weiyunsuan2 := Dufu ^ Liba
//		Weiyunsuan3 := Dufu << Liba
//		Weiyunsuan4 := 4 >> 1
//		println(Sum)
//		println(Mo)
//		fmt.Printf("Guanxi:%t", Guanxi)
//		fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan:%b", Liba, Dufu, Weiyunsuan)
//		fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan1:%b", Liba, Dufu, Weiyunsuan1)
//		fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan2:%b", Liba, Dufu, Weiyunsuan2)
//		fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan3:%b", Liba, Dufu, Weiyunsuan3)
//		fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan4:%b", Liba, Dufu, Weiyunsuan4)
//	}
//package main
//
//import "fmt"
//
//func main() {
//
//	var Liba, Dufu, Weiyunsuan int = 2, 4, 3
//
//	//++Sum 无前自增自减
//	Weiyunsuan &= Dufu
//	//fmt.Scanf("%d", &Dufu)
//	//Weiyunsuan1 := Dufu | Liba
//	//Weiyunsuan2 := Dufu ^ Liba
//	//Weiyunsuan3 := Dufu << Liba
//	Weiyunsuan4 := 4%2 | 2*2
//
//	fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan4:%b", Liba, Dufu, Weiyunsuan4)
//	fmt.Printf("Liba:%b\nDufu:%b\nWeiyunsuan:%b", Liba, Dufu, Weiyunsuan)
//	fmt.Println(&Dufu)
//}

// package main
//
// import "fmt"
//
//	func main() {
//		var score float64
//		fmt.Println("请输入您的成绩：")
//		fmt.Scanf("%f", &score)
//		if score >= 90 {
//			fmt.Println("优秀")
//		} else if score >= 70 {
//			fmt.Println("合格！")
//		} else {
//			fmt.Println("不合格！")
//		}
//
// }
// package main
//
// import "fmt"
//
//	func main() {
//		////var Score int = 90
//		//switch Score := 90; Score {
//		//case 90:
//		//	fmt.Println("A")
//		//	fallthrough
//		//case 80:
//		//	fmt.Println("B")
//		//case 70, 60:
//		//	fmt.Println("一般！")
//		//default:
//		//	fmt.Println("差")
//		//
//		//}
//		var sum int
//		for i := 1; i <= 100; i++ {
//			sum += i
//		}
//		for sum >= 100 {
//
//			fmt.Println("sum= ", sum)
//			sum--
//		}
//
// }
// package main
//
// import "fmt"
//
//	func main() {
//		const (
//			Apple int = iota
//			Xiaomi
//			Huawei
//			Redmi
//			Meizu
//			Honor
//			Oppo
//		)
//		var InitList = [7]int{Apple, Xiaomi, Huawei, Redmi, Meizu, Honor, Oppo}
//		for index, value := range InitList {
//			fmt.Println("[InitList]index=", index, "value=", value)
//
//			switch value {
//			case Apple:
//				fmt.Println("红")
//			case Xiaomi:
//				fmt.Println("橙")
//			case Huawei:
//				fmt.Println("黄")
//			case Redmi:
//				fmt.Println("绿")
//			case Meizu:
//				fmt.Println("蓝")
//			case Honor:
//				fmt.Println("青")
//			case Oppo:
//				fmt.Println("紫")
//			}
//		}
//	}
//
// package main
//
// import "fmt"
//
// func main() {
//
//	//var sum int
//	for i := 0; i <= 5; i++ {
//		for {
//			fmt.Println(i)
//			goto Loop
//		}
//	}
//	fmt.Println("hhhhh")
//
// Loop:
//
//	fmt.Println("switc Loop")
//
// }
package main

import "fmt"

func main() {
p:
	for i := 0; i < 2; i++ {
		for j := 0; i < 2; j++ {
			print(i, " ", j, " ")
			break p
		}
	}
	fmt.Println("外面世界真好！")
}
