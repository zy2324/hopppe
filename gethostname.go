package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	cmd := "hostname"
	out, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return
	}
	res := strings.Split(string(out), "\n")[0]
	fmt.Println(res)
}
