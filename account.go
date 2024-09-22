package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Balance   int
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
	}
}
