package controllers

import (
	// "fmt"
	// "strconv"

	// "io"
	// "strings"

	"net/http"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
	// "github.com/otiai10/gosseract/v2"
)

type Post struct {
	Text string
}

//func to create supply
func CreateSupply(c *gin.Context) {
	// var supply model.Supply
	// var displacement model.Displacements
	// var supplyAndDisplacement model.DisplacementsAndSupply
	// var liters, value string

	// supply.TruckId = c.Request.FormValue("truck_id")
	// supply.DriverId = c.Request.FormValue("driver_id")

	// file, header, err := c.Request.FormFile("file")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
	// 	return
	// }

	// if header != nil {
	// 	fmt.Printf("")
	// }
	// client := gosseract.NewClient()
	// defer client.Close()
	// bytes, err := io.ReadAll(file)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't open file"})
	// 	return
	// }
	// client.SetImageFromBytes(bytes)
	// text, _ := client.Text()

	// i := strings.Index(text, "ER")
	// if i != -1 {
	// 	value = string(text[i-6]) + string(text[i-5]) + string(text[i-4]) + string(text[i-3]) + string(text[i-2]) + string(text[i-1])
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't extract info"})
	// 	return
	// }
	// i = strings.Index(text, ".L")

	// if i != -1 {
	// 	liters = string(text[i-5]) + string(text[i-4]) + string(text[i-3]) + string(text[i-2]) + string(text[i-1])
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "couldn't extract info"})
	// 	return
	// }

	// if s, err := strconv.ParseFloat(strings.TrimSpace(value), 64); err == nil {
	// 	supply.Value = s
	// }
	// if x, err := strconv.ParseFloat(liters, 64); err == nil {
	// 	supply.Liters = x
	// }

	// ///gravamos a supply na bd
	// services.OpenDatabase()
	// services.Db.Save(&supply).
	// 	/// return last id from supply table
	// 	Last(&supply)
	// /// find displacement where truck_id = supply.truck_id
	// /// create displacement variable

	// services.Db.Where("truck_id = ?", supply.TruckId).Find(&displacement)
	// supplyAndDisplacement.SupplyId = supply.ID
	// supplyAndDisplacement.DisplacementId = displacement.ID
	// services.Db.Save(&supplyAndDisplacement)

	c.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

type Stats struct {
	MostDistance          float64
	DriverMostDistance    uint
	MostSupplyValue       float64
	DriverMostSupplyValue uint
}

//func to get stats
func GetStats(c *gin.Context) {
	var travelMaps []model.TravelMaps
	var stats Stats

	services.OpenDatabase()
	services.Db.Find(&travelMaps)
	defer services.Db.Close()

	for i := 0; i < len(travelMaps); i++ {
		if travelMaps[i].Distance > stats.MostDistance {
			stats.MostDistance = travelMaps[i].Distance
			stats.DriverMostDistance = uint(travelMaps[i].TruckId)
		}
		if travelMaps[i].TotalValue > stats.MostSupplyValue {
			stats.MostSupplyValue = travelMaps[i].TotalValue
			stats.DriverMostSupplyValue = uint(travelMaps[i].TruckId)
		}

	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": stats})
}

// Lista apenas os camionistas que tem camioes associados

func ListTrucksDrivers(c *gin.Context) {
	var truckAndDriver []model.TruckAndDriver
	var users []model.Users
	var truckDrivers []model.Users

	services.OpenDatabase()
	services.Db.Find(&users)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&truckAndDriver)
	defer services.Db.Close()

	for i := 0; i < len(users); i++ {
		if users[i].RoleId == 4 { // verifica se o cargo do user é camionista
			for x := 0; x < len(truckAndDriver); x++ {
				if users[i].ID == (uint)(truckAndDriver[x].FirstDriverId) {
					truckDrivers = append(truckDrivers, users[i])
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "truckDriver": truckDrivers})
}

func ListTrailers(c *gin.Context) {
	var trailers []model.Trailer
	services.OpenDatabase()
	services.Db.Find(&trailers)

	if len(trailers) <= 0 {
		defer services.Db.Close()
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "trailer": trailers})

}

func ListTruckAndDriver(c *gin.Context) {
	var trailers []model.TruckAndDriver
	services.OpenDatabase()
	services.Db.Find(&trailers)

	if len(trailers) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": trailers})

}

func CreateDisplacement(c *gin.Context) {
	var truckAndDriver model.TruckAndDriver
	var trailer model.Trailer
	var displacement model.Displacements
	var data struct {
		Coords          string  `gorm:"not null" json:"coords"`
		TruckDriver     uint    `gorm:"not null" json:"truck_driver"`
		TrailerPlate    string  `json:"trailer_plate"`
		Distance        float64 `json:distance`
		Time            string  `json:"time"`
		StartCountry    string  `json:"start_country"`
		StartCity       string  `json:"start_city"`
		StartPostalCode string  `json:"start_postal_code"`
		StartAddress    string  `json:"start_address"`
		EndCountry      string  `json:"end_country"`
		EndCity         string  `json:"end_city"`
		EndPostalCode   string  `json:"end_postal_code"`
		EndAddress      string  `json:"end_address"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&truckAndDriver, "first_driver_id = ?", data.TruckDriver)
	defer services.Db.Close()
	services.OpenDatabase()
	services.Db.Find(&trailer, "plate = ?", data.TrailerPlate)

	//aux = data.Distance;
	displacement.Coords = data.Coords
	displacement.TrailerId = trailer.ID
	displacement.TruckId = truckAndDriver.TruckId
	displacement.Distance = data.Distance
	displacement.StartAddress = data.StartAddress
	displacement.StartCity = data.StartCity
	displacement.StartCountry = data.StartCountry
	displacement.StartPostalCode = data.StartPostalCode
	displacement.EndAddress = data.EndAddress
	displacement.EndCity = data.EndCity
	displacement.EndCountry = data.EndCountry
	displacement.EndPostalCode = data.EndPostalCode
	displacement.Time = data.Time

	services.Db.Save(&displacement)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
}
