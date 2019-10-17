package myMap

import "fmt"

func MyMap() {
	// map 的零值是 nil。
	// 如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
	//personSalary := make(map[stringprac]int)
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("map is nil. Going to make one.")
		map1 = make(map[string]int)
	}
	map1["a"] = 123123
	map1["b"] = 123123312
	map1["c"] = 151235

	map2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	map2["mike"] = 9000
	fmt.Println("map contents:", map2)
	fmt.Println("map contents:", map2["mike"])
	//判断是否存在key
	value, ok := map2["mike"]
	if ok == true {
		fmt.Println("mike", "is", value)
	} else {
		fmt.Println("mike not found")
	}
	for key, value := range map2 {
		fmt.Printf("map2[%s] = %d\n", key, value)
	}
	delete(map2, "mike")
	fmt.Println(len(map2))

	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

	//map 之间不能使用 == 操作符判断，
	// == 只能用来检查 map 是否为 nil。

}
