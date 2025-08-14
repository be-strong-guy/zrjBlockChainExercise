package task1

import "strconv"

/*
*
给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func FindOnceNumber() string {
	var arry []int = []int{1, 1, 2, 2, 3, 3, 4, 4, 8, 6, 6}
	var numMap = make(map[int]int)
	var onceNumber string
	for i := 0; i < len(arry); i++ {
		if _, ok := numMap[arry[i]]; ok {
			numMap[arry[i]] += numMap[arry[i]]
		} else {
			numMap[arry[i]] = 1
		}
	}
	for key, value := range numMap {
		if value == 1 {
			onceNumber = strconv.Itoa(key)
		}
	}
	return onceNumber
}
