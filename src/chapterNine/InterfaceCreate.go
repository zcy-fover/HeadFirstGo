package main

import "HeadFirstGo/src/chapterNine/music"

func main() {
	var whistle music.MyInterface
	whistle = music.Whistle{Name: "口哨"}
	whistle.MusicName()
	whistle.PlayMusic()

	var piano music.MyInterface = music.Piano{Name: "钢琴"}
	piano.MusicName()
	piano.PlayMusic()
	//piano.Other() 编译错误
	//使用类型断言获取类型具体的值，ok 表示类型断言是否成功
	pianoType, ok := piano.(music.Piano)
	if ok {
		pianoType.Other() //这样可以调用类型自己的方法
	}
}
