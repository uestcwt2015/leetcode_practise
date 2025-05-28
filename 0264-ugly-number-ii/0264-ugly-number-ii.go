func nthUglyNumber(n int) int {
    ans := make([]int, n)
    ans[0] = 1
    p1, p2, p3 := 0, 0, 0
    for i := 1; i < n; i++ {
        v1, v2, v3 := ans[p1] * 2, ans[p2] * 3, ans[p3] * 5
        mind := min(v1, min(v2, v3))
        if mind == v1 {
            p1++
        }
        if mind == v2 {
            p2++
        }
        if mind == v3 {
            p3++
        }
        ans[i] = mind
    }
    return ans[n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}