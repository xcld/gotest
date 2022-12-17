package main

import (
	"fmt"
	"os"
	"strconv"
)

// Fahrenheit scale
//Thermodynamic temperature Kelvin  absolute temperature
//Centigrade scale   Anders Celsius

//	i := len(auto) - 1
//	//k := len(auto)
//	str := auto[:i]
//	str2, _ := strconv.ParseFloat(str, 64)
//	fmt.Printf("%v %v %v\n", auto[:i], auto[i:], len(auto))
//	fmt.Printf("%T %T %T\n", auto[:i], auto[i:], len(auto))
//	fmt.Printf("%v %T", str2, str2)

func Centigrade(Fah float64) (Cen float64) {
	Fah = Cen*9/5 + 32
	fmt.Printf("摄氏温度为:%0.2f℃", Cen)
	return Cen
}

func Fahrenheit(Cen float64) (Fah float64) {
	Cen = (Fah - 32) * 5 / 9
	fmt.Printf("华氏温度为:%0.2fF", Fah)
	return Fah
}

func Thermodynamic(Unit string, Scale float64) (Kel float64) {
	if Unit == "c" || Unit == "C" {
		Kel = Scale + 273.15
	} else if Unit == "f" || Unit == "F" {
		Kel = (Scale-32)*5/9 + 273.15
	} else {
		println("error!,please try again")
	}
	fmt.Printf("开尔文温度为:%0.2fK", Kel)
	return Kel
}
func AutoString(Str string) (Unit string, Scale float64) {
	i := len(Str) - 1
	Str2, _ := strconv.ParseFloat(Str, 64)
	Unit = Str[i:]
	Scale = Str2
	return Unit, Scale
}
func String2Float(Str1 string) (Scale1 float64) {
	Str2, _ := strconv.ParseFloat(Str1, 64)
	Scale1 = Str2
	return Scale1
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
	default:
		println("输入 convert.exe -F 摄氏温度 转换为华氏温度\nconvert.exe -h/--help 显示帮助信息，输入错误或默认情况打印\nconvert.exe -C 华氏温度 转换为摄氏温度\nconvert.exe -K 摄氏/华氏温度 单位用大小写c/f区分，自动转换为相应开尔文温度，其他情况输入错误，请重新输入。\nconvert.exe -A 摄氏/华氏温度单位用大小写c/f区分，根据当前温度自动高温或低温告警或适宜温度\nexample: convert.exe -C 97.70 输出 36.50℃,convert.exe -K 37c 输出 310.15K")
	}
}
