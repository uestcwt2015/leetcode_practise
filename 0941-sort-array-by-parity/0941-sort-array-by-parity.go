func sortArrayByParity(nums []int) []int {
    i, j := 0, len(nums)-1
    for i < j {
        if nums[i] % 2 == 0 && nums[j] % 2 != 0 {
            i++
            j--
        } else if nums[i] % 2 != 0 && nums[j] % 2 == 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        } else if nums[i] % 2 == 0 {
            i++
        } else if nums[j] % 2 != 0 {
            j--
        }
    }
    return nums
}