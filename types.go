package main

import "math/rand"

type Account struct {
	AccountId int    `json:"id`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Phone     int64  `json: "Phone"`
	Balance   int64  `json: "balance"`
}

// return an instance of Account struct
func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		AccountId: rand.Intn(1000),
		FirstName: FirstName,
		LastName:  LastName,
		Phone:     int64(rand.Intn(100000000)),
	}
}
