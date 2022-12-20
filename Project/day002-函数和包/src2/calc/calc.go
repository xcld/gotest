package calc

func Add(a, b int) int { //加
	return a + b
}

func Minus(a, b int) int { //减
	return a - b
}

func Multiply(a, b int) int { //乘
	return a * b
}

func Divide(a, b int) int { //除
	return a / b
}

type Ad func(int, int) int

func abc(a, b int, f Ad) (result int) {
	result = f(a, b)
	return
}
