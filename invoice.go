package kcrps

import (
	"github.com/google/uuid"
	"time"
)

type Invoice struct {
	ID         int       `json:"id" db:"id"`
	UUID       int       `json:"uuid" db:"uuid"`
	PosID      uuid.UUID `json:"pos-id" db:"pos_id"`
	UserID     int       `db:"user_id"`
	CreatedAt  time.Time `json:"created-at" db:"created_at"`
	Account    string    `json:"account" db:"account" binding:"required"`
	Amount     int       `json:"amount" db:"amount" binding:"required"`
	ClientName string    `json:"client-name" db:"client_name"`
	Message    string    `json:"message" db:"message"`
	Status     int       `json:"status" db:"status"`
}
