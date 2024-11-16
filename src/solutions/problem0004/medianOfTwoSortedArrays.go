package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged, i, j, len1, len2 := []int{}, 0, 0, len(nums1), len(nums2)

	// * 合并两个数组，并排序
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	// * 将剩余的数组元素添加到合并后的数组中
	for i < len1 {
		merged = append(merged, nums1[i])
		i++
	}
	for j < len2 {
		merged = append(merged, nums2[j])
		j++
	}

	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		// * 如果合并后的数组长度为偶数，则中位数为中间两个数的平均值
		return float64(merged[mid-1]+merged[mid]) / 2
	} else {
		// * 如果合并后的数组长度为奇数，则中位数为中间的数
		return float64(merged[mid])
	}
}
