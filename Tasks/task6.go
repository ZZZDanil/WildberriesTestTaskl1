package main

import (
	"fmt"
	"time"
)

func main() {
	// создание горутины с каналом для связи
	stop := make(chan interface{}, 1)
	go goroutine(&stop)
	time.Sleep(500 * time.Millisecond)
	stop <- nil // отправка сигнала в канал
	fmt.Println("Stop, press enter")
	fmt.Scanln()
	// но лучше уничтожить вычислитльное устройство
}

func goroutine(stop *chan interface{}) {
	for {
		if len(*stop) > 0 {
			return
		}
		fmt.Println("goroutine")
	}
}
