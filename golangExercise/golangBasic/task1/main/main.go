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

func main2() {
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

// 测试查找最长公共字符串前缀
func main3() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"aa", "bb", "cc"}
	s := task1.LongestCommonPrefix(strs1)
	if s != "" {
		fmt.Println("字符串数组strs1最长公共前缀是:", s)
	} else {
		fmt.Println("字符串数组strs1不存在公共前缀")
	}
	s = task1.LongestCommonPrefix(strs2)
	if s != "" {
		fmt.Println("字符串数组strs2最长公共前缀是:", s)
	} else {
		fmt.Println("字符串数组strs2不存在公共前缀")
	}
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func main4() {
	var array1 []int = []int{1, 5, 6, 3}
	var array2 []int = []int{9, 9, 9}
	fmt.Println("array1加一结果是", task1.PlusOne(array1))
	fmt.Println("array2加一结果是", task1.PlusOne(array2))
}

// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次
func main5() {
	var array []int = []int{1, 1, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7}
	fmt.Println("新数组长度是:", task1.RemoveNumber(array))
}

// 合并区间
func main6() {
	var array [][]int = [][]int{
		{1, 3},
		{4, 6},
		{5, 8},
		{7, 10},
		{20, 30},
	}
	fmt.Println("合并后的数组：", task1.MergeArray(array))
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func main() {
	var array []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var target int = 9
	fmt.Println("找到的结果是", task1.FindSumNumber(array, target))
}
