package routes

import (

	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

//func to create new supply
func CreateSupply(c *gin.Context) {
	controllers.CreateSupply(c)
}
//func to create get stats
func GetStats(c *gin.Context) {
	controllers.GetStats(c)
}

func CreateDisplacement(c *gin.Context) {
	controllers.CreateDisplacement(c)
}

func ListTrucksDrivers(c *gin.Context) {
	controllers.ListTrucksDrivers(c)
}

func ListTrailers(c *gin.Context) {
	controllers.ListTrailers(c)
}

func ListTruckAndDriver(c *gin.Context) {
	controllers.ListTruckAndDriver(c)
}
