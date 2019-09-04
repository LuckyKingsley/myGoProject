package slice

import "fmt"

func ArraySlice(){
	var a [3]int
	fmt.Println(a)
	b := [3]int{12, 45, 46}
	fmt.Println(b)
	c := [3]int {12}
	fmt.Println(c)
	d := [...]int{12, 24, 31}
	fmt.Println(d)

	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e // e copy of a is assigned to f
	f[0] = "Singapore"
	fmt.Println("e is ", e)
	fmt.Println("f is ", f)

	for i, v := range e {
		fmt.Printf("%d the element of a is %s\n", i, v)
	}
	for _, v := range e {
		fmt.Printf("the element of a is %s\n", v)
	}

	// 多维数组
	g := [3][2] string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	fmt.Println(g)
	var h [3][2]string
	h[0][0] = "apple"
	h[0][1] = "samsung"
	h[1][0] = "microsoft"
	h[1][1] = "google"
	h[2][0] = "AT&T"
	h[2][1] = "T-Mobile"

	// 切片
	j := [3]int{123, 321, 32}
	var k = j[0:2]
	fmt.Println(k)
	l := []int{6, 7, 8}
	fmt.Println(l)
	//切片做的操作会修改原有数组的值
	m := make([]int, 5, 5)
	fmt.Println(m)
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6

	// 切片类型的零值为 nil
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:",names)
	}

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)

	//多维切片
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	//切片在内存中，数组就不能被垃圾回收
	//假设我们有一个非常大的数组，我们只想处理它的一小部分。
	// 然后，我们由这个数组创建一个切片，并开始处理切片。
	// 这里需要重点注意的是，在切片引用时数组仍然存在内存中。
	// 一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。
	// 这样我们可以使用新的切片，原始数组可以被垃圾回收。
	countries := [...]string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
