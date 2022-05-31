package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/routes"
	"github.com/gestaoFrota/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Role{})
	services.Db.AutoMigrate(&model.Users{})
	services.Db.AutoMigrate(&model.PassRecover{})
	services.Db.AutoMigrate(&model.Truck{})
	services.Db.AutoMigrate(&model.Trailer{})
	services.Db.AutoMigrate(&model.TrailerHistory{})
	services.Db.AutoMigrate(&model.TruckAndDriver{})
	services.Db.AutoMigrate(&model.Locations{})
	services.Db.AutoMigrate(&model.TruckHistory{})
	services.Db.AutoMigrate(&model.TruckState{})
	services.Db.AutoMigrate(&model.Displacements{})
	services.Db.AutoMigrate(&model.Messages{})
	services.Db.AutoMigrate(&model.TrailerState{})
	services.Db.AutoMigrate(&model.TravelMaps{})
	services.Db.AutoMigrate(&model.Supply{})
	services.Db.AutoMigrate(&model.DisplacementsAndSupply{})

	services.Db.Model(&model.Displacements{}).AddForeignKey("truck_id", "trucks(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.Displacements{}).AddForeignKey("trailer_id", "trailers(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TravelMaps{}).AddForeignKey("truck_id", "trucks(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TravelMaps{}).AddForeignKey("trailer_id", "trailers(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TrailerState{}).AddForeignKey("trailer_id", "trailers(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.Users{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TruckHistory{}).AddForeignKey("truck_id", "trucks(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TrailerHistory{}).AddForeignKey("trailer_id", "trucks(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TruckAndDriver{}).AddForeignKey("truck_id", "trucks(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TruckAndDriver{}).AddForeignKey("first_driver_id", "users(id)", "RESTRICT", "RESTRICT")
	services.Db.Model(&model.TruckState{}).AddForeignKey("truck_id", "trucks(id)", "RESTRICT", "RESTRICT")
	defer services.Db.Close()
}

func main() {

	services.FormatSwagger()
	fmt.Println("inicio")
	// Creates a gin router with default middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(services.GinMiddleware("*"))

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found1"})
	})

	//create a new router group called test
	test := router.Group("/api/v1")
	// use test group and create new route to create a new supply
	test.POST("/supply", routes.CreateSupply)
	test.GET("/stats", routes.GetStats)

	test.StaticFS("/file", http.Dir("public"))

	home := router.Group("/api")
	//home.Use(services.AuthorizationRequired())
	{
		// Users
		home.GET("/users", routes.GetAllUsers)
		home.GET("/usersDrivers", routes.GetAllTruckDrivers)
		home.DELETE("/deleteUser/:id", routes.DeleteUser)
		home.POST("/updateUser", routes.UpdateUser)
		home.POST("/register", routes.RegisterUser)
		home.GET("/travelMap/:id", routes.GetTravelMap)

		// Trucks
		//endpoint to submit a new supply

		home.GET("/trucks", routes.GetAllTrucks)
		home.GET("/trucksRpmPlates", routes.GetRpmPlates)
		home.POST("/trucks/register", routes.AddTrucks)
		home.POST("/trucks/bindTruckAndDriver", routes.BindTruckAndDriver)
		home.GET("trucks/getTruckDriver/:id", routes.GetTruckDriver)
		home.DELETE("trucks/unbindTruckDriver/:id", routes.UnbindTruckDriver)
		home.DELETE("trucks/deleteTruck/:id", routes.DeleteTrucks)
		home.GET("trucks/getTruck/:id", routes.GetTruck)
		home.POST("/trucks/updateTruck", routes.UpdateTrucks)
		home.POST("/trucks/addRoute", routes.AddRoute)
		home.GET("/trucks/alltruckState", routes.AllTruckState)
		home.GET("/trucks/getTruckHistory/:id", routes.GetTruckHistory)

		// Trailer
		home.GET("/trailers", routes.GetAllTrailers)
		home.POST("/trailers/register", routes.AddTrailers)
		home.DELETE("trailers/deleteTrailer/:id", routes.DeleteTrailers)
		home.POST("/trailers/updateTrailer", routes.UpdateTrailers)
		home.GET("trailers/getTrailer/:id", routes.GetTrailer)
		home.GET("/trucks/truckCount", routes.GetTruckCount)

		// Locations
		home.GET("/locations", routes.GetAllLocations)
		home.POST("/locations/addLocation", routes.AddLocations)
		home.DELETE("/locations/deleteLocation/:id", routes.DeleteLocation)
		home.GET("/locations/getLocation/:id", routes.GetLocation)
		home.POST("/locations/updateLocation", routes.UpdateLocation)

		// Displacement
		home.POST("/createDisplacement", routes.CreateDisplacement)

		// List trucks and Trailers
		home.GET("/listTrucksDrivers", routes.ListTrucksDrivers)
		home.GET("/listTruckAndDriver", routes.ListTruckAndDriver)
		home.GET("/listTrailers", routes.ListTrailers)

		// Messages
		home.POST("/messages", routes.Messages)
		home.GET("/getMessages", routes.GetMessages)
		home.GET("/getMessagesByUser/:id", routes.GetMessagesByUser)

		// PROJETO SISTEMA REDES
		home.GET("/bestGasStation/:id/:r/:n", routes.BestGasStation)
		home.GET("/rWithoutTrafficCongestion/:id", routes.RWithoutTrafficCongestion)
	}

	mobileApi := router.Group("/api/mobile")
	{
		mobileApi.GET("/getDriverRoutes/:email", routes.GetDriverRoute)
		mobileApi.GET("/finishedRoute/:id", routes.FinishedRoute)
		mobileApi.POST("/mobileLogin", routes.MobileGenerateToken)
	}

	auth := router.Group("/api/auth")
	{
		auth.Use(cors.Default())
		auth.POST("/trucks/truckState", routes.TruckState)
		auth.POST("/trailer/trailerState", routes.TrailerState)
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/passRecover", routes.VerififyEmail)
		auth.POST("/passRecover/checktoken", routes.CheckToken)
		auth.POST("/passRecover/setPassword", routes.SetPassword)
		auth.POST("/changePassword", routes.ChangePassword)
		auth.PUT("/refresh_token", services.AuthorizationRequired(), routes.RefreshToken)
	}
	port := os.Getenv("PORT")
	//router.Run(":80")
	router.Run(":" + port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
