package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Wallet struct {
	balance atomic.Int64 // Атомарный счётчик для баланса
	version atomic.Int32 // Версия данных для оптимистичной проверки
}

// Deposit увеличивает баланс атомарно (без CAS)
func (w *Wallet) Deposit(amount int64) {
	w.balance.Add(amount)
}

// Withdraw списывает средства с проверкой через CAS
func (w *Wallet) Withdraw(amount int64) bool {
	for {
		currentBalance := w.balance.Load()
		if currentBalance < amount {
			return false // Недостаточно средств
		}

		// Сохраняем текущую версию данных
		currentVersion := w.version.Load()

		// Попытка обновить баланс и версию атомарно
		if w.version.CompareAndSwap(currentVersion, currentVersion+1) {
			// Если версия не изменилась, обновляем баланс
			if w.balance.CompareAndSwap(currentBalance, currentBalance-amount) {
				return true
			}
		}

		// Если другая горутина изменила данные, повторяем попытку
	}
}

func (w *Wallet) Balance() int64 {
	return w.balance.Load()
}

func main() {
	wallet := &Wallet{}
	wallet.balance.Store(1000) // Начальный баланс

	var wg sync.WaitGroup

	// 500 горутин пытаются списать по 2 единицы
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
	fmt.Println("Итоговый баланс:", wallet.Balance()) // Ожидается 0
}
