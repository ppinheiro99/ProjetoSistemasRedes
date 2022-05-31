package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	controllers.GetAllLocations(c)
}

func GetLocation(c *gin.Context) {
	controllers.GetLocation(c)
}

func AddLocations(c *gin.Context) {
	controllers.AddLocations(c)
}

func DeleteLocation(c *gin.Context) {
	controllers.DeleteLocation(c)
}

func UpdateLocation(c *gin.Context) {
	controllers.UpdateLocation(c)
}

func RWithoutTrafficCongestion(c *gin.Context) {
	controllers.RWithoutTrafficCongestion(c)
}

func BestGasStation(c *gin.Context) {
	controllers.BestGasStation(c)
}
