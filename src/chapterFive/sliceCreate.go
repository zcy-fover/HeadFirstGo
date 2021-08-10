package main

import "fmt"

func main() {
	var mySlice = []int{1, 2, 3}
	notes := make([]string, 10)
	notes[0] = "a"
	notes[1] = "b"
	fmt.Println(mySlice)  //[1 2 3]
	fmt.Println(notes)	//[a b        ]

	var myArray = []int{1, 2, 3, 4, 5, 6,  7, 8, 9}
	sliceArray := myArray[1:5]	//通过数组创建切片
	fmt.Println(sliceArray)    //[2 3 4 5]
	sliceArray = append(sliceArray, 11, 12, 13)
	fmt.Println(sliceArray)		//	[2 3 4 5 11 12 13]

	newSlice := append(mySlice, 4, 5, 6)
	fmt.Println(mySlice)	//[1 2 3]
	fmt.Println(newSlice)	//[1 2 3 4 5 6]
}
