package repository

import (
	"errors"

	"backend/internal/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(user *models.User) error
	UpdatePassword(user *models.User) error
}

type OrderRepository interface {
    Create(order *models.Order) error
    FindByID(id string) (*models.Order, error)
    FindByCustomer(customerID string) ([]*models.Order, error)
    FindAll() ([]*models.Order, error)
    Update(order *models.Order) error
    Delete(id string, customerID string) error
    DeleteByID(id string) error
}