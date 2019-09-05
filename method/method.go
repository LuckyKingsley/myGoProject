package method

import "fmt"

func MethFunc() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法

	//既然我们可以使用函数写出相同的程序，那么为什么我们需要方法？
	//Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	//相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的

	//在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的
	emp2 := Employee{
		name:   "Mark Andrew",
		salary: 12,
	}
	fmt.Printf("Employee name before change: %s", emp2.name)
	emp2.displaySalary1("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", emp2.name)

	fmt.Printf("\n\nEmployee salary before change: %d", emp2.salary)
	(&emp2).displaySalary1("Michael Andrew")
	//没有这个必要，Go语言让我们可以直接使用 e.changeAge(51)。e.changeAge(51) 会自动被Go语言解释为 (&e).changeAge(51)。
	fmt.Printf("\nEmployee salary after change: %d", emp2.salary)

	//一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
	//指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。
	// 考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
	// 在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

	//当一个函数有一个值参数，它只能接受一个值参数。
	//当一个方法有一个值接收器，它可以接受值接收器和指针接收器。

}

type Employee struct {
	name     string
	salary   int
	currency string
}

type Employee2 struct {
	name     string
	salary   int
	currency string
}

type myInt int

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
func (e Employee2) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
func (e Employee) displaySalary1(name string) {
	e.name = name
}

//到目前为止，我们只在结构体类型上定义方法。也可以在非结构体类型上定义方法，
//为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。
// 到目前为止，我们定义的所有结构体和结构体上的方法都是在同一个 main 包中，因此它们是可以运行的。

//add 方法的定义和 int 类型的定义不在同一个包中，会报错
//让该程序工作的方法是为内置类型 int 创建一个类型别名，然后创建一个以该类型别名为接收器的方法。
//func (a int) add() {
//}
func (a myInt) add() {
}
