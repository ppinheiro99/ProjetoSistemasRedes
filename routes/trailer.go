package routes

import (
	"github.com/gestaoFrota/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllTrailers(c *gin.Context) {
	controllers.GetAllTrailers(c)
}
func AddTrailers(c *gin.Context) {
	controllers.AddTrailers(c)
}

func DeleteTrailers(c *gin.Context) {
	controllers.DeleteTrailers(c)
}

func GetTrailer(c *gin.Context) {
	controllers.GetTrailer(c)
}

func UpdateTrailers(c *gin.Context) {
	controllers.UpdateTrailer(c)
}
func TrailerState(c *gin.Context) {
	controllers.TrailerState(c)
}