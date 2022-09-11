package main

import "fmt"

func main() {
	s := "главрыба"
	fmt.Println(s)
	fmt.Println(*переворот(&s))
}

// переход от строки к массивам рун
// последовательный проход от начала до конца массива
func переворот(input *string) *string {
	r := []rune(*input)
	l := len(r) - 1
	out := make([]rune, l+1)
	for i := 0; i <= l; i++ {
		out[i] = r[l-i]
	}
	s := string(out)
	return &s
}
