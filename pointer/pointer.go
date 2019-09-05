package pointer

import "fmt"

func PointerFunc() {
	b := 233
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	//& 操作符用于获取变量的地址

	c := 231
	var d *int
	if d == nil {
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	//解引用
	e := 255
	f := &e
	fmt.Println("address of e is", f)
	fmt.Println("value of e is", *f)

	//指针修改值
	h := 542
	i := &h
	fmt.Println("address of h is", i)
	fmt.Println("value of h is", *i)
	*i++
	fmt.Println("new value of h is", h)
	fmt.Println("address of h is", i)
	fmt.Println("value of h is", *i)

	//向函数传递指针参数
	//不要向函数传递数组的指针，而应该使用切片
	k := [3]int{12, 452, 52}
	modifyArray(k[:])
	fmt.Println(k)
}

func modifyArray(array []int) {
	array[0] = 424
}
