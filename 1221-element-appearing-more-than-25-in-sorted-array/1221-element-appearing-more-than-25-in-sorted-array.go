func findSpecialInteger(arr []int) int {
    num, cnt, n := arr[0], 0, len(arr)
    for _, v := range arr {
        if v == num {
            cnt++
            if cnt * 4 > n {
                return num
            }
        } else {
            num = v
            cnt = 1
        }
    }
    return 0
}