// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entries struct {
	Code      int64     `json:"code"`
	AccountID int64     `json:"account_id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfers struct {
	Code          int64     `json:"code"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        string    `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
