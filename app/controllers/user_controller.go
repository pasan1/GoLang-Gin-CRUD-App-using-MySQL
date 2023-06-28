package controllers

import (
	"net/http"

	"GoCRUDApplicationMySQL/app/models"
	"GoCRUDApplicationMySQL/app/repositories"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.repo.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Implement other CRUD operations like GetUser, UpdateUser, and DeleteUser
