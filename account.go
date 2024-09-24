package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstNumber"`
	LastName  string `json:"lastName"`
	Balance   int    `json:"balance"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
	}
}
