func minCount(coins []int) int {
    res := 0
    for i:=0; i<len(coins);i++{
        times := Use(coins[i])
        res += times
    }
    return res
}

func Use(nums int) int {
    count := 0
    for nums > 0 {
        if nums - 2 >= 0 {
            count ++
            nums = nums - 2
        } else {
            count ++
            break
        }
    }
    return count
}
