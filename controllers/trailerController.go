package controllers

import (
	"net/http"
	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
)

func GetAllTrailers(c *gin.Context) {
	var trailers []model.Trailer
   	services.OpenDatabase()
	services.Db.Find(&trailers)
	defer services.Db.Close()
	if len(trailers) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trailers})

}

func GetTrailer(c *gin.Context) {
	var trailer model.Trailer

	id := c.Param("id")
	
   	services.OpenDatabase()
	services.Db.Find(&trailer,id)
	defer services.Db.Close()
	if len(trailer.Plate) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Reboque não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trailer})

}

/// CORRIGIR A LATITUDE E LONGITUDE DO CAMIAO, PARA QUANDO CAMIAO É CRIADO SER IGUAL À LOCALIZAÇÃO DA EMPRESA
func AddTrailers(c *gin.Context) {
	var trailers model.Trailer
	var trailersDuplicated model.Trailer
	var trailerHistory model.TrailerHistory

		if err := c.ShouldBindJSON(&trailers); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
			return
		}
		
		if len(trailers.Plate) > 9 || len(trailers.Plate) < 5 {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula Invalida!"})
			return
		}

		if trailers.Year <= 1960 || trailers.Year >= 2040  {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Ano Invalido!"})
			return
		}
		
		services.OpenDatabase()
		services.Db.Find(&trailersDuplicated, "plate = ?", trailers.Plate)
		if len(trailersDuplicated.Plate) > 0 {
			defer services.Db.Close()
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula já existe!"})
			return
		}
		
		trailers.Latitude = latitudeUfp
		trailers.Longitude = longitudeUfp


		services.Db.Save(&trailers)

		trailerHistory.Latitude = trailers.Latitude
		trailerHistory.Longitude = trailers.Longitude
		trailerHistory.TrailerId = trailers.ID

		services.Db.Save(&trailerHistory)
		
		defer services.Db.Close()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})

}

func DeleteTrailers(c *gin.Context) {
	var trailer model.Trailer
	services.OpenDatabase()
	id := c.Param("id")

	services.Db.First(&trailer, id)

	if trailer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
		return
	}

	services.Db.Exec("delete from trailers where id = ?",trailer.ID)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func UpdateTrailer(c *gin.Context) {
	var trailers model.Trailer
	
	if err := c.ShouldBindJSON(&trailers); err != nil {		
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
		
	if len(trailers.Plate) > 9 || len(trailers.Plate) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula Invalida!"})
		return
	}

	if trailers.Year <= 1960 || trailers.Year >= 2040  {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Ano Invalido!"})
		return
	}
	
	services.OpenDatabase()
	services.Db.Exec("update trailers set plate = ?, year = ? where id = ?", trailers.Plate,trailers.Year,trailers.ID)
		
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
}



type trailerstateaux struct {
	
	TruckId string ` json:"truck_id"`
	Latitude   string ` json:"latitude"`
	Longitude   string ` json:"longitude"`
	
}

func TrailerState (c *gin.Context) {
	var truck trailerstateaux
	var trailerState, traileraux model.TrailerState

	if err := c.ShouldBindJSON(&truck); err != nil {
			
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	//Conversao dos dados para os tipos certos
	TruckIdAux, err := strconv.Atoi(truck.TruckId)
	LatitudeAux, err := strconv.ParseFloat(truck.Latitude, 8)
	LongitudeAux, err := strconv.ParseFloat(truck.Longitude, 8)
	
	
	trailerState.TrailerId = uint(TruckIdAux)
	trailerState.Latitude = LatitudeAux
	trailerState.Longitude = LongitudeAux
	
	fmt.Println(err)

	///Find para descobrir o reboque , se encontrar , atualizamos os valores senao encontrar adicionamos o reboque 
	services.OpenDatabase()
	services.Db.Find(&traileraux, "trailer_id = ?", trailerState.TrailerId)
	defer services.Db.Close()
	//Se for !=0 temos correspondencia , entao atualizamos
	if(traileraux.TrailerId != 0){
		
		services.OpenDatabase()
		services.Db.Exec("update trailer_states set latitude = ?, longitude = ? where trailer_id = ?", trailerState.Latitude, trailerState.Longitude,trailerState.TrailerId)
		
		defer services.Db.Close()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
		return
	}else {
		services.Db.Save(&trailerState)
		defer services.Db.Close()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
		return
	}

}