package main

import "github.com/google/uuid"

type Account struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
