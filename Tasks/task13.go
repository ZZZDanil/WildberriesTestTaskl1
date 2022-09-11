package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 99
	)

	fmt.Println("a: ", a, "b: ", b)
	a, b = func(a int, b int) (int, int) {
		return b, a
	}(a, b)

	fmt.Println("a: ", a, "b: ", b)
	a, b = b, a
	fmt.Println("a: ", a, "b: ", b)

}
