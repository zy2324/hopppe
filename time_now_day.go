package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	str := t.Format("2006-01-02")
	fmt.Println(str)
}
