package auth

import (
	"backend/internal/models"
	"errors"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	userRepo   UserRepository
	jwtManager *JWTManager
}

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

func NewAuthService(userRepo UserRepository, jwtManager *JWTManager) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := s.jwtManager.Generate(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.User, error) {
	if _, err := s.userRepo.FindByUsername(req.Username); err == nil {
		return nil, ErrUserExists
	}

	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return nil, ErrUserExists
	}

	user := &models.User{
		ID:       generateID(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func generateID() string {
	return "user-" + string(rune(len("temp")))
}