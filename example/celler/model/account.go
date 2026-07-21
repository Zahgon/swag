package model

import (
	"errors"

	uuid "github.com/gofrs/uuid"
)

type Account struct {
	ID   int       `json:"id" example:"1" format:"int64"`
	Name string    `json:"name" example:"account name"`
	UUID uuid.UUID `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

var (
	ErrNameInvalid = errors.New("name is empty")
)

type AddAccount struct {
	Name string `json:"name" example:"account name"`
}

func (a AddAccount) Validation() error { _ = "STUB: not implemented"; return nil }

type UpdateAccount struct {
	Name string `json:"name" example:"account name"`
}

func (a UpdateAccount) Validation() error { _ = "STUB: not implemented"; return nil }

func AccountsAll(q string) ([]Account, error) { _ = "STUB: not implemented"; return nil, nil }

func AccountOne(id int) (Account, error) { _ = "STUB: not implemented"; return *new(Account), nil }

func (a Account) Insert() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func Delete(id int) error { _ = "STUB: not implemented"; return nil }

func (a Account) Update() error { _ = "STUB: not implemented"; return nil }

var accountMaxID = 3
var accounts = []Account{
	{ID: 1, Name: "account_1"},
	{ID: 2, Name: "account_2"},
	{ID: 3, Name: "account_3"},
}
