package controllers

import (
	
	"net/http"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

type teste struct {
		
	DriverId   int  ` json:"driver_id"`
	TruckId   int ` json:"truck_id"`
	
}
func BindTruckAndDriver(c *gin.Context){
	var trucksDuplicated model.Truck
	var truck,truckaux,find model.TruckAndDriver
	var user model.Users;
	var teste teste
		//bind 
		if err := c.ShouldBindJSON(&teste); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
			return
		}
		
		//Primeiro verificamos se o camiao existe
		services.OpenDatabase()
		services.Db.Find(&trucksDuplicated, "id = ?", teste.TruckId)
		defer services.Db.Close()
		//Senao existir retornamos que o ID do camiao foi invalido
		if trucksDuplicated.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao não encontrado!"})
			return
		}
		//Verificamos se o camionista enviado existe
		services.OpenDatabase()
		services.Db.Find(&user, "id = ?", teste.DriverId)
		defer services.Db.Close()
		//Senao existir retornamos que o ID do camionista foi invalido
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camionista não encontrado!"})
			return
		}
		services.OpenDatabase()
		services.Db.Find(&find, "first_driver_id = ?", teste.DriverId)
		if(find.FirstDriverId!=0){
			defer services.Db.Close()
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camionista já possui camião associado!"})
			return
		}
		services.OpenDatabase()
		services.Db.Find(&find, "second_driver_id = ?", teste.DriverId)
		defer services.Db.Close()
		if(find.SecondDriverId!=0){
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camionista já possui camião associado!"})
			return
		}

		//Verificamos se o camiao já  existe na base dados de estados e caso exista, se já tem algum camionista associado
		
			truck.TruckId = teste.TruckId
			
			truck.FirstDriverId = teste.DriverId
			truck.SecondDriverId = 0
			services.OpenDatabase()
		services.Db.Find(&truckaux, "truck_id = ?", teste.TruckId)
		//Senao existir adicionamos à tabela
		if truckaux.TruckId == 0 {
			
			services.Db.Save(&truck)
		
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Camionista adicionado com sucesso!"})
			return
		}
		defer services.Db.Close()

		//Se existir , entao verificamos se já tem condutor principal , se tiver verificamos se tem secundario
		services.OpenDatabase()
		services.Db.Find(&truckaux, "truck_id = ?", teste.TruckId)
		if truckaux.FirstDriverId != 0 {
			if truckaux.SecondDriverId == 0  {
			
			services.Db.Exec("update truck_and_drivers set second_driver_id = ? where truck_id = ?", teste.DriverId,teste.TruckId)
			
		
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Camionista adicionado com sucesso!"})
			return
			}else {
				defer services.Db.Close()
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao com maximo motoristas!"})
			return
			}	
		}
		
		defer services.Db.Close()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request"})	
}

func AddRoute(c *gin.Context){ // SERVE APENAS PARA TESTAR PQ "TELETRANSPORTA" O CAMIAO PARA ONDE QUEREMOS
	var routes struct {
		Latitude   float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		TruckId uint `json:"truck_id"`
	}

	var truck model.Truck
	var truckHistory model.TruckHistory
	var truckState model.TruckState

	if err := c.ShouldBindJSON(&routes); err != nil {
			
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.OpenDatabase()
	// PESQUISA O CAMIAO E ATUALIZA A SUA POSICAO ESTE MÉTODO É APENAS PARA TESTAR
	services.Db.Find(&truck, "id = ?", routes.TruckId)
	services.Db.Exec("update trucks set latitude = ?, longitude = ? where id = ?", routes.Latitude, routes.Longitude ,truck.ID)

	// Atualizo no Histórico a nova posição do camião
	truckHistory.TruckId = truck.ID
	truckHistory.Longitude = routes.Longitude
	truckHistory.Latitude = routes.Latitude
	truckHistory.Rpm = 0
	truckHistory.Speed = 0
	services.Db.Save(&truckHistory)

	truckState.Latitude = routes.Latitude
	truckState.Longitude = routes.Longitude
	truckState.TruckId = truck.ID
	truckState.Rpm = 0
	truckState.Speed = 0
	services.Db.Exec("update truck_states set latitude = ?, longitude = ?, rpm = ?, speed = ? where truck_id = ?", truckState.Latitude, truckState.Longitude, truckState.Rpm,truckState.Speed,truckState.TruckId)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
}
type DriverData struct {
	
	TruckPlate   string ` json:"truck_plate"`
	TruckFirstDriver   string ` json:"truck_first"`
	TruckFirstDriverId   int ` json:"truck_firstId"`
	TruckSecondDriver   string ` json:"truck_second"`
	TruckSecondDriverId   int ` json:"truck_secondId"`
}
func GetTruckDriver(c *gin.Context){
	var truck model.TruckAndDriver
	var driverdata DriverData
	var truckdata model.Truck
	var userData,secondData model.Users
	id := c.Param("id")
	
   	services.OpenDatabase()
	services.Db.Find(&truck, "truck_id = ?", id)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&truckdata, "id = ?", id)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&userData, "id = ?", truck.FirstDriverId)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&secondData, "id = ?", truck.SecondDriverId)
	defer services.Db.Close()

	driverdata.TruckPlate = truckdata.Plate
	driverdata.TruckFirstDriver = userData.FirstName
	driverdata.TruckFirstDriverId = int(userData.ID)
	driverdata.TruckSecondDriver = secondData.FirstName
	driverdata.TruckSecondDriverId = int(secondData.ID)
	
	///falta fazer os casos em que o find pode nao dar certo
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": driverdata})
}

func UnbindTruckDriver(c *gin.Context){
	var truckDriver model.TruckAndDriver
	//Recebemos o id do camionista a retirar e vamos procurar a que camiao está associado
	id := c.Param("id")
	
	services.OpenDatabase()
	 services.Db.Find(&truckDriver, "first_driver_id = ?", id)

	//Se o nosso camionista for o primeiro condutor 
	
	
	if(truckDriver.FirstDriverId != 0) {
		//caso nao haja segundo condutor , removemos o camiao pois um camiao tem que ter sempre 1 condutor
		if(truckDriver.SecondDriverId == 0){
			//Falta adicionar À tabela historico quando removemos , para guardarmos os dados acerca dos motoristas que ja tiveram associados a um camiao
			services.Db.Exec("delete from truck_and_drivers where first_driver_id = ?",truckDriver.FirstDriverId)
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Camionista desassociado!"})
			return
		}else if (truckDriver.SecondDriverId != 0){
			// entao passamos o segundo condutor a primeiro depois de remover.
			//Falta adicionar À tabela historico quando removemos , para guardarmos os dados acerca dos motoristas que ja tiveram associados a um camiao
			services.Db.Exec("update truck_and_drivers set first_driver_id = second_driver_id , second_driver_id = 0 where first_driver_id = ?",truckDriver.FirstDriverId)
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Camionista desassociado!"})
			return
		}
	}else {
		services.Db.Find(&truckDriver, "second_driver_id = ?", id)
		if(truckDriver.SecondDriverId !=0){
			services.Db.Exec("update truck_and_drivers set second_driver_id = 0 where second_driver_id = ?",truckDriver.SecondDriverId)
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Camionista desassociado!"})
			return
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camionista nao se encontra associado a nenhum camião!"})
			defer services.Db.Close()
			return
		}

	}
	
	//Se for o segundo condutor entao simplesmente colocamos a 0
}

