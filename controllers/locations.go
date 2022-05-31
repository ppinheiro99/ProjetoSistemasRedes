package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gestaoFrota/model"
	"github.com/gestaoFrota/services"
	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	var locations []model.Locations
	services.OpenDatabase()
	services.Db.Find(&locations)
	defer services.Db.Close()
	if len(locations) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": locations})
}

func GetLocation(c *gin.Context) {
	var locations model.Locations

	id := c.Param("id")

	services.OpenDatabase()
	services.Db.Find(&locations, id)
	defer services.Db.Close()
	if locations.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": locations})
}

func AddLocations(c *gin.Context) {
	var locations model.Locations

	if err := c.ShouldBindJSON(&locations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(locations.Name) < 3 || len(locations.Name) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Nome inválido!"})
		return
	}

	services.OpenDatabase()
	services.Db.Save(&locations)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
}

func DeleteLocation(c *gin.Context) {
	var locations model.Locations

	id := c.Param("id")

	services.OpenDatabase()
	services.Db.First(&locations, id)

	if locations.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
		return
	}

	services.Db.Exec("delete from locations where id = ?", locations.ID)
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func UpdateLocation(c *gin.Context) { // nao está com a segurança máxima mas para já come-se
	var locations model.Locations
	var locationsUpdate model.Locations

	if err := c.ShouldBindJSON(&locations); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if len(locations.Name) < 3 || len(locations.Name) > 50 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Nome Inválido!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&locationsUpdate, "id = ?", locations.ID)

	if locationsUpdate.ID == 0 {
		defer services.Db.Close()
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Camiao não encontrado!"})
		return
	}
	services.Db.Exec("update locations set name = ?, latitude = ?, longitude = ? where id = ?", locations.Name, locations.Latitude, locations.Longitude, locations.ID)

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})

}

