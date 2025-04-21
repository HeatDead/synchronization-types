package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	balanceMu sync.Mutex // Мьютекс для баланса
	balance   int

	historyMu sync.Mutex // Мьютекс для истории
	history   []string   // История транзакций
}

func (w *Wallet) Deposit(amount int) {
	w.balanceMu.Lock()
	defer w.balanceMu.Unlock()
	w.balance += amount

	w.historyMu.Lock()
	defer w.historyMu.Unlock()
	w.history = append(w.history, fmt.Sprintf("Депозит: +%d", amount))
}

func (w *Wallet) Withdraw(amount int) bool {
	w.balanceMu.Lock()
	defer w.balanceMu.Unlock()

	w.balance -= amount

	w.historyMu.Lock()
	defer w.historyMu.Unlock()
	w.history = append(w.history, fmt.Sprintf("Снятие: -%d", amount))
	return true
}

func (w *Wallet) Balance() int {
	w.balanceMu.Lock()
	defer w.balanceMu.Unlock()
	return w.balance
}

func (w *Wallet) History() []string {
	w.historyMu.Lock()
	defer w.historyMu.Unlock()
	return append([]string{}, w.history...)
}

func main() {
	wallet := &Wallet{}
	var wg sync.WaitGroup

	// 500 депозитов и 500 попыток снятия
	for i := 0; i < 500; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			wallet.Deposit(1)
		}()
		go func() {
			defer wg.Done()
			wallet.Withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Println("Баланс:", wallet.Balance())
	fmt.Println("Последние 5 операций:", wallet.History()[:5])
}
