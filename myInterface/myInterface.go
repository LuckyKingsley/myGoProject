package myInterface

import "fmt"

//在 Go 语言中，接口就是方法签名（Method Signature）的集合。
// 当一个类型定义了接口中的所有方法，我们称它实现了该接口。
// 这与面向对象编程（OOP）的说法很类似。接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
func InterFunc() {
	//1
	//name := MyString("Sam Anderson")
	//var v VowelsFinder
	//v = name
	//fmt.Printf("Vowels are %c", v.FindVowels())
	//2
	name := MyString("Sam Anderson")
	//var v VowelsFinder
	//v = name
	fmt.Printf("Vowels are %c", name.FindVowels())
	//3
	SalaryMain()
	//4
	//接口的内部表示:Interface type myInterface.MyString value auth
	var v VowelsFinder
	data := MyString("auth")
	v = data
	describe(v)
	v.FindVowels()

	//没有包含方法的接口称为空接口。空接口表示为 interface{}。
	//由于空接口没有方法，因此所有类型都实现了空接口。

	//类型断言
	var s interface{} = "Steven Paul"
	assert(s)
	var t interface{} = 56
	assert(t)

	//类型选择
	findType("Naveen")
	findType(77)
	findType(89.98)
	findType(MyString("dasdasd"))

}

type MyString string
type VowelsFinder interface {
	FindVowels() []rune
}

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func describe(v VowelsFinder) {
	fmt.Printf("Interface type %T value %v\n", v, v)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("I am a stringprac and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	case VowelsFinder:
		v.FindVowels()
	default:
		fmt.Printf("Unknown type\n")
	}
}
