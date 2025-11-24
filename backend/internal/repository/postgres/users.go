package postgres

import (
	"context"
	"backend/internal/db"
	"backend/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(pool),
	}
}

func (r *UserRepository) Create(user *models.User) error {
	ctx := context.Background()
	
	dbUser, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
	if err != nil {
		return err
	}
	
	user.ID = dbUser.ID.String()
	user.CreatedAt = dbUser.CreatedAt.Time
	user.UpdatedAt = dbUser.UpdatedAt.Time
	return nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	ctx := context.Background()
	
	var userID pgtype.UUID
	if err := userID.Scan(id); err != nil {
		return nil, err
	}
	
	dbUser, err := r.queries.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	return toModelUser(&dbUser), nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	ctx := context.Background()
	
	dbUser, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	
	return toModelUser(&dbUser), nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	
	return toModelUser(&dbUser), nil
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	ctx := context.Background()
	
	dbUsers, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	
	users := make([]*models.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = toModelUser(&dbUser)
	}
	
	return users, nil
}

func toModelUser(dbUser *db.User) *models.User {
	return &models.User{
		ID:           dbUser.ID.String(),
		Username:     dbUser.Username,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		CreatedAt:    dbUser.CreatedAt.Time,
		UpdatedAt:    dbUser.UpdatedAt.Time,
	}
}

func (r *UserRepository) Update(user *models.User) error {
	ctx := context.Background()
	
	var userID pgtype.UUID
	if err := userID.Scan(user.ID); err != nil {
		return err
	}
	
	dbUser, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:       userID,
		Username: user.Username,
		Email:    user.Email,
	})
	if err != nil {
		return err
	}
	
	user.UpdatedAt = dbUser.UpdatedAt.Time
	return nil
}

func (r *UserRepository) UpdatePassword(user *models.User) error {
    ctx := context.Background()
    
    var userID pgtype.UUID
    if err := userID.Scan(user.ID); err != nil {
        return err
    }
    
    dbUser, err := r.queries.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
        ID:           userID,
        PasswordHash: user.PasswordHash,
    })
    if err != nil {
        return err
    }
    
    user.UpdatedAt = dbUser.UpdatedAt.Time
    return nil
}