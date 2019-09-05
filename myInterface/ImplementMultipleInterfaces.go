package myInterface

import "fmt"

//1.实现多个接口
//2.Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。

func ImplMoreInterfaces() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator2 = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	//实现嵌套接口
	var employ EmployeeOperations = e
	employ.DisplaySalary()
	fmt.Println("\nLeaves left =", employ.CalculateLeavesLeft())

	//对于值为 nil 的接口，由于没有底层值和具体类型，当我们试图调用它的方法时，程序会产生 panic 异常。
	var d1 Describer
	d1.Describe()
}

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator2
	LeaveCalculator
}

//接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
//对于值为 nil 的接口，由于没有底层值和具体类型，当我们试图调用它的方法时，程序会产生 panic 异常。
type Describer interface {
	Describe()
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}
