package task2

import (
	"fmt"
	"sync"
)

/**
编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。
*/

func Channel() {
	var wg sync.WaitGroup
	wg.Add(2)
	var channel chan int = make(chan int)
	go sendNumber(channel, &wg)
	go receiveNumber(channel, &wg)
	wg.Wait()

}

func sendNumber(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("发数据：", i)
		channel <- i
	}
	close(channel)

}

func receiveNumber(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range channel {
		fmt.Println("收到的数据是:", num)
	}
}
