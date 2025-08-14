package task2

import (
	"fmt"
	"sync"
)

/*
*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
消费者协程从通道中接收这些整数并打印。
*/
func ChannelBuffer() {
	var wg sync.WaitGroup
	wg.Add(2)
	var channel chan int = make(chan int, 5)
	go sendNumberBuffer(channel, &wg)
	go receiveNumberBuffer(channel, &wg)
	wg.Wait()

}
func sendNumberBuffer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Println("发数据：", i)
		channel <- i
	}
	close(channel)

}

func receiveNumberBuffer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range channel {
		fmt.Println("收到的数据是:", num)
	}
}
