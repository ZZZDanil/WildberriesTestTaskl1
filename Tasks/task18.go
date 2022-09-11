package main

import (
	"fmt"
	"time"
)

type Счётчик struct {
	Число     int
	обращения chan interface{}
	остановка chan interface{}
}

func main() {

	// горутина с постоянной отправкой значений в канал
	горутинныйУвеличитель := func(канал *chan interface{}, остановка chan interface{}) {
		for {
			select {
			case <-остановка:
				return
			default:
				*канал <- nil
			}
		}

	}

	// инициализация и запуск счётчика
	мойСчётчик := Счётчик{Число: 0}
	канал := мойСчётчик.старт()
	// создание и связывание горутин по увеличению счётчика
	канальнаяОстановка := make(chan interface{})

	go горутинныйУвеличитель(канал, канальнаяОстановка)
	go горутинныйУвеличитель(канал, канальнаяОстановка)
	go горутинныйУвеличитель(канал, канальнаяОстановка)

	time.Sleep(500 * time.Millisecond)

	// остановка горутин
	канальнаяОстановка <- nil
	канальнаяОстановка <- nil
	канальнаяОстановка <- nil
	// остановка счётчика
	мойСчётчик.стоп()

	fmt.Println("Число: ", мойСчётчик.Число)

}
func (сч *Счётчик) увеличение() {
	сч.Число++
}

// функция запуска счётчика
// происходит чтение из канала запросов на увеличение
func (сч *Счётчик) старт() *chan interface{} {
	сч.обращения = make(chan interface{})
	сч.остановка = make(chan interface{})
	go func(обр *chan interface{}, ост *chan interface{}) {
		for {
			select {
			case <-*обр:
				сч.Число++
			case <-*ост:
				return
			}

		}
	}(&сч.обращения, &сч.остановка)
	return &сч.обращения
}

// функция остановки счётчика
// в канал помещается сообщение
// и все связанные по каналу горутины останавливаются
func (сч *Счётчик) стоп() {
	сч.остановка <- nil
}