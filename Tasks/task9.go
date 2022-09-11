package main

import (
	"fmt"
)

func main() {

	// инициализация
	arr := []int{1, 2, 3, 4, 5}

	firstChanel := make(chan int, len(arr))
	secondChanel := make(chan int, len(arr))

	// горутина ввода
	go func(arr *[]int, firstChanel *chan int) {
		// запись в канал
		fmt.Print("In: ")
		for _, v := range *arr {
			*firstChanel <- v
			fmt.Print(v, " ")
		}
		close(*firstChanel) // закрытие канала
	}(&arr, &firstChanel)

	// горутина математической операции
	go func(firstChanel *chan int, secondChanel *chan int) {
		// чтение канала
		for v := range *firstChanel {
			*secondChanel <- v * v
		}
		close(*secondChanel) // закрытие канала
	}(&firstChanel, &secondChanel)

	// горутина вывода
	go func(secondChanel *chan int) {
		// чтение канала и вывод
		fmt.Print("\nOut: ")
		for v := range *secondChanel {
			fmt.Print(v, " ")
		}
	}(&secondChanel)

	fmt.Scanln() // ожидания ввода Enter
}
