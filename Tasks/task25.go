package main

import (
	"fmt"
	"time"
)

func main() {
	// работа с библиотекой time

	fmt.Println("...")
	mySleep(1)
	fmt.Println("---")
	mySleep2(1)
	fmt.Println("+++")
}

func mySleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}
func mySleep2(sec int) {
	t := time.Now()
	for time.Now().Second()-t.Second() < sec {

	}
}
