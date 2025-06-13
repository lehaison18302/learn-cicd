package users

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"my_project/internal/actionlog"
	"my_project/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetUsersHandler handles GET /users with pagination
func (h *Handler) GetUsersHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	ctx := context.Background()
	users, err := h.service.GetUsersWithPagination(ctx, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUserHandler handles POST /users to create a new user
func (h *Handler) CreateUserHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := context.Background()
	if err := h.service.CreateNewUser(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	actionlog.Log(actionlog.ActionLog{
		UserID:   c.GetInt("user_id"),
		Username: c.GetString("username"),
		Action:   "create_user",
		TargetID: user.ID,
		Time:     time.Now(),
	})
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// UpdateUserHandler handles PUT /users/:id to update a user
func (h *Handler) UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user.ID = id
	ctx := context.Background()
	if err := h.service.UpdateUser(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	actionlog.Log(actionlog.ActionLog{
		UserID:   c.GetInt("user_id"),
		Username: c.GetString("username"),
		Action:   "update_user",
		TargetID: user.ID,
		//		Time:     time.Now(),
	})
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUserHandler handles DELETE /users/:id to delete a user
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}
	ctx := context.Background()
	if err := h.service.DeleteUser(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	actionlog.Log(actionlog.ActionLog{
		UserID:   c.GetInt("user_id"),
		Username: c.GetString("username"),
		Action:   "delete_user",
		TargetID: id,
		Time:     time.Now(),
	})
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// LoginHandler handles POST /login
func (h *Handler) LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := context.Background()
	user, err := h.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := middleware.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "username": user.Username})
}
