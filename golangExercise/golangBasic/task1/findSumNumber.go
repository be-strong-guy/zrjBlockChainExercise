package task1

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/

func FindSumNumber(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	// 哈希表：key = 数字，value = 下标
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if _, ok := indexMap[complement]; ok {
			return []int{complement, num} // 找到两数
		}
		indexMap[num] = i
	}

	return nil // 没有解
}
