package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderUser struct {
	ID           pgtype.UUID      `json:"id"`
	OrderID      pgtype.UUID      `json:"order_id"`
	UserID       pgtype.UUID      `json:"user_id"`
	UserUsername string           `json:"user_username"`
	UserEmail    string           `json:"user_email"`
	Message      pgtype.Text      `json:"message"`
	Status       pgtype.Text      `json:"status"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
}