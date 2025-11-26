package handlers

import (
	"net/http"

	"backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderUserHandler struct {
	orderUserRepo *postgres.OrderUserRepository
	orderRepo     *postgres.OrderRepository
}

func NewOrderUserHandler(orderUserRepo *postgres.OrderUserRepository, orderRepo *postgres.OrderRepository) *OrderUserHandler {
	return &OrderUserHandler{
		orderUserRepo: orderUserRepo,
		orderRepo:     orderRepo,
	}
}

type CreateOrderUserRequest struct {
	Message string `json:"message"`
}

func (h *OrderUserHandler) CreateOrderUser(c *gin.Context) {
	orderID := c.Query("order_id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req CreateOrderUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	order, err := h.orderRepo.FindByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.CustomerID == userID.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot respond to your own order"})
		return
	}

	hasResponded, err := h.orderUserRepo.CheckOrderUserExists(c, orderID, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check response"})
		return
	}
	if hasResponded {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already responded to this order"})
		return
	}

	orderUser, err := h.orderUserRepo.CreateOrderUser(c, orderID, userID.(string), req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Response created successfully",
		"order_user": gin.H{
			"id":         orderUser.ID.String(),
			"order_id":   orderUser.OrderID.String(),
			"user_id":    orderUser.UserID.String(),
			"message":    orderUser.Message.String,
			"status":     orderUser.Status.String,
			"created_at": orderUser.CreatedAt.Time,
		},
	})
}

func (h *OrderUserHandler) GetOrderUsers(c *gin.Context) {
	orderID := c.Query("order_id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	order, err := h.orderRepo.FindByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.CustomerID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	orderUsers, err := h.orderUserRepo.GetOrderUsers(c, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch responses"})
		return
	}

	var orderUserList []gin.H
	for _, orderUser := range orderUsers {
		orderUserList = append(orderUserList, gin.H{
			"id":            orderUser.ID.String(),
			"order_id":      orderUser.OrderID.String(),
			"user_id":       orderUser.UserID.String(),
			"user_username": orderUser.UserUsername,
			"user_email":    orderUser.UserEmail,
			"message":       orderUser.Message.String,
			"status":        orderUser.Status.String,
			"created_at":    orderUser.CreatedAt.Time,
			"updated_at":    orderUser.UpdatedAt.Time,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"order_users": orderUserList,
		"count":       len(orderUserList),
	})
}

func (h *OrderUserHandler) DeleteOrderUser(c *gin.Context) {
	orderUserID := c.Param("orderUserId")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := h.orderUserRepo.DeleteOrderUser(c, orderUserID, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Response deleted successfully"})
}

func (h *OrderUserHandler) AssignExecutor(c *gin.Context) {    
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    var req struct {
        OrderID string `json:"order_id" binding:"required"`
        UserID  string `json:"user_id" binding:"required"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
        return
    }
    
    order, err := h.orderRepo.FindByID(req.OrderID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if order.CustomerID != userID.(string) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
        return
    }

    orderUser, err := h.orderUserRepo.GetUserOrderUser(c, req.OrderID, req.UserID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User has not responded to this order"})
        return
    }

    updatedOrder, err := h.orderUserRepo.AssignOrderExecutor(c, req.UserID, req.OrderID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign executor"})
        return
    }

    _, err = h.orderUserRepo.UpdateOrderUserStatus(c, orderUser.ID.String(), "accepted")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order user status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Executor assigned successfully",
        "order": gin.H{
            "id":          updatedOrder.ID.String(),
            "customer_id": updatedOrder.CustomerID.String(),
            "title":       updatedOrder.Title,
            "description": updatedOrder.Description,
            "budget":      getBudgetFloat(updatedOrder.Budget),
            "category":    updatedOrder.Category,
            "section":     updatedOrder.Section,
            "status":      updatedOrder.Status.String,
            "created_at":  updatedOrder.CreatedAt.Time,
            "updated_at":  updatedOrder.UpdatedAt.Time,
        },
    })
}

func (h *OrderUserHandler) RemoveExecutor(c *gin.Context) {
	orderID := c.Query("order_id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	order, err := h.orderRepo.FindByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.CustomerID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	updatedOrder, err := h.orderUserRepo.RemoveOrderExecutor(c, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove executor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Executor removed successfully",
		"order": gin.H{
			"id":          updatedOrder.ID.String(),
			"customer_id": updatedOrder.CustomerID.String(),
			"title":       updatedOrder.Title,
			"description": updatedOrder.Description,
			"budget":      getBudgetFloat(updatedOrder.Budget),
			"category":    updatedOrder.Category,
			"section":     updatedOrder.Section,
			"status":      updatedOrder.Status.String,
			"created_at":  updatedOrder.CreatedAt.Time,
			"updated_at":  updatedOrder.UpdatedAt.Time,
		},
	})
}

func (h *OrderUserHandler) CheckOrderUserExists(c *gin.Context) {	
	userID, exists := c.Get("user_id")	
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orderID := c.Query("order_id")	
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order_id query parameter is required"})
		return
	}
	
	hasResponded, err := h.orderUserRepo.CheckOrderUserExists(c, orderID, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"has_responded": hasResponded})
}

func getBudgetFloat(budget pgtype.Numeric) float64 {
	if budget.Valid {
		floatValue, err := budget.Float64Value()
		if err == nil && floatValue.Valid {
			return floatValue.Float64
		}
	}
	return 0
}