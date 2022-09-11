package main

import (
	"fmt"
)

func main() {
	var num int64
	num = 1
	fmt.Println("res: ", changeBit(&num, 3))
}

func changeBit(value *int64, i int) int64 {
	return *value ^ 1<<i // битовое изменение на основе i-й маски
}
