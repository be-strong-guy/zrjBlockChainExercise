package main

import (
	"fmt"
	"sync"
	"zrjBlockChainExercise/golangExercise/golangBasic/task2"
)

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func main1() {
	var a int = 9
	task2.PointNumber(&a)
	fmt.Println(a)
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func main2() {
	var rune1 = []int{1, 2, 3}
	task2.PointNumberRune1(rune1)
	fmt.Println(rune1)
	var rune2 = []int{4, 5, 6}
	task2.PointNumberRune2(&rune2)
	fmt.Println(rune2)
}

// 编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func main3() {
	var wg sync.WaitGroup
	wg.Add(2)

	go task2.PrintOdd(&wg)
	go task2.PrintEven(&wg)
	wg.Wait()
	fmt.Println("打印完成")
}

func main() {
	var tasks []task2.Task = []task2.Task{
		func() {
			fmt.Println("我是任务1")
		},
		func() {
			fmt.Println("我是任务2")
		},
		func() {
			fmt.Println("我是任务3")
		},
		func() {
			fmt.Println("我是任务4")
		},
	}
	task2.TaskDispatch(tasks)
}
