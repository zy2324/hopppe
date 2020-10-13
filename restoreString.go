func restoreString(s string, indices []int) string {
    var res string

    for i:=0; i < len(indices); i++ {
        res += string(s[indices[i]])
    }
    return res
}
