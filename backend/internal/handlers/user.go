package handlers

import (
	"net/http"

	"backend/internal/models"
	"backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := h.userRepo.FindByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
		return
	}

	var publicUsers []gin.H
	for _, user := range users {
		publicUsers = append(publicUsers, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": publicUsers,
		"count": len(publicUsers),
	})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser, err := h.userRepo.FindByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if req.Username != currentUser.Username {
		if existingUser, _ := h.userRepo.FindByUsername(req.Username); existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}
	}

	if req.Email != currentUser.Email {
		if existingUser, _ := h.userRepo.FindByEmail(req.Email); existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
			return
		}
	}

	currentUser.Username = req.Username
	currentUser.Email = req.Email

	if err := h.userRepo.Update(currentUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"id":       currentUser.ID,
			"username": currentUser.Username,
			"email":    currentUser.Email,
		},
	})
}