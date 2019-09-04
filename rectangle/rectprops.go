package rectangle

import (
	"math"
	"fmt"
)

//包的初始化顺序如下：
//首先初始化包级别（Package Level）的变量
//紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用。

func init(){
	//执行初始化任务，也可用于在开始执行之前验证程序的正确性
	fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}