package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T)  {
	wallet := NewWallet()
	address := wallet.GetAddress()
	// 1NSD8JSKDURYEHHMZamcRfJnZcJvMURYhr
	fmt.Println(string(address))
}