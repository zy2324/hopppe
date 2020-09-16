//使用go chan，处理大文件
package main


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

	defer file.Close()  // 文件执行结束，关闭文件

	read := bufio.NewReader(file)
	for {
		line, isPrefix, err := read.Readline() // isPrefix:line如果太长超过了,返回false
		checkError(err)

		go func() {
			res, err := AimDataV1(line)
			ch <- res
		}

		ch <- line
		go AimDataV2(ch)

	}

	for c := range ch {
		fmt.Println(c)
	}

}


func AimDataV1(word string) (string, error) {
	tmp, err := GetData(word)
	checkError(err)

	fmt.Println(tmp)
}


func AimDataV2(word chan string) {
	res := <- word

	tmp, err := GetData(res)
	checkError(err)

	fmt.Println(tmp)
}


func GetData(data string) ([]string, error) {
	return strings.Split(data, ",")
}