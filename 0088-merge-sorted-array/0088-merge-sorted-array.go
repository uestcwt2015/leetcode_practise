func merge(nums1 []int, m int, nums2 []int, n int)  {
    result := []int{}
    i, j := 0, 0
    for i < m && j < n {
        if nums1[i] <= nums2[j] {
            result = append(result, nums1[i])
            i++
        } else {
            result = append(result, nums2[j])
            j++
        }
    }

    if i < m {
        result = append(result, nums1[i:]...)
    }
 
    if j < n {
        result = append(result, nums2[j:]...)
    }

    copy(nums1, result)
}