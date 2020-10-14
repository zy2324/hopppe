package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		now := time.Now() //获取当前时间，放到now里面，要给next用
		fmt.Println(now)
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                    //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		fmt.Printf("凌晨执行一次: %v\n", time.Now())
	}
}
