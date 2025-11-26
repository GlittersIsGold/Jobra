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

type OrderUserRepository interface {
	Create(orderUser *models.OrderUser) error
	FindByID(id string) (*models.OrderUser, error)
	FindByOrderID(orderID string) ([]*models.OrderUser, error)
	FindByUserID(userID string) ([]*models.OrderUser, error)
	FindByOrderAndUser(orderID, userID string) (*models.OrderUser, error)
	ExistsByOrderAndUser(orderID, userID string) (bool, error)
	Update(orderUser *models.OrderUser) error
	Delete(id string) error
	DeleteByOrderAndUser(orderID, userID string) error
	AssignExecutor(orderID, userID string) error
	RemoveExecutor(orderID string) error
}