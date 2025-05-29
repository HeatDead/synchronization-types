package factorial

func factorial(n int, ch chan int) {
	if n == 0 || n == 1 {
		ch <- 1
		return
	}

	// Создаем новый канал для передачи результата от нижестоящей горутины
	nextCh := make(chan int)
	go factorial(n-1, nextCh) // Вызываем горутину для n-1

	// Получаем результат из следующей горутины и вычисляем текущий факториал
	ch <- n * <-nextCh
}
