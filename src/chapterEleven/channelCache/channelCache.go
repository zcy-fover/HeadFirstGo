package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 6; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main 执行结束")
}
