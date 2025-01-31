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

	// Routes untuk News
	newsRoutes := r.Group("/news")
	{
		newsRoutes.GET("", controllers.GetAllNews)
		newsRoutes.GET("/:id", controllers.GetNewsByID)
		newsRoutes.POST("", controllers.CreateNews)
		newsRoutes.PUT("/:id", controllers.UpdateNews)
		newsRoutes.DELETE("/:id", controllers.DeleteNews)
	}

	// Routes untuk Informations
	infoRoutes := r.Group("/informations")
	{
		infoRoutes.GET("", controllers.GetAllInformations)
		infoRoutes.GET("/:id", controllers.GetInformationByID)
		infoRoutes.POST("", controllers.CreateInformation)
		infoRoutes.PUT("/:id", controllers.UpdateInformation)
		infoRoutes.DELETE("/:id", controllers.DeleteInformation)
	}

	// Routes untuk Subsidies
	subsidyRoutes := r.Group("/subsidies")
	{
		subsidyRoutes.GET("", controllers.GetAllSubsidies)
		subsidyRoutes.GET("/:id", controllers.GetSubsidyByID)
		subsidyRoutes.POST("", controllers.CreateSubsidy)
		subsidyRoutes.PUT("/:id", controllers.UpdateSubsidy)
		subsidyRoutes.DELETE("/:id", controllers.DeleteSubsidy)
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

	// Routes untuk Destinations
	destinationRoutes := r.Group("/destinations")
	{
		destinationRoutes.GET("", controllers.GetAllDestinations)
		destinationRoutes.GET("/:id", controllers.GetDestinationByID)
		destinationRoutes.POST("", controllers.CreateDestination)
		destinationRoutes.PUT("/:id", controllers.UpdateDestination)
		destinationRoutes.DELETE("/:id", controllers.DeleteDestination)
	}

	// Routes untuk Testimonies
	testimonyRoutes := r.Group("/testimonies")
	{
		testimonyRoutes.GET("", controllers.GetAllTestimonies)
		testimonyRoutes.GET("/:id", controllers.GetTestimonyByID)
		testimonyRoutes.POST("", controllers.CreateTestimony)
		testimonyRoutes.PUT("/:id", controllers.UpdateTestimony)
		testimonyRoutes.DELETE("/:id", controllers.DeleteTestimony)
	}

	// Routes untuk Packages
	packageRoutes := r.Group("/packages")
	{
		packageRoutes.GET("", controllers.GetAllPackages)
		packageRoutes.GET("/:id", controllers.GetPackageByID)
		packageRoutes.POST("", controllers.CreatePackage)
		packageRoutes.PUT("/:id", controllers.UpdatePackage)
		packageRoutes.DELETE("/:id", controllers.DeletePackage)
	}

	// Routes untuk Visitors
	visitorRoutes := r.Group("/visitors")
	{
		visitorRoutes.GET("", controllers.GetAllVisitors)
		visitorRoutes.GET("/:id", controllers.GetVisitorByID)
		visitorRoutes.POST("", controllers.CreateVisitor)
		visitorRoutes.PUT("/:id", controllers.UpdateVisitor)
		visitorRoutes.DELETE("/:id", controllers.DeleteVisitor)
	}

	// Routes untuk Orders
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("", controllers.GetAllOrders)
		orderRoutes.GET("/:id", controllers.GetOrderByID)
		orderRoutes.POST("", controllers.CreateOrder)
		orderRoutes.PUT("/:id", controllers.UpdateOrder)
		orderRoutes.DELETE("/:id", controllers.DeleteOrder)
	}

	return r
}
