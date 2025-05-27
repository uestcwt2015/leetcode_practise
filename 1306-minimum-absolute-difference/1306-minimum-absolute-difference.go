func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    mind := arr[len(arr)-1] - arr[0]
    ans := [][]int{}
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i-1] < mind {
            mind = arr[i] - arr[i-1]
            ans = [][]int{[]int{arr[i-1], arr[i]}}
        } else if arr[i] - arr[i-1] == mind {
            ans = append(ans, []int{arr[i-1], arr[i]})
        }
    }
    return ans
}