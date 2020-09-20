func intervalIntersection(A [][]int, B [][]int) [][]int {
	var res [][]int
	for _, a := range A {
		for _, b := range B {
		//	tmp := []int{}
			if a[1] > b[0] && b[1] > b [0] {
		//		tmp = append(tmp, b[0], a[1])
          //      res = append(res, tmp)
          		res = append(res, []int{int(math.Max(float64(a))), int(math.Max(float64(b)))})
			}
		}
	}
    return res
}

//转自go题解
func intervalIntersection(A [][]int, B [][]int) [][]int {
	curA, curB, lenA, lenB := 0, 0, len(A), len(B)
	var result [][]int
	for curA < lenA && curB < lenB {
		leftA, rightA := A[curA][0], A[curA][1]
		leftB, rightB := B[curB][0], B[curB][1]
		// 如果有A/B有交集
		if rightB >= leftA && rightA >= leftB {
			// 计算出交集，加入 res
			result = append(result,
				[]int{int(math.Max(float64(leftA), float64(leftB))),
					int(math.Min(float64(rightA), float64(rightB)))})
		}
		if rightB < rightA {
			curB++
		} else {
			curA++
		}
	}
	return result
}

