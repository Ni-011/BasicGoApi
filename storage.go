package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

// constructor 
// form connection with our db and return an instance of PostgresStore with the db
func NewPostgresStore () (*PostgresStore, error) {
	connectionString := "host=localhost port=5432 user=postgres dbname=postgres password=myBank sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

// at init we create account table
func (store *PostgresStore) Init () error {
	return store.CreateAccountTable()
}

// creates a table for account in db
func (store *PostgresStore) CreateAccountTable () error {
	query := `create table if not exists account (
		id serial primary key,
		firstName varchar(50),
		lastName varchar(50),
		phone serial,
		balance serial,
		createdAt timestamp
	)`

	_, err := store.db.Exec(query)
	return err
}

func (store *PostgresStore) CreateAccount(account *Account) error {
	query := `Insert into account (
		firstName, lastName, phone, balance, createdAt
	)
	values ($1, $2, $3, $4, $5)`

	res, err := store.db.Query(query, account.FirstName, account.LastName, account.Phone, account.Balance, account.CreatedAt)
	if err !=nil {
		return nil
	}

	fmt.Printf("%+v\n", res)
	return nil
}

func (store *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

func (store *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (store *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}