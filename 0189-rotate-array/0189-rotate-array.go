func rotate(nums []int, k int)  {
    l := len(nums)
    temp := make([]int, l)

    for i := 0; i < l; i++ {
        idx := (i + k) % l;
        temp[idx] = nums[i]
    }

	copy(nums, temp)
}