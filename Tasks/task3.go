package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup

	// запуск горутин на каждый элемент slice
	// запист результата в sum
	for i := range arr {
		wg.Add(1)
		go func(arr *[]int, pos int, sum *int) {
			v := (*arr)[pos]
			*sum += v * v
			defer wg.Done()
		}(&arr, i, &sum)
	}
	wg.Wait()
	fmt.Println("sum: ", sum)
}
