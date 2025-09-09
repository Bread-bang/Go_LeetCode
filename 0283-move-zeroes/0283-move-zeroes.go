func moveZeroes(nums []int)  {
    arrLength := len(nums)
    if arrLength <= 1 {
        fmt.Printf("%v", nums)
        return
    }

    for i := 0; i < arrLength - 1; i++ {
        for j := 0; j < arrLength - 1; j++ {
            if nums[j] == 0 {
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
            }
        }
    }

    fmt.Println(nums)
}