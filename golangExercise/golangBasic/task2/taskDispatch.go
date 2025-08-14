package task2

import (
	"fmt"
	"sync"
	"time"
)

/**
设计一个任务调度器，接收一组任务（可以用函数表示），
并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/
//定义任务
type Task func()

// 任务调度器
func TaskDispatch(task []Task) {
	var wg sync.WaitGroup
	wg.Add(len(task))
	for i, t := range task {
		go func(idx int, task Task) {
			defer wg.Done()

			start := time.Now()
			fmt.Printf("第 %d 个任务开始执行....\n", i)
			t()
			spendTime := time.Since(start)
			fmt.Printf("第 %d 个任务执行时间是：%v\n", i, spendTime)
		}(i, t)
	}
	wg.Wait()
}
