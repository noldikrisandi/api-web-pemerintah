package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define Routes untuk Users
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUsersByID)
	r.POST("/users", controllers.CreateUsers)
	r.PUT("/users/:id", controllers.UpdateUsers)
	r.DELETE("/users/:id", controllers.DeleteUsers)

	// Define Routes untuk News
	r.GET("/news", controllers.GetAllNews)
	r.GET("/news/:id", controllers.GetNewsByID)
	r.POST("/news", controllers.CreateNews)
	r.PUT("/news/:id", controllers.UpdateNews)
	r.DELETE("/news/:id", controllers.DeleteNews)

	// Define Routes untuk Informations
	r.GET("/informations", controllers.GetAllInformations)
	r.GET("/informations/:id", controllers.GetInformationByID)
	r.POST("/informations", controllers.CreateInformation)
	r.PUT("/informations/:id", controllers.UpdateInformation)
	r.DELETE("/informations/:id", controllers.DeleteInformation)

	// Define Routes untuk Subsidies
r.GET("/subsidies", controllers.GetAllSubsidies)
r.GET("/subsidies/:id", controllers.GetSubsidyByID)
r.POST("/subsidies", controllers.CreateSubsidy)
r.PUT("/subsidies/:id", controllers.UpdateSubsidy)
r.DELETE("/subsidies/:id", controllers.DeleteSubsidy)

// Define Routes untuk Aspirations
r.GET("/aspirations", controllers.GetAllAspirations)
r.GET("/aspirations/:id", controllers.GetAspirationByID)
r.POST("/aspirations", controllers.CreateAspiration)
r.PUT("/aspirations/:id", controllers.UpdateAspiration)
r.DELETE("/aspirations/:id", controllers.DeleteAspiration)

// Define Routes untuk Destinations
r.GET("/destinations", controllers.GetAllDestinations)
r.GET("/destinations/:id", controllers.GetDestinationByID)
r.POST("/destinations", controllers.CreateDestination)
r.PUT("/destinations/:id", controllers.UpdateDestination)
r.DELETE("/destinations/:id", controllers.DeleteDestination)

// Define Routes untuk Testimonies
r.GET("/testimonies", controllers.GetAllTestimonies)
r.GET("/testimonies/:id", controllers.GetTestimonyByID)
r.POST("/testimonies", controllers.CreateTestimony)
r.PUT("/testimonies/:id", controllers.UpdateTestimony)
r.DELETE("/testimonies/:id", controllers.DeleteTestimony)

// Define Routes untuk Packages
r.GET("/packages", controllers.GetAllPackages)
r.GET("/packages/:id", controllers.GetPackageByID)
r.POST("/packages", controllers.CreatePackage)
r.PUT("/packages/:id", controllers.UpdatePackage)
r.DELETE("/packages/:id", controllers.DeletePackage)

// Define Routes untuk Visitors
r.GET("/visitors", controllers.GetAllVisitors)
r.GET("/visitors/:id", controllers.GetVisitorByID)
r.POST("/visitors", controllers.CreateVisitor)
r.PUT("/visitors/:id", controllers.UpdateVisitor)
r.DELETE("/visitors/:id", controllers.DeleteVisitor)

// Define Routes untuk Orders
r.GET("/orders", controllers.GetAllOrders)
r.GET("/orders/:id", controllers.GetOrderByID)
r.POST("/orders", controllers.CreateOrder)
r.PUT("/orders/:id", controllers.UpdateOrder)
r.DELETE("/orders/:id", controllers.DeleteOrder)





	return r
}
