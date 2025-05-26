func sortColors(nums []int)  {
    i, j, k := 0, -1, len(nums)
    for i < k {
        if nums[i] == 0 {
            j++
            if j != i {
                nums[i], nums[j] = nums[j], nums[i]
            }
            i++
        } else if nums[i] == 2 {
            k--
            if i != k {
                nums[i], nums[k] = nums[k], nums[i]
            }
        } else {
            i++
        }
    }
}

// func sortColors(nums []int)  {
//     x := []int{0, 0, 0}
//     for _, v := range nums {
//         x[v]++
//     }

//     k := 0
//     for i := range []int{0, 1, 2} {
//         for j := 0; j < x[i]; j++ {
//             nums[k] = i
//             k++
//         }
//     }
// }