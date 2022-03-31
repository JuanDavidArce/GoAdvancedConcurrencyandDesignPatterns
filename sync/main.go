package main

import (
	"fmt"
	"sync"
)

var balance int = 0

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

// 1 Deposit()
// N Balance()

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 500; i++ {
		wg.Add(1)
		go Deposit(1, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(Balance(&lock))
}
