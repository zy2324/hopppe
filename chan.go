package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func main() {
	ch := make(chan string, 3)
	go pro(ch)

	go cus(ch)

	time.Sleep(time.Second * 1)
}
