package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

var justString string
var v3 *string
var rtm runtime.MemStats
var bufAlloc int64

func someFunc() {
	var (
		startAlloc           int64
		createBigStringAlloc int64
		createSliceAlloc     int64
		updateBigStringAlloc int64
	)

	runtime.GC()
	startAlloc = getAlloc()

	v := createHugeString(1 << 10)
	createBigStringAlloc = getAlloc()
	justString = v[:100]

	createSliceAlloc = getAlloc()
	v = createHugeString(1 << 10)

	runtime.GC()
	updateBigStringAlloc = getAlloc()

	fmt.Println("createBigStringAlloc - startAlloc: ", (createBigStringAlloc - startAlloc))
	fmt.Println("createSliceAlloc - createBigStringAlloc: ", (createSliceAlloc - startAlloc))
	fmt.Println("updateBigStringAlloc - createSliceAlloc: ", (updateBigStringAlloc - startAlloc))

	// justString имеет ссылку на старую строку
	// https://go.dev/blog/slices-intro
	// https://go.dev/blog/slices

	fmt.Println("---")

	runtime.GC()
	startAlloc = getAlloc()

	v2 := createHugeString(1 << 10)
	createBigStringAlloc = getAlloc()

	justString = string([]byte(v2[:100]))

	createSliceAlloc = getAlloc()

	v2 = createHugeString(1 << 10)
	runtime.GC()
	updateBigStringAlloc = getAlloc()

	fmt.Println("createBigStringAlloc - startAlloc: ", (createBigStringAlloc - startAlloc))
	fmt.Println("createSliceAlloc - createBigStringAlloc: ", (createSliceAlloc - startAlloc))
	fmt.Println("updateBigStringAlloc - createSliceAlloc: ", (updateBigStringAlloc - startAlloc))

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

func getAlloc() int64 {
	runtime.ReadMemStats(&rtm)
	return int64(rtm.Alloc)
}
