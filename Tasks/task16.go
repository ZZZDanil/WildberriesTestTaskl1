package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// Генерация данных
	data := dataGen(100)
	fmt.Println((*data)[:])

	// Подготовка к ожиданию
	var wg sync.WaitGroup
	wg.Add(1)

	// запуск горутины сортировки
	go sort(data, 0, (len(*data) - 1), &wg)

	// Финальные действия после сортировки
	wg.Wait()
	fmt.Println((*data)[:])

}

func sort(data *[]int, l int, r int, wg *sync.WaitGroup) {
	lenght := r - l

	if lenght > 0 {
		// определения
		var center int
		center = lenght / 2
		p := (*data)[l+center]
		left := l
		right := r
		// перемещение от краёв массива к *центру*
		for left <= right {
			// находятся с краёв массива объекты для перестановки
			for left <= r && (*data)[left] < p {
				left++
			}
			for right >= l && (*data)[right] > p {
				right--
			}
			// перестановка
			if left <= right {
				(*data)[left], (*data)[right] = (*data)[right], (*data)[left]
				left++
				right--
			}

		}
		wg.Add(1)
		// разделение массива на 2 части
		go sort(data, l, left-1, wg) // элементы меньше опорного
		go sort(data, left, r, wg)   // элементы больше опорного
	} else {
		wg.Done()
	}
}

// генератор
func dataGen(size int) *[]int {
	d := make([]int, size)
	for i := range d {
		d[i] = rand.Intn(1000)
	}
	return &d
}
