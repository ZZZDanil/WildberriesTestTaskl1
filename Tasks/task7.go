package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// карта
	var m map[string]string = make(map[string]string)
	// канал связи
	stop := make(chan interface{}, 3)
	// запуск горутин
	go goroutine(&m, &stop, 1)
	go goroutine(&m, &stop, 2)
	go goroutine(&m, &stop, 3)

	time.Sleep(5 * time.Millisecond)
	// остановка
	stop <- nil
	stop <- nil
	stop <- nil
	fmt.Println("Stop, press enter")
	fmt.Scanln()

}
func goroutine(m *map[string]string, stop *chan interface{}, id int) {
	for {
		if len(*stop) > 0 {
			return
		}
		key := randomString(rand.Intn(6))
		value := randomString(rand.Intn(6))
		fmt.Println("if: ", id, " key: ", key, " value: ", value)
		(*m)[key] = value
		//time.Sleep(100 * time.Millisecond)
	}
}
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
