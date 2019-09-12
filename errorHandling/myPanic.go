package errorHandling

import (
	"fmt"
	"runtime/debug"
	"time"
)

// 你应该尽可能地使用错误，而不是使用 panic 和 recover。
// 只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制。

// panic 有两个合理的用例。
// 发生了一个不能恢复的错误，此时程序不能继续运行。 一个例子就是 web 服务器无法绑定所要求的端口。在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。
// 发生了一个编程上的错误。 假如我们有一个接收指针参数的方法，而其他人使用 nil 作为参数调用了它。在这种情况下，我们可以使用 panic，因为这是一个编程错误：用 nil 参数调用了一个只能接收合法指针的方法。

// 只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
// recover 是一个内建函数，用于重新获得 panic 协程的控制。
func MyPanicFunc() {
	//defer fmt.Println("deferred call in main")
	//firstName := "Elon"
	//fullName(&firstName, nil)
	//fmt.Println("returned normally from main")

	// 只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic
	// 当我们恢复 panic 时，我们就释放了它的堆栈跟踪。实际上，在上述程序里，恢复 panic 之后，我们就失去了堆栈跟踪
	//有办法可以打印出堆栈跟踪，使用 Debug 包中的 PrintStack 函数。
	a()
	fmt.Println("normally returned from main")
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic("runtime error: firstName cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
		debug.PrintStack()
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}
