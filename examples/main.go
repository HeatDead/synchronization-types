package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func main() {
	wallet := &Wallet{}
	var wg sync.WaitGroup

	// Запускаем 1000 горутин, которые вносят по 1 монете
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wallet.Deposit(1)
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый баланс:", wallet.Balance())
}
