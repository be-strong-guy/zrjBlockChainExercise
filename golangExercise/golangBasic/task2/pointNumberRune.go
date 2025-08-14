package task2

/*
*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func PointNumberRune1(pointNumberRune []int) {
	for i := range pointNumberRune {
		pointNumberRune[i] = pointNumberRune[i] * 2
	}
}

func PointNumberRune2(pointNumberRune *[]int) {
	for i := range *pointNumberRune {
		(*pointNumberRune)[i] *= 2
	}
}
