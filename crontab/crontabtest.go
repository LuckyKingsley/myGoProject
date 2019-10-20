package crontab

import (
	"fmt"
	"github.com/robfig/cron"
)

func CrontabTest() {
	crontabTestInner()
}

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

func crontabTestInner() {
	// 1

	//i := 0
	//c := cron.New()
	//spec := "*/5 * * * ?"
	//entryId, err := c.AddFunc(spec, func() {
	//	i++
	//	fmt.Println("cron running:", i)
	//})
	//fmt.Println(entryId, err)
	//c.Start()
	//
	//select{}

	// 2

	i := 0
	c := cron.New()

	//AddFunc
	spec := "*/5 * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
