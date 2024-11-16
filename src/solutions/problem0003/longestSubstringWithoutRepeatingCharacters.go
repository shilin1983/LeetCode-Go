package problem0003

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	left, right, maxLength, hashTable := 0, 0, 0, make(map[byte]int, length)

	for right < length {
		char := s[right]

		// * 如果哈希表中存在当前元素，则移动滑动窗口左边界到当前元素的下一个位置
		if idx, ok := hashTable[char]; ok {
			left = max(left, idx+1)
		}

		// * 将当前元素及其索引插入哈希表
		hashTable[char] = right
		// * 移动滑动窗口右边界
		right++
		// * 计算当前滑动窗口最大长度
		maxLength = max(maxLength, right-left)
	}

	return maxLength
}
