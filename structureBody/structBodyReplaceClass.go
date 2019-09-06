package structureBody

import "fmt"

//结构体取代类

func StructBodyReplaceClass() {
	e := Employee3{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()

	//var e2 Employee3
	//e2.LeavesRemaining()

	e3 := New("Sam", "Adolf", 30, 20)
	e3.LeavesRemaining()
}

type Employee3 struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee3) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

//成功地避免了创建不可用的 employee 结构体变量
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
