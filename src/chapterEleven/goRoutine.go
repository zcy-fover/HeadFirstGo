package main

import (
	"fmt"
)

func main() {
	//go playA()
	//go playB()
	//time.Sleep(1 * time.Second)
	//fmt.Println("执行完毕")

	myChan := make(chan int)
	go addNum(myChan)
	fmt.Println(<-myChan)
}

func playA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
}

func playB() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}

func addNum(myChan chan int) {
	myChan <- 5
}
