func applyOperations(nums []int) []int {
    i, j := 0, 0
    for j < len(nums) {
        if j < len(nums) - 1 && nums[j] == nums[j+1] {
            nums[j] *= 2
            nums[j+1] = 0
        }

        if nums[j] != 0 {
            nums[i] = nums[j]
            i++
        }
        j++
    }
    for i < len(nums) {
        nums[i] = 0
        i++
    }
    return nums
}