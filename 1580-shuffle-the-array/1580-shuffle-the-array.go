func shuffle(nums []int, n int) []int {
    result := make([]int, n * 2)
    i, j := 0, n
    for k := 0; k < 2 * n; k+=2 {
        result[k] = nums[i]
        result[k+1] = nums[j]
        i++
        j++
    }

    return result
}