package main

import "fmt"

func main() {

}

func deferTest() error {
	defer fmt.Println("此段代码延迟执行")
	fmt.Println("正常执行")
	return fmt.Errorf("错误返回")
}

func freakOut() {
	defer calmDown()
	panic("程序奔溃")
}

func calmDown() {
	recover()
}
