func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    ans := []int{}
    for k, v := range nums {
        if v == target {
            ans = append(ans, k)
        }
    }
    return ans
}