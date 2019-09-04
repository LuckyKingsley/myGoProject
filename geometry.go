package main

import (
	"log"
	_ "math"
	"geometry/string"
)

//在Go中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

//var _ = math.Min // 错误屏蔽器

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	//fmt.Println("Geometrical shape properties")
	///*Area function of rectangle package used*/
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	///*Diagonal function of rectangle package used*/
	//fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
	//
	//if num := 10; num%2 == 0 { //checks if number is even
	//	fmt.Println(num, "is even")
	//} else {
	//	fmt.Println(num, "is odd")
	//}

	//base.Base()
	//slice.ArraySlice()
	//myMap.MyMap()
	string.StringFunc()
}
