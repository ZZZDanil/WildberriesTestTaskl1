package main

import (
	"fmt"
	"unsafe"
)

type A interface {
	To()
}

type B struct {
}

func (b *B) To() {
	fmt.Println("sssssss")
}

type T struct {
	a bool
}

func main2() {
	var ggg A
	yyy := B{}
	ggg = &yyy
	ggg.To()

	s := unsafe.Sizeof(struct{}{})
	fmt.Println(s)
	s = unsafe.Sizeof(T{true})
	fmt.Println(s)

	m := map[int]int{}
	m[0] = 1
	m[1] = 124
	m[2] = 281
	fmt.Println(m[0], m[1], m[2])

	fmt.Println(m)

}
func main() {
	slice := []string{"a", "a"}
	func(slice *[]string) {
		*slice = append((*slice), "a")
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		fmt.Print(slice)
	}(&slice)
	fmt.Print(slice)
}
