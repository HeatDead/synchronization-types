package main

import (
	"fmt"
	"sync"
)

type Wallet struct {
	once     sync.Once // Контролирует однократную инициализацию
	mu       sync.Mutex
	balance  int
	isInited bool // Флаг инициализации
}

// initBalance выполняется только один раз, даже при конкурентных вызовах
func (w *Wallet) initBalance() {
	fmt.Println("Инициализация баланса...")
	w.balance = 0 // Начальное значение
	w.isInited = true
}

// Deposit вносит средства, инициализируя баланс при первом вызове
func (w *Wallet) Deposit(amount int) {
	w.once.Do(w.initBalance) // Ленивая инициализация

	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

// Balance возвращает баланс, инициализируя его при первом вызове
func (w *Wallet) Balance() int {
	w.once.Do(w.initBalance) // Ленивая инициализация

	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}

func main() {
	wallet := &Wallet{}

	var wg sync.WaitGroup

	// Горутина 1: вызов Balance() до внесения средств
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Баланс (горутина 1):", wallet.Balance()) // 0
	}()

	// Горутина 2: внесение средств
	wg.Add(1)
	go func() {
		defer wg.Done()
		wallet.Deposit(100)
		fmt.Println("Внесено 100")
	}()

	// Горутина 3: вызов Balance() после внесения
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Баланс (горутина 3):", wallet.Balance()) // 100
	}()

	wg.Wait()
}
