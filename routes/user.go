package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	controllers.GetAllUsers(c)
}

func GetAllTruckDrivers(c *gin.Context) {
	controllers.GetAllTruckDrivers(c)
}

func DeleteUser(c *gin.Context) {
	controllers.DeleteUser(c)
}

func UpdateUser(c *gin.Context){
	controllers.UpdateUser(c)
}