//使用go chan，处理大文件
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	//"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 按行读文件，放入chan中
func HandleFile(filepath string) (bool, error) {
	ch := make(chan string, 5)

	file, err := os.Open(filepath)
	checkError(err)

	defer file.Close() // 文件执行结束，关闭文件

	read := bufio.NewReader(file)
	for {
		line, _, err := read.ReadLine() // isPrefix:读到行末尾，范围false

		if err == io.EOF {
			//close(ch)
			break
		}

		checkError(err)

		//if !isPrefix {
		//	fmt.Println("line is too long")
		//}

		go func() {
			res, err := AimDataV1(string(line))
			checkError(err)
			ch <- res
			//	select {
			//	case ch <- res:
			//		fmt.Println("ok")
			//	default:
			//		fmt.Println("channel is full")
			//	}
		}()

		//ch <- line
		//go AimDataV2(ch)

	}

	for c := range ch {
		fmt.Println(c)
		if c == "g" {
			break
		}
	}

	return true, nil
}

func AimDataV1(word string) (string, error) {
	tmp, err := GetData(word)
	checkError(err)
	return tmp, err
}

func AimDataV2(word chan string) (string, error) {
	res := <-word

	tmp, err := GetData(res)
	checkError(err)
	return tmp, err
}

func GetData(data string) (string, error) {
	return strings.Split(data, ",")[0], nil
}

func main() {
	//ch := make(chan string, 5)
	path := "test"
	res, err := HandleFile(path)
	checkError(err)

	if res {
		fmt.Println("read handle ok")
	}
}
