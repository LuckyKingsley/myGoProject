package structureBody

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

//匿名字段
type Person struct {
	string
	int
}

type Address struct {
	city, state string
}
type Man struct {
	name    string
	age     int
	address Address
}

type image struct {
	data map[int]int
}

func SBFunc() {
	//111111111111111111111111111111111111
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//之所以称这种结构体是匿名的，
	//是因为它只是创建一个新的结构体变量 em3，而没有定义任何结构体类型。
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	emp4 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Printf("Salary: $%d", emp4.salary)

	var emp5 Employee
	emp5.firstName = "Jack"
	emp5.lastName = "Adams"
	fmt.Println("Employee 5:", emp5)

	//指针
	//可以使用 emp6.firstName 来代替显式的解引用 (*emp6).firstName
	emp6 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp6).firstName)
	fmt.Println("Age:", (*emp6).age)

	//22222222222222222222222222222222222222222
	p := Person{"Naveen", 50}
	fmt.Println(p)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	//333333333333333333333333333333333333
	var spec Spec
	spec.Maker = "apple"
	spec.Price = 50000
	spec.model = "dasd"
	fmt.Println("Spec:", spec)

	//结构体是值类型。
	// 如果它的每一个字段都是可比较的，则该结构体也是可比较的。
	// 如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
	//结构体包含不可比较的字段，则结构体变量也不可比较。
	//image1 := image{data: map[int]int{
	//	0: 155,
	//}}
	//image2 := image{data: map[int]int{
	//	0: 155,
	//}}
	//if image1 == image2 {
	//	fmt.Println("image1 and image2 are equal")
	//}
}
