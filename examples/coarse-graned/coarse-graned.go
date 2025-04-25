package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	mu      sync.Mutex
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

func (w *Wallet) Balance() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}

func main() {
	wallet := &Wallet{}
	var wg sync.WaitGroup

	// Те же 1000 горутин
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
