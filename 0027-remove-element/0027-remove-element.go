func removeElement(nums []int, val int) int {
    i, j := 0, 0
    for j < len(nums) {
        if nums[j] != val {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
        j++
    }
    return i
}