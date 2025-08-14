package task1

/*
*
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1)
额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，
并将 i 后移一位。
*/
func RemoveNumber(array []int) int {
	if len(array) == 0 {
		return 0
	}
	var newLen = 0
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			array[newLen] = array[i]
			newLen++
		}
	}
	return newLen
}
