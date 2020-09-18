//一个生产者．多个消费者
package main

import (
	"time"
	"fmt"
)

func Pro(ch chan int) {
	for i:=1; i <= 10; i++ {
		ch <- i
	}
}

func Con(ch chan int) {
	res := <- ch
	fmt.Println(res)
}



func main() {
	ch := make(chan int)

	go Pro(ch)

	for i:=1; i <= 5; i++ {
		go Con(ch)
	}

	time.Sleep(5*time.Second)

}
// 运行结果：较符合预期
// 如何不使用time