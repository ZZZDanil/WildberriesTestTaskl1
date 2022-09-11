package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// считывание воркеров
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("Число воркеров: ")
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
			fmt.Printf("Введи число воркеров: ")
		}
	}
	// создание каналов обмена данными
	mainChanel := make(chan string)
	doneChanel := make(chan interface{}, N+1)

	// запуск генератора  и слушателей
	go stringGenerator(&mainChanel, &doneChanel, &wg)
	for i := N; i > 0; i-- {
		go stringChanelListener(&mainChanel, i, &doneChanel, &wg)
	}
	wg.Add(N + 1)

	// обработчик команды завершения процесса
	// если была нажата комбинация CTRL + C,
	// то через канал doneзавершаются все горутины
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(done *chan interface{}) {
		for sig := range c {

			if sig != nil {
				for i := N + 1; i > 0; i-- {
					*done <- 1
				}
				fmt.Println("!!! CTRL + C !!!")
				return
			}
		}

	}(&doneChanel)

	wg.Wait()

}
func stringGenerator(mainChanel *chan string, done *chan interface{}, wg *sync.WaitGroup) {
	for {
		if len(*done) > 0 {
			for {
				if len(*done) == 1 {
					fmt.Println("StringGenerator Stopped")
					wg.Done()
					return
				}
			}
		}
		s := randomString(rand.Intn(10))
		fmt.Println("Put ", s)
		*mainChanel <- s
		time.Sleep(500 * time.Millisecond)

	}
}
func stringChanelListener(mainChanel *chan string, id int, done *chan interface{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-*done:
			fmt.Println("Listener ", id, " Stopped")
			wg.Done()
			return
		case m := <-*mainChanel:
			fmt.Println("Listener ", id, " message: ", m)
		}
	}
}

// генератор строки
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
