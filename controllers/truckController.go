package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

const latitudeUfp = 41.17323531013999
const longitudeUfp = -8.611167669296265

func GetAllTrucks(c *gin.Context) {
	var trucks []model.Truck
	services.OpenDatabase()
	services.Db.Find(&trucks)
	defer services.Db.Close()
	if len(trucks) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trucks})

}

func GetTruck(c *gin.Context) {
	var truck model.Truck

	id := c.Param("id")

	services.OpenDatabase()
	services.Db.Find(&truck, id)
	defer services.Db.Close()
	if len(truck.Plate) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": truck})

}

/// CORRIGIR A LATITUDE E LONGITUDE DO CAMIAO, PARA QUANDO CAMIAO É CRIADO SER IGUAL À LOCALIZAÇÃO DA EMPRESA
func AddTrucks(c *gin.Context) {
	var trucks model.Truck
	var trucksDuplicated model.Truck
	var truckHistory model.TruckHistory
	var truckState model.TruckState

	if err := c.ShouldBindJSON(&trucks); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(trucks.Plate) > 9 || len(trucks.Plate) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula Invalida!"})
		return
	}

	if trucks.Year <= 1960 || trucks.Year >= 2040 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Ano Invalido!"})
		return
	}

	if trucks.Month < 1 || trucks.Month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Mes Invalido!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&trucksDuplicated, "plate = ?", trucks.Plate)
	if len(trucksDuplicated.Plate) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula já existe!"})
		return
	}

	defer services.Db.Close()
	trucks.Latitude = latitudeUfp
	trucks.Longitude = longitudeUfp

	services.OpenDatabase()
	services.Db.Save(&trucks)
	defer services.Db.Close()
	truckHistory.Latitude = trucks.Latitude
	truckHistory.Longitude = trucks.Longitude
	truckHistory.TruckId = trucks.ID
	truckHistory.Rpm = 0
	truckHistory.Speed = 0

	services.OpenDatabase()
	services.Db.Save(&truckHistory)

	truckState.Latitude = latitudeUfp
	truckState.Longitude = longitudeUfp
	truckState.TruckId = trucks.ID
	truckState.Rpm = 0
	truckState.Speed = 0
	services.Db.Save(&truckState)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})

}

func DeleteTrucks(c *gin.Context) {
	var truck model.Truck

	id := c.Param("id")
	services.OpenDatabase()
	services.Db.First(&truck, id)

	if truck.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
		return
	}

	services.Db.Exec("delete from trucks where id = ?", truck.ID)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func UpdateTruck(c *gin.Context) {
	var trucks model.Truck
	var trucksDuplicated model.Truck

	id := c.Param("id")
	if err := c.ShouldBindJSON(&trucks); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(trucks.Plate) > 9 || len(trucks.Plate) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Matricula Invalida!"})
		return
	}
	if trucks.Year <= 1960 || trucks.Year >= 2040 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Ano Invalido!"})
		return
	}

	if trucks.Month < 1 || trucks.Month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Mes Invalido!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&trucksDuplicated, "id = ?", id)
	if len(trucksDuplicated.Plate) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao não encontrado!"})
		return
	}
	services.Db.Exec("update trucks set plate = ?, year = ?, month = ?, km = ? , brand = ? where id = ?", trucks.Plate, trucks.Year, trucks.Month, trucks.Km, trucks.Brand, trucks.ID)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
}

type tstateaux struct {
	TruckId   string ` json:"truck_id"`
	Latitude  string ` json:"latitude"`
	Longitude string ` json:"longitude"`
	Rpm       string ` json:"rpm"`
	Speed     string ` json:"speed"`
}

