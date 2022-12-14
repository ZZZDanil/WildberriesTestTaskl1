# Вопросы
1. Какой самый эффективный способ конкатенации строк?

Ответ:

strings.Builder 

2. Что такое интерфейсы, как они применяются в Go?

Интерфейс - динамический тип, который способен иметь соглашения о функциях.

Как данные - интерфейс используется как универсальный тип данных без привязки к типу.

Как соглашение - интерфейс используется для полиморфизма

3. Чем отличаются RWMutex от Mutex?

RWMutex в отличие от Mutex имеет возможность параллельного считывания данных во время блокировки.

4. Чем отличаются буферизированные и не буферизированные каналы?

Не буферизированные каналы имеют 1 свободную ячейку под данные, заполнив которую, канал блокируется до считывания данных.

Буфферизированный имеет Т ячеек и блокируется после заполнения их всех.

5. Какой размер у структуры struct{}{}?

unsafe.Sizeof показывает 0

6. Есть ли в Go перегрузка методов или операторов?

Нет

7. В какой последовательности будут выведены элементы map[int]int?

```go
m[0]=1
m[1]=124
m[2]=281
```

Выведятся в последовательности вывода по ключам

8. В чем разница make и new?

make - создаёт map, slice, chan. Имеет входной параметр Тип и Размер области

new - создаёт указатель на новый тип. Имеет входной параметр Тип 

9. Сколько существует способов задать переменную типа slice или map?

Фантазия безгранична

10. Что выведет данная программа и почему?

```go
func update(p *int) {
b := 2
p = &b
}
func main() {
var (
a = 1
p = &a
) fmt.Println(*p)
update(p)
fmt.Println(*p)
}
```

Вывод: 1 1

Так как функции передают параметры по значению.

Заменив p = &b => \*p = b будет: 1 2

11. Что выведет данная программа и почему?

```go
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}
```

fatal error: all goroutines are asleep - deadlock!

Необходимо передать указатель на wg

12. Что выведет данная программа и почему?

```go
func main() {
n := 0
if true {
n := 1
n++
} fmt.Println(n)
}
```

Вывод: 0

Создана локальная переменная

n := 1 => n = 1

13. Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
```

Вывод: [100 2 3 4 5]

Замена элемента и расширение slice

Заменив передачу по указатель получен вид

Вывод: [100 2 3 4 5 6]

14. Что выведет данная программа и почему?

```go
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
```

Вывод: [b b a][a a]

Передача копии slice, создание локальной копии и вывод в консоль 2х slice

Заменив передачу по указателю создании локальной копии предотвращается

