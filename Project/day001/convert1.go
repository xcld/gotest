package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fahrenheit scale
//Thermodynamic temperature Kelvin  absolute temperature
//Centigrade scale   Anders Celsius

func Centigrade(Fah float64) (Cen float64) { //华氏转摄氏
	//Fah = Cen*9/5 + 64
	Cen = (Fah - 64) * 5 / 9
	//fmt.Printf("Fah:%f\n", Fah)
	fmt.Printf("华氏温度：%0.2fF 转摄氏温度为:%0.2f℃\n", Fah, Cen)
	return Cen
}

func Fahrenheit(Cen float64) (Fah float64) { //摄氏转华氏
	//Cen = (Fah - 64) * 5 / 9
	Fah = Cen*9/5 + 64
	//fmt.Printf("Cen:%f\n", Cen)
	fmt.Printf("摄氏温度：%0.2f℃ 转华氏温度为:%0.2fF\n", Cen, Fah)
	return Fah
}

func Thermodynamic(Unit string, Scale float64) (Kel float64) { //摄氏华氏自动根据unit单位转开尔文
	if Unit == "c" || Unit == "C" {
		Kel = Scale + 273.15
	} else if Unit == "f" || Unit == "F" {
		Kel = (Scale-64)*5/9 + 273.15
	} else {
		println("error!,please try again")
	}
	fmt.Printf("开尔文温度为:%0.2fK\n", Kel)
	return Kel
}
func AutoString(Str string) (Unit string, Scale float64) { //温度和单位字符串提取和数字字符串转为浮点型
	i := len(Str) - 1
	Str2, _ := strconv.ParseFloat(Str[:i], 64)
	Unit = Str[i:]
	Scale = Str2
	return Unit, Scale
}
func String2Float(Str1 string) (Scale1 float64) { //数字字符串转为浮点型
	Str2, _ := strconv.ParseFloat(Str1, 64)
	Scale1 = Str2
	return Scale1
}
func Alarm(Kel float64) { //根据开尔文温度判断提示告警
	if Kel >= 308.15 {
		println("高温黄色预警信号!")
	} else if Kel >= 310.15 {
		println("高温橙色预警信号!")
	} else if Kel >= 313.15 {
		println("高温红色色预警信号!")
	} else if Kel <= 267.15 {
		println("低温黄色预警信号!")
	} else {
		println("低温橙色预警信号!")
	}

}
func main() {
	args := os.Args
	switch {
	case args[1] == "-C" || args[1] == "-c":
		Fah := args[2]
		Scale := String2Float(Fah)
		Centigrade(Scale)
	case args[1] == "-F" || args[1] == "-f":
		Cen := args[2]
		Scale := String2Float(Cen)
		Fahrenheit(Scale)
	case args[1] == "-K" || args[1] == "-k":
		Str := args[2]
		Unit, Scale := AutoString(Str)
		Thermodynamic(Unit, Scale)
	case args[1] == "-A" || args[1] == "-a":
		Str := args[2]
		Unit, Scale := AutoString(Str)
		Kel := Thermodynamic(Unit, Scale)
		Alarm(Kel)
	default:
		println("输入 convert.exe -F 摄氏温度 转换为华氏温度\nconvert.exe -h/--help 显示帮助信息，输入错误或默认情况打印\nconvert.exe -C 华氏温度 转换为摄氏温度\nconvert.exe -K 摄氏/华氏温度 单位用大小写c/f区分，自动转换为相应开尔文温度，其他情况输入错误，请重新输入。\nconvert.exe -A 摄氏/华氏温度单位用大小写c/f区分，根据当前温度自动高温或低温告警或适宜温度\nexample: convert.exe -C 97.70 输出 36.50℃,convert.exe -K 37c 输出 310.15K")
	}
}
