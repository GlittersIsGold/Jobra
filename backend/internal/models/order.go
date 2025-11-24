package models

import "time"

type Order struct {
    ID          string    `json:"id"`
    CustomerID  string    `json:"customer_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Budget      float64   `json:"budget"`
    Category    string    `json:"category"`
    Section     string    `json:"section"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
    Title       string  `json:"title" binding:"required,min=3"`
    Description string  `json:"description" binding:"required,min=10"`
    Budget      float64 `json:"budget" binding:"required,gt=0"`
    Category    string  `json:"category" binding:"required"`
    Section     string  `json:"section" binding:"required"`
}

type UpdateOrderRequest struct {
    Title       string  `json:"title" binding:"required,min=3"`
    Description string  `json:"description" binding:"required,min=10"`
    Budget      float64 `json:"budget" binding:"required,gt=0"`
    Category    string  `json:"category" binding:"required"`
    Section     string  `json:"section" binding:"required"`
    Status      string  `json:"status" binding:"required"`
}