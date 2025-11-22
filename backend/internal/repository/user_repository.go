package repository

import (
	"errors"
	"sync"

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
}

type inMemoryUserRepository struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *inMemoryUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Username == user.Username {
			return ErrUserExists
		}
		if u.Email == user.Email {
			return ErrUserExists
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepository) FindByID(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *inMemoryUserRepository) FindByUsername(username string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *inMemoryUserRepository) FindByEmail(email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *inMemoryUserRepository) FindAll() ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}