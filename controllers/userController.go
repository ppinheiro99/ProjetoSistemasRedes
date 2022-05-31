package controllers

import (
	"fmt"
	"net/http"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []model.Users
	services.OpenDatabase()
	services.Db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		defer services.Db.Close()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
	defer services.Db.Close()
}

func DeleteUser(c *gin.Context) {
	var user model.Users
	services.OpenDatabase()
	id := c.Param("id")

	services.Db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	// Nao permite que elimine cargos superiores aos seus AINDA POR TERMINAR !!!!!!!!!!!!!!!!!!!

	if user.RoleId == 1{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Nao é possivel eliminar o super admin!"})
		return
	}

	services.Db.Exec("delete from users where email = ?",user.Email)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func UpdateUser(c *gin.Context){
	var creds model.Users

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(creds.FirstName) < 3 || len(creds.FirstName) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Nome!"})
		return
	}

	if len(creds.Country) < 2 || len(creds.Country) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Pais!"})
		return
	}

	if len(creds.LastName) < 3 || len(creds.LastName) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Apelido!"})
		return
	}

	services.OpenDatabase()
	services.Db.Exec("update users set first_name = ?, last_name = ?, address = ?, country = ? where email = ?", creds.FirstName,creds.LastName, creds.Address, creds.Country, creds.Email)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "User ID": creds.ID, "user": creds})

}

func GetAllTruckDrivers(c *gin.Context){
	var users []model.Users
	services.OpenDatabase()
	services.Db.Find(&users, "role_id = 4")
	defer services.Db.Close()
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func GetTravelMap(c *gin.Context){
	// Vou buscar o Truck à BD
	// de seguida vou à BD de travel maps e retorno os mapas que possuem o camiao e retorno o mapa
	var travelMaps []model.TravelMaps
	var truck model.Truck
	id := c.Param("id")
	fmt.Println(id)
	services.OpenDatabase()
	services.Db.First(&truck, id)
	defer services.Db.Close()
	fmt.Println(truck)
	if truck.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	fmt.Println(truck.ID)
	services.OpenDatabase()
	services.Db.Find(&travelMaps, "truck_id = ?", truck.ID)
	defer services.Db.Close()
	fmt.Println(travelMaps)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": travelMaps})

}