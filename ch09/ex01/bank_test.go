package bank_test

import (
	"fmt"
	"testing"

	bank "./"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Withdraw Success
	go func() {
		bank.WithDraw(100)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Wait for both transactions.
	<-done

	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Withdraw Fail
	go func() {
		bank.WithDraw(300)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Wait for both transactions.
	<-done

	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