type getTruckLocation struct {
	Type     string    `json:"type"`
	Query    []float64 `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			Foursquare string `json:"foursquare"`
			Landmark   bool   `json:"landmark"`
			Wikidata   string `json:"wikidata"`
			Category   string `json:"category"`
			Maki       string `json:"maki"`
		} `json:"properties,omitempty"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			Text      string `json:"text"`
			Wikidata  string `json:"wikidata,omitempty"`
			ShortCode string `json:"short_code,omitempty"`
		} `json:"context,omitempty"`
		_ struct {
		} `json:"properties,omitempty"`
		Bbox []float64 `json:"bbox,omitempty"`
		_    struct {
			Wikidata string `json:"wikidata"`
		} `json:"properties,omitempty"`
		_ struct {
			Wikidata string `json:"wikidata"`
		} `json:"properties,omitempty"`
		_ struct {
			Wikidata  string `json:"wikidata"`
			ShortCode string `json:"short_code"`
		} `json:"properties,omitempty"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}

type DistritoCode struct {
	Status    bool   `json:"status"`
	Mensagem  string `json:"mensagem"`
	Resultado []struct {
		Descritivo string `json:"Descritivo"`
		ID         int    `json:"Id"`
	} `json:"resultado"`
	Token interface{} `json:"token"`
}

type MunicipioCode struct {
	Status    bool   `json:"status"`
	Mensagem  string `json:"mensagem"`
	Resultado []struct {
		Descritivo string `json:"Descritivo"`
		IDDistrito int    `json:"IdDistrito"`
		Distrito   struct {
			Descritivo string `json:"Descritivo"`
			ID         int    `json:"Id"`
		} `json:"Distrito"`
		ID int `json:"Id"`
	} `json:"resultado"`
	Token interface{} `json:"token"`
}

type GasStationStruct struct {
	Status    bool   `json:"status"`
	Mensagem  string `json:"mensagem"`
	Resultado []struct {
		ID              int     `json:"Id"`
		Nome            string  `json:"Nome"`
		TipoPosto       string  `json:"TipoPosto"`
		Municipio       string  `json:"Municipio"`
		Preco           string  `json:"Preco"`
		Marca           string  `json:"Marca"`
		Combustivel     string  `json:"Combustivel"`
		DataAtualizacao string  `json:"DataAtualizacao"`
		Distrito        string  `json:"Distrito"`
		Morada          string  `json:"Morada"`
		Localidade      string  `json:"Localidade"`
		CodPostal       string  `json:"CodPostal"`
		Latitude        float64 `json:"Latitude"`
		Longitude       float64 `json:"Longitude"`
		Quantidade      int     `json:"Quantidade"`
	} `json:"resultado"`
	Token interface{} `json:"token"`
}

type DistanceBetCoordsStruct struct {
	Routes []struct {
		CountryCrossed bool    `json:"country_crossed"`
		WeightName     string  `json:"weight_name"`
		Weight         float64 `json:"weight"`
		Duration       float64 `json:"duration"`
		Distance       float64 `json:"distance"`
		Legs           []struct {
			ViaWaypoints []interface{} `json:"via_waypoints"`
			Admins       []struct {
				Iso31661Alpha3 string `json:"iso_3166_1_alpha3"`
				Iso31661       string `json:"iso_3166_1"`
			} `json:"admins"`
			Weight   float64 `json:"weight"`
			Duration float64 `json:"duration"`
			Steps    []struct {
				BannerInstructions []struct {
					Sub struct {
						Components []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						} `json:"components"`
						Type     string `json:"type"`
						Modifier string `json:"modifier"`
						Text     string `json:"text"`
					} `json:"sub"`
					Primary struct {
						Components []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						} `json:"components"`
						Type     string `json:"type"`
						Modifier string `json:"modifier"`
						Text     string `json:"text"`
					} `json:"primary"`
					DistanceAlongGeometry float64 `json:"distanceAlongGeometry"`
				} `json:"bannerInstructions"`
				VoiceInstructions []struct {
					SsmlAnnouncement      string  `json:"ssmlAnnouncement"`
					Announcement          string  `json:"announcement"`
					DistanceAlongGeometry float64 `json:"distanceAlongGeometry"`
				} `json:"voiceInstructions"`
				Intersections []struct {
					Entry           []bool  `json:"entry"`
					Bearings        []int   `json:"bearings"`
					Duration        float64 `json:"duration,omitempty"`
					MapboxStreetsV8 struct {
						Class string `json:"class"`
					} `json:"mapbox_streets_v8"`
					IsUrban       bool      `json:"is_urban"`
					AdminIndex    int       `json:"admin_index"`
					Out           int       `json:"out"`
					Weight        float64   `json:"weight,omitempty"`
					GeometryIndex int       `json:"geometry_index"`
					Location      []float64 `json:"location"`
					In            int       `json:"in,omitempty"`
					TurnWeight    float64   `json:"turn_weight,omitempty"`
					TurnDuration  float64   `json:"turn_duration,omitempty"`
				} `json:"intersections"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Name        string  `json:"name"`
				Duration    float64 `json:"duration"`
				Distance    float64 `json:"distance"`
				DrivingSide string  `json:"driving_side"`
				Weight      float64 `json:"weight"`
				Mode        string  `json:"mode"`
				Geometry    string  `json:"geometry"`
				_           struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Ref string `json:"ref,omitempty"`
				_   struct {
					Type          string    `json:"type"`
					Exit          int       `json:"exit"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				RotaryName string `json:"rotary_name,omitempty"`
				_          struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
			} `json:"steps"`
			Distance float64 `json:"distance"`
			Summary  string  `json:"summary"`
		} `json:"legs"`
		Geometry    string `json:"geometry"`
		VoiceLocale string `json:"voiceLocale"`
	} `json:"routes"`
	Waypoints []struct {
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
	Code string `json:"code"`
	UUID string `json:"uuid"`
}

type GasStationInfoStruct struct {
	Nome        string  `json:"nome"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Distrito    string  `json:"distrito"`
	Localidade  string  `json:"localidade"`
	Preco       string  `json:"preco"`
	Combustivel string  `json:"combustivel"`
	Distancia   float64 `json:"distancia"`
}

