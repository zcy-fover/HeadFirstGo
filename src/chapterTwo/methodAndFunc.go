package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var now = time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.String())

	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(nil)
	}
	fmt.Println(input)
}
