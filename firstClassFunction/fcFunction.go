package firstClassFunction

import "fmt"

func MyFcFunc() {
	// 1
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
	// 2
	func() {
		fmt.Println("hello world first class function again")
	}()
	// 3
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
	// 4
	var c add = func(a int, b int) int {
		return a + b
	}
	s := c(5, 6)
	fmt.Println("Sum", s)
	// 5
	//高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：
	//接收一个或多个函数作为参数
	//返回值是一个函数
	f := func(a, b int) int {
		return a + b
	}
	simple1(f)
	s2 := simple2()
	fmt.Println(s2(60, 7))
	// 6
	// 闭包（Closure）是匿名函数的一个特例。
	// 当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
	d := appendStr()
	fmt.Println(d("World"))
	fmt.Println(d("Gopher"))

}

// 正如我们定义自己的结构体类型一样，我们可以定义自己的函数类型。
type add func(a int, b int) int

func simple1(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}
func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
