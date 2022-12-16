package main

// // 无参无返回
//
//	func No_two() {
//		fmt.Println("这里是无参无返回")
//	}
//
// // 有参无返回
//
//	func fan_can(a int) {
//		fmt.Println("有参无返回参数a： ", a)
//	}
//
// // 有参有返回
//
//	func two_you(a int) (b int) {
//		b = 10
//		b += a
//		return b
//	}
//
// // 可接受任意个数参数
//
//	func any_arg(args ...int) {
//		for _, n := range args {
//			fmt.Println("任意个数的参数n： ", n)
//		}
//	}
//
// // 递归函数，函数自己直接或间接调用自己；注意要有递归终止条件
//
//	func digui(num int) int {
//		if num == 1 {
//			return 1
//		}
//		return num + digui(num-1)
//	}
//
// // 给函数起别名
// type liangshu_add func(int, int) int // 声明一个函数的类型
//
//	func Add(a, b int) int {
//		return a + b
//	}
//
// // 局部作用域
//
//	func jubu_bianliang(a int) int {
//		a += 10
//		return a
//	}
//
// var b int = 100
// var c int = 9
//
// // 全局变量
//
//	func quanjun_bianliang() int {
//		b -= 1
//		return b
//	}
//
// // 测试defer的特性，先传参，值传递，最后才运行
//
//	func defer_test(b int) int {
//		fmt.Println("这里是defer_test的b，和c  ", b, "  -- ", c)
//		return b
//	}
//
//	func main() {
//		No_two()
//		fan_can(111)
//		any_arg(1, 2, 3, 4)
//		fmt.Println("有参有返回  ", two_you(10))
//		any_arg(1, 2)
//
//		// 递归示例
//		fmt.Println("递归演示： ", digui(100))
//
//		var f liangshu_add = Add
//		fmt.Println("两数和：", f(1, 2))
//
//		// 局部变量测试：
//		var a int = 2
//		fmt.Println("这里是局部的a： ", jubu_bianliang(a))
//		fmt.Println("这里是main的a： ", a)
//
//		// 全局变量
//		fmt.Println("全局的main的b： ", b)
//		fmt.Println("全局变量的b： ", quanjun_bianliang())
//		fmt.Println("全局的main====的b： ", b)
//
//		// 匿名函数
//		func() {
//			i := 12
//			fmt.Println("这是匿名函数的i： ", i)
//		}() // 直接调用
//
//		// defer延迟调用
//		defer fmt.Println("BBBBBBBBBBB") // 多个defer以栈的顺序来运行
//		fmt.Println("ccccccccc")
//		fmt.Println("全局的main  开始 的b： ", b)
//		defer fmt.Println("defer运行后的 b： ", defer_test(b)) // 把当前的b传进函数了，只是还没有运行，不是传最后的b
//		b += 1000
//		fmt.Println("最终运行的 b： ", b)
//		fmt.Println("AAAAAAAA")
//
//		// 获取命令行参数
//		terminal_args := os.Args // 获取用户输入的多个参数，以空格区分
//		// 判空和值个数不够两个，提示重新输入
//		if terminal_args == nil || len(terminal_args) < 2 {
//			fmt.Println("err: 2个， ip port")
//			return
//		}
//
//		more_ren := terminal_args[0]
//		ip := terminal_args[1]
//		port := terminal_args[2]
//		fmt.Println("输出more_ren： ", more_ren)
//		fmt.Println("这里是接收的命令行参数：  ip = %s,  port = %s\n： ", ip, port)
//
// }
