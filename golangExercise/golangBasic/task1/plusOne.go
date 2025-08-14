package task1

/*
*给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
 */
func PlusOne(array []int) []int {

	n := len(array)

	for i := n - 1; i >= 0; i-- {
		array[i]++
		if array[i] < 10 {
			return array // 没有进位，直接返回
		}
		array[i] = 0 // 进位
	}

	// 如果最高位有进位，比如 999 变成 1000
	newArray := make([]int, n+1)
	newArray[0] = 1
	return newArray

}
