package main

import (
	"math/rand"
)

var justString string

func someFunc() {

	v := createHugeString(1 << 10)

	// создание slice со ссылкой на большую строку
	// может привести к утечки памяти
	// justString = v[:100]
	// создание новой строки копии от части большой строки
	justString = string([]byte(v[:100]))

}
func main() {
	someFunc()
}

func createHugeString(v int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, v)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
