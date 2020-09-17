package main

import (
	"fmt"
	//"strings"
	"time"
)

// 使用NewTicker定时打印
func PriTicker() {
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("ticker")
	}
}

// 使用NewTimer定时打印
func PriTimer() {
	for {
		t := time.NewTimer(time.Second * 3)
		<-t.C
		fmt.Println("timer")
	}
}

//只打印一次
func PriTimer1() {
	t := time.NewTimer(time.Second * 3)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("timer")
	}
}

func main() {
	//timeStr := time.Now().Format("2006-01-02 15:04:05")
	//day := strings.Split(timeStr, " ")[0]
	//fmt.Println(day)

	go PriTicker()
	go PriTimer1()

	time.Sleep(time.Minute * 1)
}

// 区别
