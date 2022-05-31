package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllTrucks(c *gin.Context) {
	controllers.GetAllTrucks(c)
}
func AddTrucks(c *gin.Context) {
	controllers.AddTrucks(c)
}

func DeleteTrucks(c *gin.Context) {
	controllers.DeleteTrucks(c)
}

func GetTruck(c *gin.Context) {
	controllers.GetTruck(c)
}

func UpdateTrucks(c *gin.Context) {
	controllers.UpdateTruck(c)
}

func BindTruckAndDriver(c *gin.Context) {
	controllers.BindTruckAndDriver(c)
}

func AddRoute(c *gin.Context) {
	controllers.AddRoute(c)
}
func GetTruckDriver(c *gin.Context) {
	controllers.GetTruckDriver(c)
}
func UnbindTruckDriver(c *gin.Context) {
	controllers.UnbindTruckDriver(c)
}
func TruckState(c *gin.Context) {
	controllers.TruckState(c)
}
func GetTruckCount(c *gin.Context) {
	controllers.GetTruckCount(c)
}
func GetRpmPlates(c *gin.Context) {
	controllers.GetRpmPlates(c)
}

func AllTruckState(c *gin.Context) {
	controllers.AllTruckState(c)
}

func GetTruckHistory(c *gin.Context) {
	controllers.GetTruckHistory(c)
}