func BestGasStation(c *gin.Context) {

	// Obter a localização do camião
	var truck model.Truck
	// APIS DATA
	var dataApi getTruckLocation
	var dataDistritoCodes DistritoCode
	var dataMunicipioCodes MunicipioCode
	var dataGasStation GasStationStruct
	var distanceBetCoordsData DistanceBetCoordsStruct
	var gasStationInfo []GasStationInfoStruct
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&truck, "id = ?", id)

	if truck.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Camiao não encontrado!"})
		return
	}

	// Obter os "códigos" dos distritos para utilizar nos próximos endpoints
	response, err := http.Get("https://precoscombustiveis.dgeg.gov.pt/api/PrecoComb/GetDistritos")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
		return
	}

	districtsCodes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	lng := fmt.Sprint(truck.Longitude)
	lat := fmt.Sprint(truck.Latitude)

	// Obter a minha localização (distrito e municipio) com base na latitude e longitude do camião
	response1, err := http.Get("https://api.mapbox.com/geocoding/v5/mapbox.places/" + lng + "," + lat + ".json?access_token=pk.eyJ1IjoicHBpbmhlaXJvOTkiLCJhIjoiY2tvbXI3N3NuMXIzejJxcm1uZTJxOWlmYyJ9.kJ7suLcfyHd8DBr34XygRw")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
		return
	}

	truckLocationData, err := ioutil.ReadAll(response1.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(districtsCodes), &dataDistritoCodes)
	json.Unmarshal([]byte(truckLocationData), &dataApi)

	// Ver o código pertencente ao Distrito e ao Municipio do camião
	var districtCode int
	for _, s := range dataDistritoCodes.Resultado {
		if s.Descritivo == dataApi.Features[0].Context[2].Text {
			districtCode = s.ID
		}
	}
	districtCodeString := fmt.Sprint(districtCode)
	response2, err := http.Get("https://precoscombustiveis.dgeg.gov.pt/api/PrecoComb/GetMunicipios?idDistrito=" + districtCodeString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
		return
	}

	MunCodes, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(MunCodes), &dataMunicipioCodes)

	var municipioCode int
	for _, s := range dataMunicipioCodes.Resultado {
		if s.Descritivo == dataApi.Features[0].Context[1].Text {
			municipioCode = s.ID
		}
	}

	municipioCodeString := fmt.Sprint(municipioCode)

	// Obter os Postos de abastecimento do concelho
	response3, err := http.Get("https://precoscombustiveis.dgeg.gov.pt/api/PrecoComb/PesquisarPostos?idsTiposComb=2101&idMarca=&idTipoPosto=&idDistrito=" + districtCodeString + "&idsMunicipios=" + municipioCodeString + "&qtdPorPagina=1000&pagina=1")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
		return
	}

	allgasStationData, err := ioutil.ReadAll(response3.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(allgasStationData), &dataGasStation)
	raioKmString := c.Param("r")
	raioKm, _ := strconv.ParseFloat(raioKmString, 64)
	raioKm = raioKm * 1000
	// Percorrer todas as bombas de combustível, filtrar pelas que estão "dentro" do raio de km e colocar em um array de tamanho 3(as 3 bombas mais baratas no raio de km)
	for _, s := range dataGasStation.Resultado {
		auxLng := fmt.Sprint(s.Longitude)
		auxLat := fmt.Sprint(s.Latitude)
		// Para isso, em cada bomba de combustível chamo uma Api(ver no notePad++ o endpoint para isto) que me dá a distancia entre 2 coordenadas
		response4, err := http.Get("https://api.mapbox.com/directions/v5/mapbox/driving/" + lng + "," + lat + ";" + auxLng + "," + auxLat + "?waypoints=0;1&steps=true&voice_instructions=true&banner_instructions=true&voice_units=imperial&roundabout_exits=true&access_token=pk.eyJ1IjoicHBpbmhlaXJvOTkiLCJhIjoiY2tvbXI3N3NuMXIzejJxcm1uZTJxOWlmYyJ9.kJ7suLcfyHd8DBr34XygRw")
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
			return
		}
		distanceBetCoords, err := ioutil.ReadAll(response4.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(distanceBetCoords), &distanceBetCoordsData)
		// Comparar as distancia obtida com o raio de km que desejo procurar e adiciono ao array
		if distanceBetCoordsData.Routes[0].Distance <= raioKm {
			gasStationInfo = append(gasStationInfo, GasStationInfoStruct{
				Nome:        s.Nome,
				Latitude:    s.Latitude,
				Longitude:   s.Longitude,
				Distrito:    s.Distrito,
				Localidade:  s.Localidade,
				Preco:       s.Preco,
				Combustivel: s.Combustivel,
				Distancia:   distanceBetCoordsData.Routes[0].Distance,
			})
		}
	}
	if len(gasStationInfo) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Não existem bombas de combustível no raio de " + raioKmString + "km"})
		return
	}
	numBombasAEncontrar := c.Param("n")
	numBombasAEncontrarInt, _ := strconv.Atoi(numBombasAEncontrar)
	var warning = ""
	if len(gasStationInfo) < numBombasAEncontrarInt {
		numBombasAEncontrarInt = len(gasStationInfo)
		warning = "No raio introduzido não foram encontradas bombas de combustível suficientes para o número de bombas pedido"
	}

	var resultBestGasStation []GasStationInfoStruct
	for i := 0; i < numBombasAEncontrarInt; i++ {
		var minStruct, index = FindMin(gasStationInfo)
		resultBestGasStation = append(resultBestGasStation, minStruct)
		gasStationInfo = append(gasStationInfo[:index], gasStationInfo[index+1:]...)
	}

	if warning != "" {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": warning, "gasStationInfo": resultBestGasStation})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Bombas de combustível encontradas", "gasStationInfo": resultBestGasStation})
}

