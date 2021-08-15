package main

import "fmt"

func abcPlay(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func defPlay(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func channelPlay() {
	c1 := make(chan string) //创建两个channel使得交替打印 adbecf
	c2 := make(chan string)
	go abcPlay(c1) //分别启动两个 goroutine
	go defPlay(c2)
	fmt.Println(<-c1) //阻塞，等到abcPlay处理完成输出c1发送回来的值
	fmt.Println(<-c2) //这里可以看出来c1和c2在交替阻塞主 goroutine 的执行，否则
	fmt.Println(<-c1) //程序就不会始终按照 adbecf 的顺序来打印了
	fmt.Println(<-c2)
	fmt.Println(<-c1)
	fmt.Println(<-c2)
}
