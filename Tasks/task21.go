package main

import "fmt"

// старя реализация чего-то
type OldInterface interface {
	oldHelloFunc() *string
}

type OldRealisation struct {
}

// старая функция
func (o *OldRealisation) oldHelloFunc() *string {
	s := "Здрасте!"
	return &s
}

// адаптер для старой реализации
// по выводу старых данных в новой форме
type IAdapter interface {
	normalHelloFunc()
}
type Adapter struct {
	OldInterface interface{}
	text         string
}

func (a *Adapter) print() *string {
	s := a.text + "Здрасте!"
	return &s
}

func main() {
	old := OldRealisation{}          // старая реализация
	new := Adapter{old, "Привет и "} // адаптер
	fmt.Println(*new.print())        // новые данные

}
