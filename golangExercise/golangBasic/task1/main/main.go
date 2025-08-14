package main

import (
	"fmt"
	"strconv"
	"zrjBlockChainExercise/golangExercise/golangBasic/task1"
)

func main() {
	// 寻找只出现一次的数字
	numberStr := task1.FindOnceNumber()
	if numberStr == "" {
		fmt.Println("没有找到出现一次的数字！")
		return
	}
	number, _ := strconv.Atoi(numberStr)
	fmt.Println("只出现一次的数字是：", number)
}
