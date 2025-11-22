package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"backend/internal/auth"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth.AuthService
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest

	body, _ := c.GetRawData()
	log.Printf("=== REGISTER REQUEST ===")
	log.Printf("Raw body: %s", string(body))
	log.Printf("Content-Type: %s", c.GetHeader("Content-Type"))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Parsed request: Username='%s', Email='%s', Password='%s'",
		req.Username, req.Email, req.Password)

	user, err := h.authService.Register(&req)
	if err != nil {
		log.Printf("Auth service error: %v", err)
		if err == auth.ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	log.Printf("Registration successful: UserID=%s", user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Login JSON bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Login attempt: Username='%s'", req.Username)

	token, user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		log.Printf("Login auth service error: %v", err)
		if err == auth.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not login"})
		return
	}

	log.Printf("Login successful: UserID=%s", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}