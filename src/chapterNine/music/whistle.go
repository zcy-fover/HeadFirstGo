package music

import "fmt"

type Whistle struct {
	Name string
}

func (w Whistle) PlayMusic() {
	fmt.Println(w.Name, "可以吹哨子")
}

func (w Whistle) MusicName() {
	fmt.Println("我是口哨")
}
