package main

import (
	"fmt"
	"math/rand"
)

func main() {
	массив1 := dataGen(10)
	массив2 := dataGen(15)
	массив3 := пересечениеМассивов(массив1, массив2)

	fmt.Println("массив 1: ", массив1)
	fmt.Println("массив 2: ", массив2)
	fmt.Println("массив 3: ", массив3)
}

func пересечениеМассивов(массив1 *[]int, массив2 *[]int) *[]int {

	// нахождение меньшего массива
	var меньшийМассив *[]int
	var большийМассив *[]int
	if len(*массив1) > len(*массив2) {
		меньшийМассив = массив1
		большийМассив = массив2
	} else {
		меньшийМассив = массив2
		большийМассив = массив1
	}
	// помещение уникальных совпадений
	// между большим и меньшим массивами
	// в 3ю переменную
	вывод := []int{}
	уникальность := map[int]bool{}
	for _, a := range *меньшийМассив {
		for _, b := range *большийМассив {
			if a == b {
				if уникальность[b] != true {
					уникальность[b] = true
					вывод = append(вывод, b)
				}
				break
			}
		}
	}
	return &вывод
}
func dataGen(size int) *[]int {
	d := make([]int, size)
	for i := range d {
		d[i] = rand.Intn(10)
	}
	return &d
}
