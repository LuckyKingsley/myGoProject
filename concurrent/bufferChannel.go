package concurrent

import (
	"fmt"
	"sync"
	"time"
)

//我们还可以创建一个有缓冲（Buffer）的信道。
//只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
//同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。

func BcFunc() {
	ch := make(chan string, 2)
	ch <- "abc"
	ch <- "def"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}

	//缓冲信道的容量是指信道可以存储的值的数量。我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
	//缓冲信道的长度是指信道中当前排队的元素个数。
	ch3 := make(chan string, 3)
	ch3 <- "naveen"
	ch3 <- "paul"
	fmt.Println("capacity is", cap(ch3))
	fmt.Println("length is", len(ch3))
	fmt.Println("read value", <-ch3)
	fmt.Println("read value", <-ch3)
	fmt.Println("new length is", len(ch3))

	//而 WaitGroup 用于实现工作池，因此要理解工作池，我们首先需要学习 WaitGroup。
	//WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
	//假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。Go 主协程需要等待这 3 个协程执行结束后，才会终止。
	//这就可以用 WaitGroup 来实现。
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		//传递 wg 的地址是很重要的。如果没有传递 wg 的地址，那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，
		// 因而当它们执行结束时，main 函数并不会知道。
		go process(i, &wg)
	}
	wg.Wait()

}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
