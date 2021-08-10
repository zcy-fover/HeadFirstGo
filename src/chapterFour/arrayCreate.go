package main

import (
	"fmt"
)

func main() {
	var ageArray [2]int
	ageArray[1] =  15
	fmt.Println(ageArray[0])	//输出：0
	var nameArray = [10]string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj"}
	fmt.Println(nameArray[0])	//输出：aa
	fmt.Printf("%#v\n", nameArray)

	for i := 0; i < len(nameArray); i++ {
		fmt.Print(nameArray[i], "--")
	}
	fmt.Println()
	for index, value := range nameArray {
		fmt.Println(index, "---", value)
	}
}
