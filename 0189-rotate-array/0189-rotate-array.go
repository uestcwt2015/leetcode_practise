func rotate(nums []int, k int)  {
    n := len(nums)
    step := k % n
    copy(nums, append(nums[n-step:], nums[:n-step]...))
}