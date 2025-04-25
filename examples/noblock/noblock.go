package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Wallet struct {
	balance atomic.Int64 // Атомарный счётчик баланса
}

// Deposit атомарно увеличивает баланс
func (w *Wallet) Deposit(amount int64) {
	w.balance.Add(amount)
}

// Withdraw атомарно списывает средства (с проверкой через CAS)
func (w *Wallet) Withdraw(amount int64) bool {
	for {
		current := w.balance.Load()
		if current < amount {
			return false // Недостаточно средств
		}

		// Пытаемся атомарно обновить баланс, если он не изменился
		if w.balance.CompareAndSwap(current, current-amount) {
			return true
		}
		// Если другая горутина изменила баланс, повторяем попытку
	}
}

func (w *Wallet) Balance() int64 {
	return w.balance.Load()
}

func main() {
	wallet := &Wallet{}
	wallet.balance.Store(1000) // Начальный баланс

	var wg sync.WaitGroup

	// 500 горутин списывают по 2 единицы
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for !wallet.Withdraw(2) {
				// Повторяем, пока не получится
			}
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый баланс:", wallet.Balance()) // 0
}
