package handlers

import (
	"net/http"
	"log"
	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepo repository.OrderRepository
}

func NewOrderHandler(orderRepo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo: orderRepo,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Creating order for user %s: %+v", userID.(string), req)

	order := &models.Order{
		CustomerID:  userID.(string),
		Title:       req.Title,
		Description: req.Description,
		Budget:      req.Budget,
		Category:    req.Category,
		Section:     req.Section,
		Status:      "active",
	}

	if err := h.orderRepo.Create(order); err != nil {
		log.Printf("Order creation error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	log.Printf("Order created successfully: %s", order.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orders, err := h.orderRepo.FindByCustomer(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"count":  len(orders),
	})
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"count":  len(orders),
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := h.orderRepo.FindByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orderID := c.Param("id")

	var req models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	order.Title = req.Title
	order.Description = req.Description
	order.Budget = req.Budget
	order.Category = req.Category
	order.Section = req.Section
	order.Status = req.Status

	if err := h.orderRepo.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order":   order,
	})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orderID := c.Param("id")

	if err := h.orderRepo.Delete(orderID, userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}