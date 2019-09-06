package concurrent

import (
	"fmt"
	"sync"
)

func MutexFunc() {
	//当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码称为临界区
	//如果在任意时刻只允许一个 Go 协程访问临界区，那么就可以避免竞态条件。而使用 Mutex 可以达到这个目的

	//用mutex处理竞争临界值
	//var w sync.WaitGroup
	//var m sync.Mutex
	//for i := 0; i < 100; i++ {
	//	w.Add(1)
	//	go increment1(&w, &m)
	//}
	//w.Wait()
	//fmt.Println("final value of x is", x)

	//用信道处理竞争临界值
	y := 0
	var wg sync.WaitGroup
	var cn chan int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment2(y, cn, &wg)
		y = <-cn
	}
	wg.Wait()
	fmt.Println("final value of x is", x)

}

var x = 0

func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func increment2(y int, cn chan int, wg *sync.WaitGroup) {
	y = y + 1
	cn <- y
	wg.Done()
}
