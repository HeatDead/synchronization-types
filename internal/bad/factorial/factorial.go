package factorial

func factorial(n int, ch chan int) {
	if n == 0 || n == 1 {
		ch <- 1
		return
	}
	ch <- n * <-ch // Не защищенный доступ к каналу
}
