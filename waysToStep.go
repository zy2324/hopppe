package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func waysToStep(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	dp1, dp2, dp3 := 1, 2, 4
	for i := 4; i <= n; i++ {
		dp1, dp2, dp3 = dp2, dp3, (dp1+dp2+dp3)%1000000007
	}
	return dp3
}

// 为什么是100000007，最大的十位质数．可以防止结果溢出