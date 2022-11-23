package main

import (
	"fmt"
	"reflect"
)

var _ int = 18
var Gender, Height, Weight int = 0, 178, 48
var (
	Bust  float32 = 90
	Waist float32 = 60
	Hip   float32 = 90
)

const Pi float64 = 3.14159
const e = 2.71828
const (
	G float64 = 6.7e-11
	C         = 2.99e8
)
const (
	Zero = iota
	Fir
	Sec
	Thi
	Fou01, Fou02, Fou03 = iota, iota, iota
	Fif                 = "fifth"
	Six
	Seventh = iota + 3
)

type OrderStatus int

const (
	Cancelled OrderStatus = iota
)

const GoldenRatio, Bernstein, Hafner_Sarnak_McCurley = 0.61803, 0.70258, 4.66920
const Feigenbaum, C2, M1 float64 = 2.50290, 0.66016, 0.26149

func main() {
	NoPay := 1
	Scale := float32(5 / 8)
	fmt.Printf("国际时装周模特%d标准为:\n%dcm的身高\n，%dkg的体重\n，三围：\n胸围为：%fcm\n,腰围为：%fcm\n,臀围为：%fcm\n，并且上下身比例为：%f\n", Gender, Height, Weight, Bust, Waist, Hip, Scale)
	//fmt.Printf("It's over %f\n", Power)
	fmt.Println(reflect.TypeOf(Scale))
	fmt.Printf("圆周率:%f\n自然对数的底:%f\n重力常数:%E\n真空光速:%E", Pi, e, G, C)
	fmt.Printf("黄金比φ:%f\nEmbree-Trefethen 常数β:%f\n第一费根鲍姆常数δ:%f\n第二费根鲍姆常数α:%f\n", GoldenRatio, Bernstein, Hafner_Sarnak_McCurley, Feigenbaum)
	fmt.Printf("孪生质数常数C2:%f\nMeissel-Mertens常数M1:%f\n", C2, M1)
	fmt.Println(Zero, Fir, Sec, Thi, Fou01, Fou02, Fou03, Fif, Six, Seventh)
	fmt.Println(Cancelled, NoPay)
}
