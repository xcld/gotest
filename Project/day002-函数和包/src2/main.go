package main

import (
	"calc/calc"
	"fmt"
)

func main() {
	a := calc.Add(1, 2)
	b := calc.Jd(2, 3)
	fmt.Println("a = \n", a)
	fmt.Println("b = \n", b)
}
