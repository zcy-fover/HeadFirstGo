package main

import (
	"HeadFirstGo/src/chapterEight/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetDate(2021, 8, 10)
	if err != nil {
		log.Fatal(err)
	}
	date.Display()

	event := calendar.Event{}
	event.Title = "测试标题"
	err = event.SetDate(2, 3, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year(), event.Month(), event.Day())

	event.DisplayEvent()
}
