package main

import (
	"fmt"
	"sort"
)

func main() {
	// сортировка
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(arr)
	fmt.Println(arr)
	// функция вывода slice,
	// где разница между большим и меньшим элементом равна длине шага
	f := func(v *[]float64, step int) {
		f := float64(step)
		buf := 0
		for i := range *v {
			if (*v)[i] > ((*v)[buf] + f) {
				fmt.Println((*v)[buf:i]) // вывод
				buf = i                  // смещение опорной точки
			}
		}
	}
	f(&arr, 10)
}