func TruckState(c *gin.Context) {
	var truck tstateaux
	var truckState, truckaux model.TruckState
	var truckHistory model.TruckHistory
	var updateTruck model.Truck
	var displacements model.Displacements
	var trailers model.Trailer
	var trailerState model.TrailerState

	if err := c.ShouldBindJSON(&truck); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	//Conversao dos dados para os tipos certos
	TruckIdAux, err := strconv.Atoi(truck.TruckId)
	LatitudeAux, err := strconv.ParseFloat(truck.Latitude, 8)
	LongitudeAux, err := strconv.ParseFloat(truck.Longitude, 8)
	RpmAux, err := strconv.Atoi(truck.Rpm)
	SpeedAux, err := strconv.Atoi(truck.Speed)

	truckState.TruckId = uint(TruckIdAux)
	truckState.Latitude = LatitudeAux
	truckState.Longitude = LongitudeAux
	truckState.Rpm = uint(RpmAux)
	truckState.Speed = uint(SpeedAux)
	fmt.Println(err)

	if truckState.Latitude != 0 && truckState.Longitude != 0 { // se for diferente de 0 então deixa entrar na bd

		///Find para descobrir o camiao , se encontrar , atualizamos os valores senao encontrar adicionamos o camiao
		services.OpenDatabase()
		services.Db.Find(&truckaux, "truck_id = ?", truckState.TruckId)
		defer services.Db.Close()
		services.OpenDatabase()
		services.Db.Find(&updateTruck, "id = ?", truck.TruckId) // atualizamos a posição do camiao tb
		defer services.Db.Close()
		//Se for !=0 temos correspondencia , entao atualizamos
		if truckaux.TruckId != 0 {
			//Se os rpm estao a 0 entao a query de update é diferente
			if truckState.Rpm == 0 {
				services.OpenDatabase()
				services.Db.Exec("update truck_states set latitude = ?, longitude = ?, rpm = ?, speed = ? where truck_id = ?", truckState.Latitude, truckState.Longitude, truckState.Rpm, truckState.Speed, truckState.TruckId)
				services.Db.Exec("update trucks set latitude = ?, longitude = ? where id = ?", truckState.Latitude, truckState.Longitude, truckState.TruckId)

				// Sempre que atualizamos o estado colocamos as novas localizações no historico do camiao
				truckHistory.TruckId = truckState.TruckId
				truckHistory.Latitude = truckState.Latitude
				truckHistory.Longitude = truckState.Longitude
				truckHistory.Rpm = truckState.Rpm
				truckHistory.Speed = truckState.Speed
				services.Db.Save(&truckHistory)
				defer services.Db.Close()
				// Verificamos se o camião tem deslocamento
				// Se tiver deslocamento pegamos na localizaçao do camiao e do reboque e verificamos se estão no mesmo raio
				// Se tiverem no mesmo raio entao o reboque vai ser "teletransportado para junto do Camião"
				services.OpenDatabase()
				services.Db.Find(&displacements, "truck_id = ?", truckState.TruckId)
				defer services.Db.Close()
				if displacements.ID != 0 {
					services.OpenDatabase()
					services.Db.Find(&trailers, "id = ?", displacements.TrailerId) // pego no trailer
					// truck eu ja tenho por isso é so comparar
					if trailers.ID != 0 {
						// Fazer o parse do float para retirar a coordenada exata e só meter a aproximada
						latTrailer := fmt.Sprintf("%.2f", trailers.Latitude)
						longTrailer := fmt.Sprintf("%.2f", trailers.Longitude)
						latTruck := fmt.Sprintf("%.2f", truckState.Latitude)
						longTruck := fmt.Sprintf("%.2f", truckState.Longitude)
						if latTrailer == latTruck && longTrailer == longTruck {
							services.Db.Exec("update trailers set latitude = ?, longitude = ? where id = ?", truckState.Latitude, truckState.Longitude, trailers.ID) // colocamos o trailer a atualizar com o camiao
							trailerState.Latitude = truckState.Latitude
							trailerState.Longitude = truckState.Longitude
							services.Db.Save(trailerState)
						}
					}
					defer services.Db.Close()
				}
				//
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
				return
			}
			services.OpenDatabase()
			services.Db.Exec("update truck_states set latitude = ?, longitude = ?, rpm = ?, speed = ? where truck_id = ?", truckState.Latitude, truckState.Longitude, truckState.Rpm, truckState.Speed, truckState.TruckId)
			services.Db.Exec("update trucks set latitude = ?, longitude = ? where id = ?", truckState.Latitude, truckState.Longitude, truckState.TruckId)
			// Sempre que atualizamos o estado colocamos as novas localizações no historico do camiao
			truckHistory.TruckId = truckState.TruckId
			truckHistory.Latitude = truckState.Latitude
			truckHistory.Longitude = truckState.Longitude
			truckHistory.Rpm = truckState.Rpm
			truckHistory.Speed = truckState.Speed
			services.Db.Save(&truckHistory)

			// Verificamos se o camião tem deslocamento
			// Se tiver deslocamento pegamos na localizaçao do camiao e do reboque e verificamos se estão no mesmo raio
			// Se tiverem no mesmo raio entao o reboque vai ser "teletransportado para junto do Camião"
			services.OpenDatabase()
			services.Db.Find(&displacements, "truck_id = ?", truckState.TruckId)
			if displacements.ID != 0 {
				services.OpenDatabase()
				services.Db.Find(&trailers, "id = ?", displacements.TrailerId) // pego no trailer
				// truck eu ja tenho por isso é so comparar
				if trailers.ID != 0 {
					// Fazer o parse do float para retirar a coordenada exata e só meter a aproximada
					latTrailer := fmt.Sprintf("%.2f", trailers.Latitude)
					longTrailer := fmt.Sprintf("%.2f", trailers.Longitude)
					latTruck := fmt.Sprintf("%.2f", truckState.Latitude)
					longTruck := fmt.Sprintf("%.2f", truckState.Longitude)
					if latTrailer == latTruck && longTrailer == longTruck {
						services.Db.Exec("update trailers set latitude = ?, longitude = ? where id = ?", truckState.Latitude, truckState.Longitude, trailers.ID) // colocamos o trailer a atualizar com o camiao
						trailerState.Latitude = truckState.Latitude
						trailerState.Longitude = truckState.Longitude
						services.Db.Save(trailerState)
					}
				}
			}
			//
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
			return
		} else {
			services.OpenDatabase()
			services.Db.Save(&truckState)
			defer services.Db.Close()
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
			return
		}
	}

}

