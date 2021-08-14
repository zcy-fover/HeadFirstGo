package music

import "fmt"

type Piano struct {
	Name string
}

func (p Piano) PlayMusic() {
	fmt.Println(p.Name, "演奏钢琴曲")
}

func (p Piano) MusicName() {
	fmt.Println("我是钢琴")
}

func (p Piano) Other() {
	fmt.Println("钢琴可以演奏其他的吗？")
}
