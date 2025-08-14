package task2

import (
	"fmt"
	"sync"
)

/*
*
编写一个程序，使用 go 关键字启动两个协程，
一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/
func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时调用 Done
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
	}
}

func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
	}
}
