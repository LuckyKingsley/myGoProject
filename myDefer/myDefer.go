package myDefer

import (
	"fmt"
	"sync"
)

// 含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数
// 不仅限于函数的调用，调用方法也是合法的

func DeferFunc() {
	//nums := []int{78, 109, 2, 563, 300}
	//largest(nums)

	// 并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值
	//a := 5
	//defer printA(a)
	// 当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，
	// 随后按照后进先出（Last In First Out, LIFO）的顺序执行
	//defer printA(1)
	//a = 10

	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}
