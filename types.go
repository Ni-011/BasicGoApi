package main

import (
	"math/rand"
	"time"
)

type Account struct {
	AccountId int    `json:"id`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Phone     int64  `json: "Phone"`
	Balance   int64  `json: "balance"`
	CreatedAt time.Time `json: "createdAt"`
}

type createAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json: "LastName"`
}

// return an instance of Account struct
func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		FirstName: FirstName,
		LastName:  LastName,
		Phone:     int64(rand.Intn(100000000)),
		CreatedAt: time.Now().UTC(),
	}
}
