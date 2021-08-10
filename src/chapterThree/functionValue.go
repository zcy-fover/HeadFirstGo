package main

import (
	stringUtils "HeadFirstGo/src/common"
	"fmt"
)

func main() {
	aa := 5
	paramChange(aa)
	fmt.Println(aa)
	paramAddress(aa)
	paramPoint()

	fmt.Println(*funcPoint(&aa))
	fmt.Println(aa)
	stringUtils.Test()
}

func calculateArea(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("输入异常：width{%.2f}, height{%.2f}", width, height)
	}
	return width * height,  nil
}

func paramChange(num int)  {
	num =  10
	fmt.Println(num)
}

func paramAddress(num int)  {
	fmt.Println(&num)
}

func paramPoint() {
	var myInt  = 10
	myIntPointer :=  &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	*myIntPointer =  20
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
	var  myFloat float64
	var myFloatPointer *float64
	myFloatPointer  = &myFloat
	fmt.Println(myFloatPointer)
	//myFloatPointer = &myInt，编译异常
}

func funcPoint(numPoint *int) *int {
	var result = *numPoint * 10
	*numPoint = 10
	return &result
}
