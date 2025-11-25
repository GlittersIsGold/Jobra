package postgres

import (
	"context"

	"backend/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderUserRepository struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewOrderUserRepository(pool *pgxpool.Pool) *OrderUserRepository {
	return &OrderUserRepository{
		pool:    pool,
		queries: db.New(pool),
	}
}

func toUUID(s string) pgtype.UUID {
	var uuid pgtype.UUID
	uuid.Scan(s)
	return uuid
}

func toText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: s != ""}
}

func (r *OrderUserRepository) CreateOrderUser(ctx context.Context, orderID, userID, message string) (db.OrderUser, error) {
	return r.queries.CreateOrderUser(ctx, db.CreateOrderUserParams{
		OrderID: toUUID(orderID),
		UserID:  toUUID(userID),
		Message: toText(message),
	})
}

func (r *OrderUserRepository) GetOrderUsers(ctx context.Context, orderID string) ([]db.GetOrderUsersRow, error) {
	return r.queries.GetOrderUsers(ctx, toUUID(orderID))
}

func (r *OrderUserRepository) GetUserOrderUser(ctx context.Context, orderID, userID string) (db.OrderUser, error) {
	return r.queries.GetUserOrderUser(ctx, db.GetUserOrderUserParams{
		OrderID: toUUID(orderID),
		UserID:  toUUID(userID),
	})
}

func (r *OrderUserRepository) UpdateOrderUserStatus(ctx context.Context, orderUserID, status string) (db.OrderUser, error) {
	return r.queries.UpdateOrderUserStatus(ctx, db.UpdateOrderUserStatusParams{
		Status: toText(status),
		ID:     toUUID(orderUserID),
	})
}

func (r *OrderUserRepository) DeleteOrderUser(ctx context.Context, orderUserID, userID string) error {
	return r.queries.DeleteOrderUser(ctx, db.DeleteOrderUserParams{
		ID:     toUUID(orderUserID),
		UserID: toUUID(userID),
	})
}

func (r *OrderUserRepository) AssignOrderExecutor(ctx context.Context, executorID, orderID string) (db.Order, error) {
	return r.queries.AssignOrderExecutor(ctx, db.AssignOrderExecutorParams{
		ExecutorID: toUUID(executorID),
		ID:         toUUID(orderID),
	})
}

func (r *OrderUserRepository) RemoveOrderExecutor(ctx context.Context, orderID string) (db.Order, error) {
	return r.queries.RemoveOrderExecutor(ctx, toUUID(orderID))
}

func (r *OrderUserRepository) CheckOrderUserExists(ctx context.Context, orderID, userID string) (bool, error) {
	return r.queries.CheckOrderUserExists(ctx, db.CheckOrderUserExistsParams{
		OrderID: toUUID(orderID),
		UserID:  toUUID(userID),
	})
}

func (r *OrderUserRepository) GetOrderWithDetails(ctx context.Context, orderID string) (db.GetOrderWithDetailsRow, error) {
	return r.queries.GetOrderWithDetails(ctx, toUUID(orderID))
}