package routes

import (
	"api-web-pemerintah/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi semua routes API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware Logging dan Recovery
	r.Use(gin.Logger(), gin.Recovery())

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Sesuaikan dengan domain frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Tambahkan OPTIONS
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Endpoint untuk login menggunakan LoginUserController
	r.POST("/login", controllers.LoginUserController)

	// Rute untuk register user
	r.POST("/register", controllers.RegisterUser)

	// Routes untuk Users
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUsersByID)
		userRoutes.POST("", controllers.CreateUsers)
		userRoutes.PUT("/:id", controllers.UpdateUsers)
		userRoutes.DELETE("/:id", controllers.DeleteUsers)
	}

	// Routes untuk Admin
	adminRoutes := r.Group("/admins")
	{
		adminRoutes.GET("", controllers.GetAllAdmins)
		adminRoutes.GET("/:id", controllers.GetAdminsByID)
		adminRoutes.POST("", controllers.CreateAdmins)
		adminRoutes.PUT("/:id", controllers.UpdateAdmins)
		adminRoutes.DELETE("/:id", controllers.DeleteAdmins)
		adminRoutes.POST("/login", controllers.LoginAdminController)
	}

	// Routes untuk Aspirations
	aspirationRoutes := r.Group("/aspirations")
	{

		// aspirationRoutes.POST("", controllers.CreateAspiration)
		aspirationRoutes.POST("", controllers.CreateAspiration2)
		aspirationRoutes.GET("", controllers.GetAllAspirations)
		aspirationRoutes.GET("/:id", controllers.GetAspirationByID)
		aspirationRoutes.GET("/user/:user_id", controllers.GetAspirationsByUserID)
		// aspirationRoutes.GET("/count/:id_pengirim", controllers.GetAspirationsCountByUserID)
		aspirationRoutes.PUT("/:id", controllers.UpdateAspiration)
		aspirationRoutes.DELETE("/:id", controllers.DeleteAspiration)
		// aspirationRoutes.GET("/user/:user_id/week", controllers.GetAspirationsCountByUserIDLast7Days)
		// aspirationRoutes.GET("/user/:user_id/week", controllers.CreateAspiration2)

	}

	return r
}
