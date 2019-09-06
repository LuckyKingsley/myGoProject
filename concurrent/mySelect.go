package concurrent

import (
	"fmt"
	"time"
)

//select 语句用于在多个发送/接收信道操作中进行选择。
//select 语句会一直阻塞，直到发送/接收操作准备就绪。
// 如果有多个信道操作准备完毕，select 会随机地选取其中之一执行

//选择首先响应的服务器，而忽略其它的响应。
//使用这种方法，我们可以向多个服务器发送请求，并给用户返回最快的响应

func MySeclect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	flag := false
	for {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			flag = true
		case s2 := <-output2:
			fmt.Println(s2)
			flag = true
		default:
			fmt.Println("no value received")
		}
		if flag == true {
			break
		}
	}

}

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server2"
}
