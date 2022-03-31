package main

import (
	"fmt"
	"sync"
)

var balance int = 0

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 1; i <= 500; i++ {
		wg.Add(1)
		go Deposit(1, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(Balance())
}
