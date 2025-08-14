package task2

import (
	"fmt"
	"sync"
)

/*
*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func SyncMutexTest() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	var number int = 0
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				number++
				mutex.Unlock()
			}
		}()

	}
	wg.Wait()
	fmt.Println("10个协程累计1000的结果是:", number)
}
