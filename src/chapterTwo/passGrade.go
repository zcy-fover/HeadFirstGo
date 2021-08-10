package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for i := 10; i > 0; i-- {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("输入异常")
			log.Fatal(err)
		}
		guessNum, err :=  strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("数字输入异常")
			log.Fatal(err)
		}
		if guessNum > 60 {
			 fmt.Println("合格：",  guessNum)
		} else {
			fmt.Println("不合格：",  guessNum)
		}
	}
}