func FindMin(products []GasStationInfoStruct) (min GasStationInfoStruct, index int) {
	min = products[0]
	for i, product := range products {
		if product.Preco < min.Preco {
			min = product
			index = i
		}
	}
	return min, index
}

type NewRouteStruct struct {
	Routes []struct {
		CountryCrossed  bool    `json:"country_crossed"`
		WeightTypical   float64 `json:"weight_typical"`
		DurationTypical float64 `json:"duration_typical"`
		WeightName      string  `json:"weight_name"`
		Weight          float64 `json:"weight"`
		Duration        float64 `json:"duration"`
		Distance        float64 `json:"distance"`
		Legs            []struct {
			ViaWaypoints []interface{} `json:"via_waypoints"`
			Admins       []struct {
				Iso31661Alpha3 string `json:"iso_3166_1_alpha3"`
				Iso31661       string `json:"iso_3166_1"`
			} `json:"admins"`
			WeightTypical   float64 `json:"weight_typical"`
			DurationTypical float64 `json:"duration_typical"`
			Weight          float64 `json:"weight"`
			Duration        float64 `json:"duration"`
			Steps           []struct {
				BannerInstructions []struct {
					Primary struct {
						Components []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						} `json:"components"`
						Type     string `json:"type"`
						Modifier string `json:"modifier"`
						Text     string `json:"text"`
					} `json:"primary"`
					DistanceAlongGeometry float64 `json:"distanceAlongGeometry"`
				} `json:"bannerInstructions"`
				VoiceInstructions []struct {
					SsmlAnnouncement      string  `json:"ssmlAnnouncement"`
					Announcement          string  `json:"announcement"`
					DistanceAlongGeometry float64 `json:"distanceAlongGeometry"`
				} `json:"voiceInstructions"`
				Intersections []struct {
					Bearings        []int  `json:"bearings"`
					Entry           []bool `json:"entry"`
					MapboxStreetsV8 struct {
						Class string `json:"class"`
					} `json:"mapbox_streets_v8"`
					IsUrban       bool      `json:"is_urban"`
					AdminIndex    int       `json:"admin_index"`
					Out           int       `json:"out"`
					GeometryIndex int       `json:"geometry_index"`
					Location      []float64 `json:"location"`
				} `json:"intersections"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Name            string  `json:"name"`
				WeightTypical   float64 `json:"weight_typical"`
				DurationTypical float64 `json:"duration_typical"`
				Duration        float64 `json:"duration"`
				Distance        float64 `json:"distance"`
				DrivingSide     string  `json:"driving_side"`
				Weight          float64 `json:"weight"`
				Mode            string  `json:"mode"`
				Geometry        string  `json:"geometry"`
				_               struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Destinations string `json:"destinations,omitempty"`
				_            struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Ref string `json:"ref,omitempty"`
				_   struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				Exits string `json:"exits,omitempty"`
				_     struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Exit          int       `json:"exit"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Exit          int       `json:"exit"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Exit          int       `json:"exit"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
				_ struct {
					Type          string    `json:"type"`
					Instruction   string    `json:"instruction"`
					Modifier      string    `json:"modifier"`
					BearingAfter  int       `json:"bearing_after"`
					BearingBefore int       `json:"bearing_before"`
					Location      []float64 `json:"location"`
				} `json:"_,omitempty"`
			} `json:"steps"`
			Distance float64 `json:"distance"`
			Summary  string  `json:"summary"`
		} `json:"legs"`
		Geometry    string `json:"geometry"`
		VoiceLocale string `json:"voiceLocale"`
	} `json:"routes"`
	Waypoints []struct {
		Distance int       `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
	Code string `json:"code"`
	UUID string `json:"uuid"`
}

func RWithoutTrafficCongestion(c *gin.Context) {
	// Obter a localização do camião aonde o caminista se encontra
	var truck model.Truck
	var displacement model.Displacements
	var newRouteDataStruct NewRouteStruct
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.OpenDatabase()
	services.Db.Find(&truck, "id = ?", id)

	if truck.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Camiao não encontrado!"})
		return
	}

	// fmt.Println(truck.ID)
	// fmt.Println(truck.Latitude)
	// fmt.Println(truck.Longitude)

	// Ver a rota atribuida ao camião e pegar na localização de destino
	// Ver o deslocamento associado ao camião
	// services.Db.Find(&displacement, "truck_id = ?", truck.ID)

	var_select := "SELECT * FROM displacements "
	var_select += "WHERE truck_id = " + fmt.Sprintf("%d", truck.ID)
	var_select += " LIMIT 1"

	// fmt.Println(var_select)

	services.Db.Raw(var_select).Scan(&displacement)
	// fmt.Println(displacement.StartCity)
	// fmt.Println(displacement.EndCity)
	// fmt.Println(displacement.ID)
	//fmt.Println(displacement.Coords)
	// fmt.Println()
	// Obter as coordenadas iniciais e finais do deslocamento
	s := strings.Split(displacement.Coords, ",")
	beginLat := strings.Replace(s[0], "[", "", -1)
	beginLng := s[1]
	endLat := s[len(s)-2]
	endLng := strings.Replace(s[len(s)-1], "]", "", -1)

	// fmt.Println(displacement.Time)
	// fmt.Println(displacement.Distance)
	// fmt.Println(beginLat)
	// fmt.Println(beginLng)
	// fmt.Println(endLat)
	// fmt.Println(endLng)
	// Gerar a melhor rota (rota menos demorada, baseado no km e no transito, usar a api mapquest)
	response, err := http.Get("https://api.mapbox.com/directions/v5/mapbox/driving-traffic/" + beginLng + "," + beginLat + ";" + endLng + "," + endLat + "?waypoints=0;1&steps=true&voice_instructions=true&banner_instructions=true&voice_units=imperial&roundabout_exits=true&access_token=pk.eyJ1IjoicHBpbmhlaXJvOTkiLCJhIjoiY2tvbXI3N3NuMXIzejJxcm1uZTJxOWlmYyJ9.kJ7suLcfyHd8DBr34XygRw")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Erro no pedido à API"})
		return
	}
	distanceBetCoords, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(distanceBetCoords), &newRouteDataStruct)
	durationString := fmt.Sprintf("%d", int(newRouteDataStruct.Routes[0].Duration))
	hours, _ := time.ParseDuration(durationString + "s") // converter para Horas
	h := strings.Split(displacement.Time, " ")
	oldRouteTime, _ := time.ParseDuration(h[1] + h[2]) // converter para Horas

	var coordsString []string
	// Comparar os tempos da rota "originalmente" gerada e a nova rota gerada, se a nova rota tiver um tempo melhor então atualizamos a deslocação com as novas coords
	if hours.Seconds() < oldRouteTime.Seconds() {
		// Fazer parse dos dados e guardar todas as coordenadas numa string para de segida guardar na base de dados
		for _, s := range newRouteDataStruct.Routes[0].Legs[0].Steps { // CICLOS PARA CHEGAR AS COORDENADAS (API ENVIA MUITAS INFORMAÇÕES)
			for _, c := range s.Intersections {
				invertCoords := fmt.Sprintf("%f", c.Location[1]) + "," + fmt.Sprintf("%f", c.Location[0])
				coordsString = append(coordsString, invertCoords)
			}
		}
		// Atualizar a rota na BD
		displacement.Coords = "[" + strings.Join(coordsString, ",") + "]"
		displacement.Time = hours.String()
		displacement.Distance = (newRouteDataStruct.Routes[0].Distance / 1000) // converter para Km como tenho na BD
		// fmt.Println(displacement.Coords)
		// fmt.Println(displacement.Time)
		// fmt.Println(displacement.Distance)

		services.OpenDatabase()
		services.Db.Save(&displacement)
		defer services.Db.Close()
		services.Db.Exec("update displacements set coords = ?, distance = ?, time = ? where id = ?", displacement.Coords, displacement.Distance, displacement.Time, displacement.ID)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Sucesso!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Nova rota gerada possui tempo superior!"})
}
