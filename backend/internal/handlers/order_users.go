package handlers

import (
	"log"
	"net/http"
	"time"

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
    log.Println("=== AssignExecutor handler called ===")
    log.Printf("Method: %s, Path: %s", c.Request.Method, c.Request.URL.Path)
    
    var req struct {
        OrderID string `json:"order_id"`
        UserID  string `json:"user_id"`
    }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Bind JSON error: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
        return
    }
    
    log.Printf("Request data: order_id=%s, user_id=%s", req.OrderID, req.UserID)
    
    // Проверяем, что ID не пустые
    if req.OrderID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "order_id is required"})
        return
    }
    
    if req.UserID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
        return
    }
    
    // Проверяем существование заказа
    order, err := h.orderRepo.FindByID(req.OrderID)
    if err != nil {
        log.Printf("Order not found error: %v", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found with id: " + req.OrderID})
        return
    }
    
    log.Printf("Found order: ID=%s, CustomerID=%s", order.ID, order.CustomerID)
    
    // Проверяем аутентификацию пользователя
    authUserID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }
    
    log.Printf("Authenticated user ID: %s", authUserID.(string))
    
    if order.CustomerID != authUserID.(string) {
        log.Printf("Permission denied: order customer %s != authenticated user %s", 
            order.CustomerID, authUserID.(string))
        c.JSON(http.StatusForbidden, gin.H{"error": "Only order owner can assign executor"})
        return
    }
    
    if req.UserID == authUserID.(string) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot assign yourself as executor"})
        return
    }
    
    hasResponded, err := h.orderUserRepo.CheckOrderUserExists(c, req.OrderID, req.UserID)
    if err != nil {
        log.Printf("Error checking order user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check response"})
        return
    }
    
    if !hasResponded {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User has not responded to this order"})
        return
    }
    
    updatedOrder, err := h.orderUserRepo.AssignOrderExecutor(c, req.UserID, req.OrderID)
    if err != nil {
        log.Printf("AssignOrderExecutor error: %v", err)
        if err.Error() == "no rows in result set" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or cannot be updated"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign executor: " + err.Error()})
        }
        return
    }
    
    log.Printf("Successfully assigned executor. Updated order ID: %s", updatedOrder.ID.String())
    
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Executor assigned successfully",
        "order": gin.H{
            "id":                    updatedOrder.ID.String(),
            "customer_id":           updatedOrder.CustomerID.String(),
            "executor_id":           updatedOrder.ExecutorID.String(),
            "title":                 updatedOrder.Title,
            "description":           updatedOrder.Description,
            "budget":                getBudgetFloat(updatedOrder.Budget),
            "category":              updatedOrder.Category,
            "section":               updatedOrder.Section,
            "status":                updatedOrder.Status.String,
            "executor_assigned_at":  getNullableTime(updatedOrder.ExecutorAssignedAt),
            "created_at":            updatedOrder.CreatedAt.Time,
            "updated_at":            updatedOrder.UpdatedAt.Time,
        },
    })
}

func getNullableString(s pgtype.Text) *string {
    if s.Valid {
        return &s.String
    }
    return nil
}

func getNullableTime(t pgtype.Timestamptz) *time.Time {
    if t.Valid {
        return &t.Time
    }
    return nil
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