func merge(nums1 []int, m int, nums2 []int, n int)  {
    x, y := 0, 0
	l := m + n
	temp := make([]int, l)

	for i := 0; i < l; i++ {
		if x >= m {
			temp[i] = nums2[y]
			y++
			continue
		}

		if y >= n {
			temp[i] = nums1[x]
			x++
			continue
		}

		if nums1[x] < nums2[y] {
			temp[i] = nums1[x]
			x++
			continue
		}

		if nums1[x] == nums2[y] {
			temp[i] = nums1[x]
			temp[i+1] = nums2[y]
			x++
			y++
			i++
			continue
		}

		if nums1[x] > nums2[y] {
			temp[i] = nums2[y]
			y++
			continue
		}
	}

	copy(nums1, temp)
}