package main

import (
	"fmt"
	"strconv"
	"zrjBlockChainExercise/golangExercise/golangBasic/task1"
)

func main1() {
	// 寻找只出现一次的数字
	numberStr := task1.FindOnceNumber()
	if numberStr == "" {
		fmt.Println("没有找到出现一次的数字！")
		return
	}
	number, _ := strconv.Atoi(numberStr)
	fmt.Println("只出现一次的数字是：", number)
}

func main() {
	// 验证字符串是否是有效字符串
	s1 := "([{}])"
	s2 := "{[}]"
	s3 := "[{(}]"
	if task1.IsValidCharacter(s1) {
		fmt.Println("字符串s1是有效字符串！")
	} else {
		fmt.Println("字符串s1不是有效字符串！")
	}
	if task1.IsValidCharacter(s2) {
		fmt.Println("字符串s2是有效字符串！")
	} else {
		fmt.Println("字符串s2不是有效字符串！")
	}
	if task1.IsValidCharacter(s3) {
		fmt.Println("字符串s3是有效字符串！")
	} else {
		fmt.Println("字符串s3不是有效字符串！")
	}
}