func GetTruckCount(c *gin.Context) {
	var data [3]int

	///Fazemos o count para os camioes desligados
	services.OpenDatabase()
	services.Db.Model(&model.TruckState{}).Where("rpm = 0").Count(&data[0])
	defer services.Db.Close()
	///Fazemos o count para os camioes ligados mas ao relantim
	services.OpenDatabase()
	services.Db.Model(&model.TruckState{}).Where("rpm >= 600 and rpm <= 1000").Count(&data[1])
	defer services.Db.Close()
	services.OpenDatabase()
	///Fazemos o count para os camioes em deslocação
	services.Db.Model(&model.TruckState{}).Where("rpm > 1001").Count(&data[2])

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})

}

func GetRpmPlates(c *gin.Context) {
	var rpmzero, rpmmid, rpmrun []model.TruckState
	var rzero, rmid, rrun []model.Truck
	var rzeroaux, rmidaux, rrunaux []model.Truck
	services.OpenDatabase()
	///Find dos camioes com rpm a 0
	services.Db.Find(&rpmzero, "rpm = 0")
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&rzero)
	defer services.Db.Close()

	for i := 0; i < len(rpmzero); i++ {
		for x := 0; x < len(rzero); x++ {
			if rpmzero[i].TruckId == (uint)(rzero[x].ID) {
				rzeroaux = append(rzeroaux, rzero[i])
			}
		}
	}

	///Find dos camioes ao relantim
	services.OpenDatabase()
	services.Db.Where("rpm >= 600 AND rpm <= 1000").Find(&rpmmid)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&rmid)
	//services.Db.Find(&rpmmid, "rpm >= 600 and rpm <= 1000")
	for i := 0; i < len(rpmmid); i++ {
		for x := 0; x < len(rmid); x++ {
			if rpmmid[i].TruckId == (uint)(rmid[x].ID) {
				rmidaux = append(rmidaux, rmid[i])
			}
		}
	}
	defer services.Db.Close()
	///Find dos camioes em andamento
	services.OpenDatabase()
	services.Db.Find(&rpmrun, "rpm > 1001")
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&rrun)
	for i := 0; i < len(rpmrun); i++ {
		for x := 0; x < len(rrun); x++ {
			if rpmrun[i].TruckId == (uint)(rrun[x].ID) {
				rrunaux = append(rrunaux, rrun[i])
			}
		}
	}

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "rpmzero": rzeroaux, "rpmmid": rmidaux, "rpmrun": rrunaux})

}

type data struct {
	TruckId   uint    ` json:"truck_id"`
	Latitude  float64 ` json:"latitude"`
	Longitude float64 ` json:"longitude"`
	Rpm       uint    ` json:"rpm"`
	Speed     uint    ` json:"speed"`
	Plate     string  ` json:"plate"`
}

func AllTruckState(c *gin.Context) {

	var dados []data
	var aux data
	var trucksState []model.TruckState
	var trucks []model.Truck
	services.OpenDatabase()
	services.Db.Find(&trucks)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&trucksState)

	for i := 0; i < len(trucks); i++ {
		for x := 0; x < len(trucksState); x++ {
			if trucks[i].ID == (trucksState[x].TruckId) {
				aux.TruckId = trucks[i].ID
				aux.Plate = trucks[i].Plate
				aux.Latitude = trucksState[x].Latitude
				aux.Longitude = trucksState[x].Longitude
				aux.Rpm = trucksState[x].Rpm
				aux.Speed = trucksState[x].Speed
				dados = append(dados, aux)
			}
		}
	}

	defer services.Db.Close()

	fmt.Print(dados)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dados})
}

func GetTruckHistory(c *gin.Context) {
	var truckHistory []model.TruckHistory
	id := c.Param("id")

	services.OpenDatabase()
	services.Db.Find(&truckHistory, "truck_id = ?", id)
	if len(truckHistory) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Camião não tem histórico"})
		defer services.Db.Close()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": truckHistory})
	defer services.Db.Close()
}
