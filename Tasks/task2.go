package main

import (
	"fmt"
	"sync"
)

func main() {
	// инициализация
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(arr)
	var wg sync.WaitGroup

	// запуск горутин на каждый элемент slice
	for i := range arr {
		wg.Add(1)
		go func(arr *[]int, pos int) {
			v := (*arr)[pos]
			(*arr)[pos] = v * v
			wg.Done()
		}(&arr, i)
	}

	wg.Wait()
	fmt.Println(arr)
}
