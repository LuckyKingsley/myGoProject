package errorHandling

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// 错误处理

func ErrorHandingFunc() {

}

func errorHanding1() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// Java当中有instanceof这样的关键字判断类型 Go当中自然也有相应的方法来判断类型
// 即Comma-ok断言
// 写法为value, ok := em.(T)
// 如果确保em 是同类型的时候可以直接使用value:=em.(T)一般用于switch语句中下面将会讲解到
// em代表要判断的变量
// T代表被判断的类型
// value代表返回的值
// ok代表是否为改类型
// em必须为initerface类型才可以进行类型断言
// 当函数作为参数并且被调用函数将参数类型指定为interface{}的时候是没有办法直接调用该方法的

// 断言底层结构体类型，使用结构体字段获取更多信息
func errorHanding2() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// 断言底层结构体类型，调用方法获取更多信息
func errorHanding3() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func errorHanding4() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}
