package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func GetDriverRoute(c *gin.Context){
	controllers.GetDriverRoute(c)
}

func FinishedRoute(c *gin.Context){
	controllers.FinishedRoute(c)
}