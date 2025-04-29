package main

func main() {
	// Создали nil канал. nil потому что не проинициализирован через make
	var ch chan int

	// Закрываем nil канал. Получаем панику.
	close(ch)
}
