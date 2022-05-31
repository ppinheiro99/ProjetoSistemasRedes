package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func Messages(c *gin.Context){
	controllers.Messages(c)
}

func GetMessagesByUser(c *gin.Context){
	controllers.GetMessagesByUser(c)
}

func GetMessages(c *gin.Context){
	controllers.GetMessages(c)
}