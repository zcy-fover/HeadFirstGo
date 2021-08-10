package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target :=  50
	seed  := time.Now().Unix()
	rand.Seed(seed)
	for i := 10; i > 0; i-- {
		guessTemp := rand.Intn(100)  +  1
		if guessTemp ==  target {
			fmt.Println("bingo!", guessTemp)
			continue
		} else if guessTemp > target {
			fmt.Println("large", guessTemp)
		} else {
			fmt.Println("small", guessTemp)
		}
	}
}