package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi semua routes API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware Logging
	r.Use(gin.Logger())

	// Route untuk login (public)
	r.POST("/login", controllers.LoginHandler)

	// Routes untuk Users (Public, tanpa proteksi)
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUsersByID)
		userRoutes.POST("", controllers.CreateUsers)
		userRoutes.PUT("/:id", controllers.UpdateUsers)
		userRoutes.DELETE("/:id", controllers.DeleteUsers)
	}

	// Routes untuk Admin (Protected)
	adminRoutes := r.Group("/admins")
	adminRoutes.Use(middlewares.AuthMiddleware()) // Melindungi route
	{
		adminRoutes.GET("", controllers.GetAllAdmins)
		adminRoutes.GET("/:id", controllers.GetAdminsByID)
		adminRoutes.POST("", controllers.CreateAdmins)
		adminRoutes.PUT("/:id", controllers.UpdateAdmins)
		adminRoutes.DELETE("/:id", controllers.DeleteAdmins)
	}

	// Routes untuk Aspirations (Protected)
	aspirationRoutes := r.Group("/aspirations")
	aspirationRoutes.Use(middlewares.AuthMiddleware()) // Melindungi route
	{
		aspirationRoutes.GET("", controllers.GetAllAspirations)
		aspirationRoutes.GET("/:id", controllers.GetAspirationByID)
		aspirationRoutes.POST("", controllers.CreateAspiration)
		aspirationRoutes.PUT("/:id", controllers.UpdateAspiration)
		aspirationRoutes.DELETE("/:id", controllers.DeleteAspiration)
	}

	return r
}
