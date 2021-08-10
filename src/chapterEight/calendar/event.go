package calendar

import "fmt"

type Event struct {
	Title string
	Date  //嵌入匿名Date类型，Event可以拥有 Date 属性字段
}

//这里体现了包内非导出数据的可见性
var eventTest = Event{"aa", Date{1, 2, 4}}

func (e Event) DisplayEvent() {
	fmt.Println(eventTest)
}
