package controllers

import (
	"fmt"
	"net/http"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

func GetDriverRoute(c *gin.Context) {
	var usr model.Users
	var truckAndDriver model.TruckAndDriver
	email := c.Param("email")
	services.OpenDatabase()
	services.Db.Find(&usr, "email = ?", email)
	defer services.Db.Close()
	if usr.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		defer services.Db.Close()
		return
	}
	services.OpenDatabase()
	services.Db.Find(&truckAndDriver, "first_driver_id = ?", usr.ID)
	if truckAndDriver.ID == 0 { // caso o camionista que faz a pesquisa seja um 2 motorista e nao 1
		services.Db.Find(&truckAndDriver, "second_driver_id = ?", usr.ID)
		searchRoute(truckAndDriver, c)
		defer services.Db.Close()
		return
	}
	searchRoute(truckAndDriver, c)
	defer services.Db.Close()
}

func searchRoute(truckAndDriver model.TruckAndDriver, c *gin.Context) {
	var routes model.Displacements
	services.OpenDatabase()
	services.Db.Find(&routes, "truck_id = ?", truckAndDriver.TruckId)
	if routes.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "User has no route!"})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": routes.TruckId, "coords": routes.Coords})
	defer services.Db.Close()

}

func FinishedRoute(c *gin.Context) { // api envia o id da rota que o camionista executou e passamos essa rota para o mapa de viagem do ou dos camionistas e eliminamos da bd de rotas
	var travelMap model.TravelMaps
	var displacements model.Displacements
	var suppliesIds[] model.DisplacementsAndSupply
	var currentSupply model.Supply
	var TotalValue, TotalLiters float64
	id := c.Param("id")
	

	services.OpenDatabase()
	services.Db.Find(&displacements, "truck_id = ?", id)
	if displacements.Coords == "" { // os deslocamentos tem obrigatoriamente de ter coordenadas. se nao tiver entao nao existe
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Displacement!"})
		return
	}
	//
	fmt.Println(displacements.ID)

	services.Db.Find(&suppliesIds, "displacement_id = ?", displacements.ID)

	fmt.Println(suppliesIds)
	// percorrer lista de suppliesIds e somar o total de litros e o valor total
	for _, supplyId := range suppliesIds {
		services.Db.Find(&currentSupply, "id = ?", supplyId.SupplyId)
		TotalValue += currentSupply.Value
		TotalLiters += currentSupply.Liters
	}
		///services.Db.Find(&currentSupply, "id = ?", suppliesIds.SupplyId)

		//TotalValue += currentSupply.Value
		fmt.Println("total value: ", TotalValue)
		//TotalLiters += currentSupply.Liters
		fmt.Println("total liters: ", TotalLiters)
		
	
	// Criamos um Mapa de Viagens com o deslocamento terminado pelo camionista
	travelMap.Coords = displacements.Coords
	travelMap.TruckId = displacements.TruckId
	travelMap.TrailerId = displacements.TrailerId
	travelMap.Distance = displacements.Distance
	travelMap.StartAddress = displacements.StartAddress
	travelMap.StartCity = displacements.StartCity
	travelMap.StartCountry = displacements.StartCountry
	travelMap.StartPostalCode = displacements.StartPostalCode
	travelMap.EndAddress = displacements.EndAddress
	travelMap.EndCity = displacements.EndCity
	travelMap.EndCountry = displacements.EndCountry
	travelMap.EndPostalCode = displacements.EndPostalCode
	travelMap.Time = displacements.Time
	travelMap.TotalValue = TotalValue
	travelMap.TotalLiters = TotalLiters
	services.Db.Save(&travelMap)

	// Eliminar da BD de deslocamentos
	services.Db.Exec("delete from displacements where id = ?", displacements.ID)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Succeeded!"})
}
