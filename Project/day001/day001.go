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

func main() {
	Scale := float32(5 / 8)
	fmt.Printf("国际时装周模特%d标准为:\n%dcm的身高\n，%dkg的体重\n，三围：\n胸围为：%fcm\n,腰围为：%fcm\n,臀围为：%fcm\n，并且上下身比例为：%f\n", Gender, Height, Weight, Bust, Waist, Hip, Scale)
	//fmt.Printf("It's over %f\n", Power)
	fmt.Println(reflect.TypeOf(Scale))
}
