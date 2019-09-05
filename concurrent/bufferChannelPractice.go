package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//我们工作池的核心功能如下：
//创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
//将作业添加到该输入型缓冲信道中。
//作业完成后，再将结果写入一个输出型缓冲信道。
//从输出型缓冲信道读取并打印结果。

func BcpFunc() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits2(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits2(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
