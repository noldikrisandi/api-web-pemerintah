package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi semua routes API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware Logging
	r.Use(gin.Logger())

	// Routes untuk Users
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUsersByID)
		userRoutes.POST("", controllers.CreateUsers)
		userRoutes.PUT("/:id", controllers.UpdateUsers)
		userRoutes.DELETE("/:id", controllers.DeleteUsers)
	}

	// Routes untuk Aspirations
	aspirationRoutes := r.Group("/aspirations")
	{
		aspirationRoutes.GET("", controllers.GetAllAspirations)
		aspirationRoutes.GET("/:id", controllers.GetAspirationByID)
		aspirationRoutes.POST("", controllers.CreateAspiration)
		aspirationRoutes.PUT("/:id", controllers.UpdateAspiration)
		aspirationRoutes.DELETE("/:id", controllers.DeleteAspiration)
	}

	return r
}
