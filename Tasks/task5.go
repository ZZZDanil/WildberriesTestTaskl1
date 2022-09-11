package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// анализ ввода
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("Время выполния в секундах: ")
	N := 0
	for sc.Scan() {
		txt := sc.Text()
		i, err := strconv.Atoi(txt)
		if err == nil {
			N = i
			if N > 0 {
				break
			}
		} else {
			fmt.Printf("Введи время выполнения в секундах: ")
		}
	}

	// создание каналов связи
	mainChanel := make(chan string)
	doneChanel := make(chan interface{}, 2)

	wg.Add(2)
	// создание генератора и слушателя
	go stringGenerator(&mainChanel, &doneChanel, &wg)
	go stringChanelListener(&mainChanel, &doneChanel, &wg)

	// ожидание введнного времени
	// и отправка сигналов о завершении программы
	time.Sleep(time.Second * time.Duration(N))
	for i := 2; i > 0; i-- {
		doneChanel <- nil
	}
	wg.Wait()

}
func stringGenerator(mainChanel *chan string, done *chan interface{}, wg *sync.WaitGroup) {
	for {
		// завершение после завершения слушателя
		if len(*done) > 0 {
			for {
				if len(*done) == 1 {
					fmt.Println("StringGenerator Stopped")
					wg.Done()
					break
				}
			}
			break
		}
		s := randomString(rand.Intn(10))
		fmt.Println("Put ", s)
		*mainChanel <- s
		time.Sleep(500 * time.Millisecond)

	}
}
func stringChanelListener(mainChanel *chan string, done *chan interface{}, wg *sync.WaitGroup) {
	for b := 1; b > 0; {
		select {
		case <-*done:
			fmt.Println("Listener ", " Stopped")
			wg.Done()
			b--
		case m := <-*mainChanel:
			fmt.Println("Listener ", " message: ", m)
		}
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
