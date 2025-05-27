func sortedSquares(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    i, j, k := 0, n-1, n-1
    for j >= i {
        p, q := nums[i] * nums[i], nums[j] * nums[j]
        if p < q {
            result[k] = q
            j--
        } else {
            result[k] = p
            i++
        }
        k--
    }
    return result
}