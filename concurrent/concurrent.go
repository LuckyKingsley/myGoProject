package concurrent

import (
	"fmt"
	"time"
)

func CcFunc() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")

	//go numbers()
	//go alphabets()
	//time.Sleep(3000 * time.Millisecond)

	//所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
	//a := make(chan int)
	//当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
	// 与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	//data := <- a //读取信道a
	//a <- data // 写入信道a

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	sqrches, cubeches := <-sqrch, <-cubech
	fmt.Println("Final output", sqrches+cubeches)

	//单向信道
	//sendch := make(chan<- int)
	//go sendData(sendch)
	//fmt.Println(<-sendch)

	//信道转换
	//把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)

	//数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
	//当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
	ch := make(chan int)
	go producer(ch)
	//for {
	//	v, ok := <- ch
	//	if ok == false {
	//		break
	//	} else {
	//		fmt.Println("Received ", v, ok)
	//	}
	//}
	for v := range ch {
		fmt.Println("Received ", v)
	}

}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func digits(number int, digChnl chan int) {
	for number != 0 {
		digit := number % 10
		number /= 10
		digChnl <- digit
	}
	close(digChnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	digChnl := make(chan int)
	go digits(number, digChnl)
	for digit := range digChnl {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	digchnl := make(chan int)
	go digits(number, digchnl)
	for digit := range digchnl {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 5; i++ {
		chnl <- i
	}
	close(chnl)
}
