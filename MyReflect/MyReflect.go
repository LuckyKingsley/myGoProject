package MyReflect

import (
	"fmt"
	"reflect"
)

func MyReflectFunc() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery2(o)

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

type order struct {
	ordId      int
	customerId int
}

type employeeSelf struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	k := t.Kind()
	// Type 表示 interface{} 的实际类型，而 Kind 表示该类型的特定类别。
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind", k)
}

func createQuery2(o interface{}) {
	// NumField() 方法返回结构体中字段的数量，而 Field(i int) 方法返回字段 i 的 reflect.Value。
	if reflect.ValueOf(o).Kind() == reflect.Struct {
		v := reflect.ValueOf(o)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d / type:%T / value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}
