package main

import (
	"fmt"
	"reflect"
)

// тип переменной можно узнать при помощи библиотеки reflect

func main() {
	var a interface{} = 1
	var b interface{} = "text"
	whatIsIt(a)
	whatIsIt(b)
	whatIsIt([]byte{1})
}
func whatIsIt(x interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(x), " Value", reflect.ValueOf(x))
}
