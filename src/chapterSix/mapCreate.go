package main

import (
	stringUtils "HeadFirstGo/src/common"
	"fmt"
	"log"
)

func main() {
	lines, err := stringUtils.GetStrings("src/chapterSix/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var myMap map[string]int
	myMap = make(map[string]int)
	fmt.Println(myMap)
	initMap := map[string]int{"a": 1, "b": 2}
	value, ok := initMap["a"]
	if ok {
		fmt.Println("存在：", value)
	} else {
		fmt.Println("不存在")
	}
	delete(initMap, "b")
	fmt.Println(initMap)
}
