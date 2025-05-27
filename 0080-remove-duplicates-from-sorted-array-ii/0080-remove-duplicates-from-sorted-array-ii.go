func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }
    i, j := 2, 2

    for j < len(nums) {
        if nums[j] != nums[i-2] {
            nums[i] = nums[j]
            i++
        }
        j++
    }
    return i
}