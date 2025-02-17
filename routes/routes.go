package routes

import (
	"api-web-pemerintah/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // untuk ases frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/login", controllers.LoginUserController)

	r.POST("/register", controllers.RegisterUser)

	// Routes untuk Users
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUsersByID)
		userRoutes.POST("", controllers.CreateUsers)
		userRoutes.PUT("/:id", controllers.UpdateUsers)
		// userRoutes.DELETE("/:id", controllers.DeleteUsers) saya tidak beri akses hapus
	}

	// Routes untuk Admin
	adminRoutes := r.Group("/admins")
	{
		adminRoutes.GET("", controllers.GetAllAdmins)
		adminRoutes.GET("/:id", controllers.GetAdminsByID)
		adminRoutes.POST("", controllers.CreateAdmins)
		adminRoutes.PUT("/:id", controllers.UpdateAdmins)
		// adminRoutes.DELETE("/:id", controllers.DeleteAdmins) saya tidak beri akses hapus
		adminRoutes.POST("/login", controllers.LoginAdminController)
	}

	// Routes untuk Aspirations
	aspirationRoutes := r.Group("/aspirations")
	{

		aspirationRoutes.POST("", controllers.CreateAspiration)
		aspirationRoutes.GET("", controllers.GetAllAspirations)
		aspirationRoutes.GET("/:id", controllers.GetAspirationByID)
		aspirationRoutes.GET("/user/:user_id", controllers.GetAspirationsByUserID)
		aspirationRoutes.PUT("/:id", controllers.UpdateAspiration)
		// aspirationRoutes.DELETE("/:id", controllers.DeleteAspiration) saya tidak beri akses hapus

	}

	return r
}
