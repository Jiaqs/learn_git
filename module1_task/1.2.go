/*
基于Channel编写一个简单的单线程生产者消费者模型：
队列：
队列长度为10，队列元素为int
生产者：
每一秒往队列中放入一个int的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
package main

import (
	"fmt"
	"time"
)

func producer(queue chan<- int) {
	for i := 0; ; i++ {
		queue <- i              // 将元素放入队列
		time.Sleep(time.Second) // 等待 1 秒
	}
}

func consumer(queue <-chan int) {
	for {
		item := <-queue // 从队列中获取元素
		fmt.Println(item)
		time.Sleep(time.Second) // 等待 1 秒
	}
}

func main() {
	queue := make(chan int, 10) // 创建带有长度为 10 的缓冲队列

	go producer(queue) // 启动生产者协程
	consumer(queue)    // 执行消费者函数（主线程中执行）
}
