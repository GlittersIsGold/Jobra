package postgres

import (
	"backend/internal/db"
	"backend/internal/models"
	"context"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	queries *db.Queries
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		queries: db.New(pool),
	}
}

func (r *OrderRepository) Create(order *models.Order) error {
	ctx := context.Background()
	
	var customerID pgtype.UUID
	if err := customerID.Scan(order.CustomerID); err != nil {
		return err
	}
	
	budget := pgtype.Numeric{}
	if err := budget.Scan(order.Budget); err != nil {
		budget = pgtype.Numeric{
			Int:   big.NewInt(int64(order.Budget * 100)), 
			Exp:   -2, 
			Valid: true,
		}
	}
	
	status := pgtype.Text{String: order.Status, Valid: true}
	
	dbOrder, err := r.queries.CreateOrder(ctx, db.CreateOrderParams{
		CustomerID:  customerID,
		Title:       order.Title,
		Description: order.Description,
		Budget:      budget,
		Category:    order.Category,
		Section:     order.Section,
		Status:      status,
	})
	if err != nil {
		return err
	}
	
	order.ID = dbOrder.ID.String()
	order.CreatedAt = dbOrder.CreatedAt.Time
	order.UpdatedAt = dbOrder.UpdatedAt.Time
	return nil
}

func (r *OrderRepository) FindByID(id string) (*models.Order, error) {
	ctx := context.Background()
	
	var orderID pgtype.UUID
	if err := orderID.Scan(id); err != nil {
		return nil, err
	}
	
	dbOrder, err := r.queries.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	
	return toModelOrderFromSingle(&dbOrder), nil
}

func (r *OrderRepository) FindByCustomer(customerID string) ([]*models.Order, error) {
	ctx := context.Background()
	
	var customerUUID pgtype.UUID
	if err := customerUUID.Scan(customerID); err != nil {
		return nil, err
	}
	
	dbOrders, err := r.queries.GetOrdersByCustomer(ctx, customerUUID)
	if err != nil {
		return nil, err
	}
	
	orders := make([]*models.Order, len(dbOrders))
	for i, dbOrder := range dbOrders {
		orders[i] = toModelOrderFromCustomerList(&dbOrder)
	}
	
	return orders, nil
}

func (r *OrderRepository) FindAll() ([]*models.Order, error) {
	ctx := context.Background()
	
	dbOrders, err := r.queries.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	
	orders := make([]*models.Order, len(dbOrders))
	for i, dbOrder := range dbOrders {
		orders[i] = toModelOrderFromAllList(&dbOrder)
	}
	
	return orders, nil
}

func (r *OrderRepository) Update(order *models.Order) error {
	ctx := context.Background()
	
	var orderID pgtype.UUID
	if err := orderID.Scan(order.ID); err != nil {
		return err
	}
	
	var budget pgtype.Numeric
	if err := budget.Scan(order.Budget); err != nil {
		return err
	}
	
	status := pgtype.Text{String: order.Status, Valid: true}
	
	dbOrder, err := r.queries.UpdateOrder(ctx, db.UpdateOrderParams{
		ID:          orderID,
		Title:       order.Title,
		Description: order.Description,
		Budget:      budget,
		Category:    order.Category,
		Section:     order.Section,
		Status:      status,
	})
	if err != nil {
		return err
	}
	
	order.UpdatedAt = dbOrder.UpdatedAt.Time
	return nil
}

func (r *OrderRepository) Delete(id string, customerID string) error {
	ctx := context.Background()
	
	var orderUUID, customerUUID pgtype.UUID
	if err := orderUUID.Scan(id); err != nil {
		return err
	}
	if err := customerUUID.Scan(customerID); err != nil {
		return err
	}
	
	return r.queries.DeleteOrder(ctx, db.DeleteOrderParams{
		ID:         orderUUID,
		CustomerID: customerUUID,
	})
}

func (r *OrderRepository) DeleteByID(id string) error {
	ctx := context.Background()
	
	var orderUUID pgtype.UUID
	if err := orderUUID.Scan(id); err != nil {
		return err
	}
	
	return r.queries.DeleteOrderByID(ctx, orderUUID)
}

func toModelOrderFromSingle(dbOrder *db.GetOrderByIDRow) *models.Order {
	var budget float64
	if dbOrder.Budget.Valid {
		floatValue, err := dbOrder.Budget.Float64Value()
		if err == nil && floatValue.Valid {
			budget = floatValue.Float64
		}
	}
	
	return &models.Order{
		ID:          dbOrder.ID.String(),
		CustomerID:  dbOrder.CustomerID.String(),
		Title:       dbOrder.Title,
		Description: dbOrder.Description,
		Budget:      budget,
		Category:    dbOrder.Category,
		Section:     dbOrder.Section,
		Status:      dbOrder.Status.String,
		CreatedAt:   dbOrder.CreatedAt.Time,
		UpdatedAt:   dbOrder.UpdatedAt.Time,
	}
}

func toModelOrderFromCustomerList(dbOrder *db.GetOrdersByCustomerRow) *models.Order {
	var budget float64
	if dbOrder.Budget.Valid {
		floatValue, err := dbOrder.Budget.Float64Value()
		if err == nil && floatValue.Valid {
			budget = floatValue.Float64
		}
	}
	
	return &models.Order{
		ID:          dbOrder.ID.String(),
		CustomerID:  dbOrder.CustomerID.String(),
		Title:       dbOrder.Title,
		Description: dbOrder.Description,
		Budget:      budget,
		Category:    dbOrder.Category,
		Section:     dbOrder.Section,
		Status:      dbOrder.Status.String,
		CreatedAt:   dbOrder.CreatedAt.Time,
		UpdatedAt:   dbOrder.UpdatedAt.Time,
	}
}

func toModelOrderFromAllList(dbOrder *db.GetAllOrdersRow) *models.Order {
	var budget float64
	if dbOrder.Budget.Valid {
		floatValue, err := dbOrder.Budget.Float64Value()
		if err == nil && floatValue.Valid {
			budget = floatValue.Float64
		}
	}
	
	return &models.Order{
		ID:          dbOrder.ID.String(),
		CustomerID:  dbOrder.CustomerID.String(),
		Title:       dbOrder.Title,
		Description: dbOrder.Description,
		Budget:      budget,
		Category:    dbOrder.Category,
		Section:     dbOrder.Section,
		Status:      dbOrder.Status.String,
		CreatedAt:   dbOrder.CreatedAt.Time,
		UpdatedAt:   dbOrder.UpdatedAt.Time,
	}
}