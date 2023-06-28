package routes

import (
	"database/sql"

	"GoCRUDApplicationMySQL/app/controllers"
	"GoCRUDApplicationMySQL/app/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUsers)
		v1.POST("/users", userController.CreateUser)
		// Add other routes for GetUser, UpdateUser, and DeleteUser
	}

	return r
}
