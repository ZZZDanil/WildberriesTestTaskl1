package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Println("cap: ", cap(arr))

	RemoveIndex(&arr, 5) // удаление по индексу
	fmt.Println(arr)
	fmt.Println("cap: ", cap(arr))

	// реальный размер не изменился
	// альтернатива - создать через make

}

// объединение левой и правой части слайса
func RemoveIndex(s *[]int, index int) {
	(*s) = append((*s)[:index], (*s)[index+1:]...)
}
