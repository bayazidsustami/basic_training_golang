package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditionWithMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter :", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (bankAccount *BankAccount) AddBalance(amount int) {
	bankAccount.RWMutex.Lock()
	bankAccount.Balance = bankAccount.Balance + amount
	bankAccount.RWMutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.RWMutex.RLock()
	balance := bankAccount.Balance
	bankAccount.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("total :", account.GetBalance())
}
