//命令行获取参数实现,示范代码1：

package main

import (
	"fmt"
	"os"
	"strconv"
)

// 保留两位小数
func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func main() {
	//s, sep := "", ""
	//for index, arg := range os.Args[1:] {
	//	s += sep + arg
	//	fmt.Println(index, arg)
	//}
	//fmt.Println("++++++", s)

	st, temp, typ := "", "", ""
	for index, arg := range os.Args[1:] {

		if index == 0 && (arg == "-F" || arg == "-C" || arg == "-h" || arg == "-K" || arg == "-A") {
			st = arg
		} else if index == 1 {
			temp = arg
		} else if index == 2 {
			typ = arg
		}
		//fmt.Println(index, arg, reflect.TypeOf(arg))
	}
	FloatValue, _ := strconv.ParseFloat(temp, 64)
	//fmt.Printf("%T,%f \n", FloatValue, FloatValue)

	if st == "-F" { // go run .\convert.go -F 37.5
		fmt.Printf("摄氏%d,  转华氏%d：", temp, Decimal((FloatValue*9/5)+32))
	} else if st == "-C" { // go run .\convert.go -C 97.5
		fmt.Printf("华氏%d,  转摄氏%d：", temp, Decimal((FloatValue-32.0)*5/9))
	} else if st == "-h" {
		fmt.Println("帮助")
	} else if st == "-K" {
		if typ == "c" { // go run .\convert.go -K 37.5 c
			fmt.Printf("摄氏%d,  转开尔文%d：", temp, Decimal(FloatValue+273.15))
		} else if typ == "f" {
			fmt.Printf("华氏%d,  转开尔文%d：", temp, Decimal((FloatValue*9/5)+32+273.15))
		}
	} else if st == "-A" { //go run .\convert.go -A 37.5 c
		var floa float64
		if typ == "c" {
			floa = Decimal(FloatValue)
		} else if typ == "f" {
			floa = Decimal((FloatValue - 32) * 5 / 9)
		}
		if floa <= -3.0 {
			fmt.Println("低温黄色预警信号")
		} else if floa <= -6.0 {
			fmt.Println("低温橙色预警信号")
		} else if floa >= 40.0 {
			fmt.Println("高温红色预警信号")
		} else if floa >= 37.0 {
			fmt.Println("高温橙色预警信号")
		} else if floa >= 35.0 {
			fmt.Println("高温黄色预警信号")
		}
	} else {
		fmt.Println("输入不正确")
	}

}

// // os.Args变量是一个字符串（string）的切片（slice）,类似python的切片  s[m:n]这个切片，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素。os.Args输出的是执行文件信息
// // PS D:\GO_workspace\BASIC> go run .\命令行参数.go    111 222 333 444 555
// // 0 C:\Users\XUE~1.XIO\AppData\Local\Temp\go-build1439121805\b001\exe\命令行参数.exe
// // 1 111
// // 2 222
// // 3 333
// // 4 444
//
