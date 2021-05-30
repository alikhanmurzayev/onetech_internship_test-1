# calculator

Необходимо реализовать метод *Start* структуры *Calculator*.
Структура содержит два поля:
```golang
type Calculator struct {
	Input  <-chan int
	Output chan<- int
}
```
При вызове метода *Start* необходимо запустить поцесс, который будет читать из канала *Input* числа и записывать их квадраты в канал *Output*. При этом функция должна сразу завершиться. Считывать числа из канала следует до тех пор, пока он открыт. После закрытия канала *Input* необходимо закрыть канал *Output* и завершить работу.

* [Работа с горутинами](https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/)
* [Пример работы с каналами](https://tour.golang.org/concurrency/4)
