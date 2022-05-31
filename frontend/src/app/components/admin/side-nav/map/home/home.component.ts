import { Component, AfterViewInit, Output, EventEmitter} from '@angular/core';
import { TrucksService } from "../../../../../services/trucks/trucks.service";

import "leaflet";
import { MapService } from 'src/app/services/map/map.service';
import { TrailersService } from 'src/app/services/trailers/trailers.service';
import { DatePipe } from '@angular/common';
import { UsersService } from 'src/app/services/user/users.service';
import { interval } from 'rxjs';
declare let L;

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements AfterViewInit {
  private map
  trucksMarks: any
  trailersMarks: any
  leafletIconTruck: any
  leafletIconTruckGreen: any
  leafletIconTruckOrange: any
  leafletIconTrailer: any
  wayPoints:any
  arrayCoords = []
  arrayCoordsAux = []
  polylineArray = []
  index = 0
  polyLineIndex = 0;
  arrayTruckDrivers = []
  arrayTrailers = []
  arrayTrucks = []
  arrayTruckAndDriver = []
  truckCoordsHistory = []
  truckCoordsHistoryAux = []
  distance :any
  distanceHistoryAux = 0
  time: any
  dataLocation :any
  user: any
  arrayCoordsRoute = []
  markersArray = []
  loopForever
  constructor(private trucksService: TrucksService,private trailerService: TrailersService ,private mapService: MapService, private datePipe: DatePipe, private userService: UsersService) {}

  ngAfterViewInit():void{
    this.initMap();
  }

  private initMap():void{
    this.map = L.map('map', {
      center: [ 41.17323531013999, -8.611167669296265 ],
      zoom: 10
    }).on('click', (e) => {
      this.arrayCoords.push(e)
      if(this.arrayCoords.length > 1){
        var pointA = new L.LatLng(this.arrayCoords[this.index-1].latlng.lat, this.arrayCoords[this.index-1].latlng.lng);
        var pointB = new L.LatLng(this.arrayCoords[this.index].latlng.lat, this.arrayCoords[this.index].latlng.lng);
        var pointList = [pointA, pointB];
       
        var polyline = new L.Polyline(pointList, {
            color: 'red',
            weight: 3,
            opacity: 0.5,
            smoothFactor: 1
        });
        this.polylineArray.push(polyline)
        polyline.addTo(this.map);
        this.polyLineIndex++    
      }
    this.index++
    }); // coordenadas do sítio que queremos mostrar no mapa
    
    const tiles = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      maxZoom:18,
      minZoom:4,
      attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    });

    tiles.addTo(this.map);

    // Get para ir buscar os camioes, para saber a localização deles
    this.leafletIconTruck = L.icon({
      iconUrl: 'assets/images/icon/truck_icon_red.png',
     // shadowUrl: 'leaf-shadow.png',
      iconSize:     [40, 60], // size of the icon
      shadowSize:   [40, 60], // size of the shadow
      iconAnchor:   [22, 50], // point of the icon which will correspond to marker's location
      shadowAnchor: [4, 62],  // the same for the shadow
      popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
    })
    this.leafletIconTruckGreen = L.icon({
      iconUrl: 'assets/images/icon/truck_icon_green.png',
     // shadowUrl: 'leaf-shadow.png',
      iconSize:     [40, 60], // size of the icon
      shadowSize:   [40, 60], // size of the shadow
      iconAnchor:   [22, 50], // point of the icon which will correspond to marker's location
      shadowAnchor: [4, 62],  // the same for the shadow
      popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
    })
    this.leafletIconTruckOrange = L.icon({
      iconUrl: 'assets/images/icon/truck_icon_orange.png',
     // shadowUrl: 'leaf-shadow.png',
      iconSize:     [40, 60], // size of the icon
      shadowSize:   [40, 60], // size of the shadow
      iconAnchor:   [22, 50], // point of the icon which will correspond to marker's location
      shadowAnchor: [4, 62],  // the same for the shadow
      popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
    })

    // Get para ir buscar os reboques, para saber a localização deles
    this.leafletIconTrailer = L.icon({
      iconUrl: 'assets/images/icon/trailer.png',
      iconSize:     [40, 60], // size of the icon
      shadowSize:   [40, 60], // size of the shadow
      iconAnchor:   [22, 50], // point of the icon which will correspond to marker's location
      shadowAnchor: [4, 62],  // the same for the shadow
      popupAnchor:  [-3, -76] // point from which the popup should open relative to the iconAnchor
    })

    this.loopForever = setInterval(() => {
      this.addTrucksAndTraillersStatus(1)
    }, 10000);

    this.addTrucksAndTraillersStatus(0)

  }

public addTrucksAndTraillersStatus(int){
  if (int == 1){
    this.markersArray.forEach((element, index) => {
      this.map.removeLayer(this.markersArray[index])
      console.warn(this.markersArray)
      console.warn(element)
      console.warn(index + " index ")
    });
  // this.map.removeLayer(this.markersArray[0])
  }
  this.trucksService.getAllStates().subscribe(async (data:any) =>{
   // marker = null;
    for(this.trucksMarks of data.data){
      const lat = this.trucksMarks.latitude
      const lon = this.trucksMarks.longitude
      const id = this.trucksMarks.truck_id
      if(this.trucksMarks.rpm == 0){
        var marker = L.marker([lat,lon],{icon: this.leafletIconTruck })
        this.map.addLayer(marker);
        this.markersArray.push(marker)
      }else if (this.trucksMarks.rpm >0 && this.trucksMarks.rpm < 1000){
        var marker = L.marker([lat,lon],{icon: this.leafletIconTruckOrange })
        this.map.addLayer(marker);
        this.markersArray.push(marker)
      }else {
        var marker = L.marker([lat,lon],{icon: this.leafletIconTruckGreen })
        this.map.addLayer(marker);
        this.markersArray.push(marker)
      }
      marker.on('click', async (e) => {
        this.clearHistoryTruckRoutes() // limpamos as rotas anteriores 
        this.clearMap() // limpo as linhas anteriores pq já não precisamos delas
        //this.map.removeLayer(marker)
        // Pega na rota 
        this.mapService.listTruckAndDriver().subscribe(
          data => {
            this.arrayTruckAndDriver= data.data
            const auxTruckAndDriver = this.arrayTruckAndDriver.find(x => x.truck_id == id )
            // vamos buscar o user para saber-mos o email dele 
            this.userService.getData().subscribe(
              data =>{
                this.user = data
                const userAux = this.user.data.find(x => x.ID == auxTruckAndDriver.first_driver_id )
                this.trucksService.getRoutes(userAux.email).subscribe(
                  data =>{
                    var array = JSON.parse("[" + data.coords + "]");
                    for (let index = 0; index < array[0].length; index= index + 2) {
                      //
                        if( index >= 3){
                          var pointA = new L.LatLng(array[0][index+0], array[0][index+1]);
                          if(index+3 < array[0].length ){
                            var pointB = new L.LatLng(array[0][index+2], array[0][index+3]);
                          }
                          var pointList = [pointA, pointB];
                        
                          var polyline = new L.Polyline(pointList, {
                              color: 'red',
                              weight: 5,
                              opacity: 1,
                              smoothFactor: 1
                          });
                            polyline.addTo(this.map); // traça as linhas no mapa
                        
                        }
                      //
                    }
                  }
                )
              }
            )
          }
        )
        await this.delay(1000); // delay para nao fazerem sincronos e esperar que o this.mapService.distance receba os valores
        this.getTruckHistory(id) // vai buscar o historico do camiao ( por onde passou ) e desenha no mapa
      })
      marker.bindPopup(
        "<b>Matrícula:</b>" + this.trucksMarks.plate + "<br></br>" + 
        "<b>Velocidade:</b>" + this.trucksMarks.speed  + "<br></br>" +
        "<b>Rpm:</b>" + this.trucksMarks.rpm + "<br></br>"
      )
    }
  })
  // Trailers
  this.trailerService.getData().subscribe((data:any) =>{
    for(this.trailersMarks of data.data){
      const lat = this.trailersMarks.latitude
      const lon = this.trailersMarks.longitude

      var marker = L.marker([lat,lon],{icon: this.leafletIconTrailer})
      this.map.addLayer(marker);
      this.markersArray.push(marker)
      marker.bindPopup(this.trailersMarks.plate)
    }
  })
}   

public getTruckHistory(id){
  // Desenhar as rotas do hitórico de rotas do camiao, para isso tenho de fazer um pedido à api que me retorne o array de localizações antigas do camiao
  this.trucksService.getTruckHistory(id).subscribe(
    data =>{
      this.distanceHistoryAux = 0
      this.truckCoordsHistory = data.data
      this.truckCoordsHistoryAux = []
      if(data.data.length > 0){ 
        this.truckCoordsHistory.forEach( (currentValue, index) => {
          this.truckCoordsHistoryAux.push(this.truckCoordsHistory[index].longitude+","+this.truckCoordsHistory[index].latitude)
        })
        let cycle;
        cycle = (this.truckCoordsHistoryAux.length) / 24
        if(Number.isInteger(cycle)==false){
          cycle++
        }
        console.log(cycle)
        console.log(this.truckCoordsHistoryAux[0]);
        let auxTamArray = this.truckCoordsHistoryAux.length
        let auxTamArrayIndex = 0
        this.wayPoints = 0
        for(let i = 1; i <= cycle; i++){
          if(this.truckCoordsHistoryAux.length <= 24){
            // CASO O HISTÓRICO TENHA MENOS DE 24 PONTOS 
              console.log(id) 
              console.log("1 if")
              var coords = this.truckCoordsHistoryAux.map(function (item) {
                return item;
              }).join(';'); //separação dos indices do array passa de , para ;
              this.wayPoints = "0"+";"+(this.truckCoordsHistoryAux.length-1)
              this.apiMapBox(coords,this.wayPoints, true, id);
          }
          else{
              if(auxTamArray > 24){ // Como tem mais de 24 pontos, temos de pegar em um array aux que vai ficar com as posições 
                // ex 1x - temos 50 pontos
                // array aux pega nas posicoes de 0-24 (temos uma variavél index que vai sendo incrementada, para ir acompanhando as posicoes que o array aux vai ficar)
                // no fim da condiçao temos de tirar 24 ao tamanho do "array" (auxTamArray isto porque ja temos as primeiras 24 posicoes) (50 -24 = 26) e enviar para a map box 
                // 2x como a variavel aux foi incrementada, o array aux já pega nas posicoes de 24-48 e voltamos a decrementar aux tam( 26-24 = 2) e enviar para a map box 
                // como depois o aux tam tem tamanho 2 e 2 < 24 saimos desta condição e vamos a condiçao "final"
                let auxArray = []
                let auxAIndex = auxTamArrayIndex
                // Fazer ciclo for do auxTamArray até auxAuxTamArray + 24 (aux tam array vai sempre acompanhando )
                for(auxTamArrayIndex; auxTamArrayIndex < (auxAIndex+24); auxTamArrayIndex++){
                  auxArray.push(this.truckCoordsHistoryAux[auxTamArrayIndex])
                }
                var coords = auxArray.map(function (item) {
                  return item;
                }).join(';'); //separação dos indices do array passa de , para ;
                this.wayPoints = "0"+";"+23 // Vai andar sempre 24 de cada vez
                this.apiMapBox(coords,this.wayPoints, false, id); // false pq nao vai entrar numa condicao do apiMapBox (por nao ser o calculo final)
                auxTamArray = auxTamArray - 24
              }else{ // Esta coindição é quando nos encontramos na fase final
                // aux tamIndex chega aqui com o valor 48 e o aux tam com valor 2 (<24 por isso viemos para aqui)
                // vamos percorrer de auxTamIndex ao tamanho do array original (50)
                // vai colocar no array aux as 2 posicoes (50-2) e vai enviar para o map box
                let auxArray = []
                let auxAIndex = auxTamArrayIndex
                // Fazer ciclo for do auxTamArray até auxAuxTamArray + 24 (aux tam array vai sempre acompanhando )
                for(auxTamArrayIndex; auxTamArrayIndex <= (this.truckCoordsHistoryAux.length-1); auxTamArrayIndex++){
                  auxArray.push(this.truckCoordsHistoryAux[auxTamArrayIndex])
                  }
                var coords = auxArray.map(function (item) {
                  return item;
                }).join(';'); //separação dos indices do array passa de , para ;
                this.wayPoints = "0"+";"+((this.truckCoordsHistoryAux.length - auxAIndex)-1) // tamanha do array (pq auxTamArray já chegou ao fim do array) - posição do index no inicio da 3 condição 
                this.apiMapBox(coords,this.wayPoints, true, id);
                auxTamArray = 0;
              }
          }
        }  
      }        
    })
    // Buscar a rota que o camionista criou e desenhar em Vermelho, para ser possível vizualizar, não apenas o que o camionista percorreu mas o que ainda falta percorrer
    // Para tar temos de pedir ao golang que aceda à BD de displacements aonde tem as rotas dos chefes de tráfego 
    // De seguida desenhamos no mapa
  }

  private apiMapBox(coords,wayPoints, boolean, id){

    this.mapService.generateRoute(coords.toString(),wayPoints).subscribe( // Enviar as cordenadas para a api (conforme a api pede) e o número de "vertíces"
    data => {
      console.warn(data)
      let mapBoxCoordsInters = []
      this.distanceHistoryAux += (data.routes[0].distance)/1000 //a distancia que nos dão é em metros entao tenho de converter para KM
      if(boolean){
        this.mapService.distanceAndIdTruckHistory.push({
           truck_id:id,
           km:this.distanceHistoryAux
        })
        console.log(this.mapService.distanceAndIdTruckHistory.toString)
      }

      const steps = data.routes[0].legs[0].steps // para aceder ao campo das localizações
      steps.forEach( (currentValue, index) => {
        steps[index].intersections.forEach((currentValue, indexAux) => {
          mapBoxCoordsInters.push(steps[index].intersections[indexAux])
        })
      })

      var arrayLocation = [] // vai guardar o array de location para depois converter para json e enviar para a BD
      mapBoxCoordsInters.forEach((currentValue, index) => {
        arrayLocation.push(mapBoxCoordsInters[index].location[1], mapBoxCoordsInters[index].location[0])
        if(index > 0)
          {
            var pointA = new L.LatLng(mapBoxCoordsInters[index-1].location[1], mapBoxCoordsInters[index-1].location[0]);
            var pointB = new L.LatLng(mapBoxCoordsInters[index].location[1], mapBoxCoordsInters[index].location[0]);
            var pointList = [pointA, pointB];
        
            var polyline = new L.Polyline(pointList, {
                color: 'blue',
                weight: 5,
                opacity: 1,
                smoothFactor: 1
            });
            polyline.addTo(this.map); // traça as linhas no mapa
        }
      })
    })
  }

  deleteCoords(){
    this.index--
    this.polyLineIndex--
    this.map.removeLayer(this.polylineArray[this.polyLineIndex]);
    
    this.polylineArray.splice(this.polylineArray[this.polyLineIndex],1)
    this.arrayCoords.splice(this.index, 1)

  }

  listTrucks(){
    this.mapService.listTrucksDrivers().subscribe(
      data => {
        this.arrayTruckDrivers= data.truckDriver
      }
    )
    this.mapService.listTrailers().subscribe(
      data => {
        this.arrayTrailers= data.trailer
      }
    )
    this.mapService.listTrucks().subscribe(
      data => {
        this.arrayTrucks= data.data
      }
    )
    this.mapService.listTruckAndDriver().subscribe(
      data => {
        this.arrayTruckAndDriver= data.data
      }
    )

  }

  addCoordTruck(plate,id, trailerLat, trailerLong){ // ADICIONA A ROTA CRIADA A UM CAMIONISTA (CAMIAO) E A UM REBOQUE

    // temos de limpar o array com as informaçoes das localizaçoes
    this.mapService.locationData = []
    this.mapService.time = 0
    const auxTruckAndDriver = this.arrayTruckAndDriver.find(x => x.first_driver_id == id )
    const truck = this.arrayTrucks.find(x => x.ID == auxTruckAndDriver.truck_id)

    this.arrayCoordsAux.push(truck.longitude+","+truck.latitude) // adiciono os pontos da localização do camiao
    this.arrayCoordsAux.push(trailerLong+","+trailerLat) // adiciono os pontos da localização do reboque (para fazer o percurso entre o camiao e reboque selecionados e só de seguida a rota que defenimos)
    this.arrayCoords.forEach( (currentValue, index) => {
      this.arrayCoordsAux.push(this.arrayCoords[index].latlng.lng+","+this.arrayCoords[index].latlng.lat) 
    })

    console.warn(this.arrayCoordsAux)
    let initRoute = this.arrayCoordsAux[0] // para saber as coordenadas do primeiro ponto da rota (localização do reboque)
    let lastRoute = this.arrayCoordsAux[(this.arrayCoordsAux.length)-1] // para saber as coordenadas do ultimo ponto da rota
    this.mapService.getLocations(initRoute).subscribe(
      data =>{
        this.dataLocation = data
        let country
        let postalCode
        let address
        if(this.dataLocation.features[0].context[0].id.includes("neighborhood")){
            postalCode = this.dataLocation.features[0].context[1].text
            address = this.dataLocation.features[0].context[2].text
        }else{
          postalCode = this.dataLocation.features[0].context[0].text
          address = this.dataLocation.features[0].context[1].text
        }
        if(this.dataLocation.features[0].context[3].id.includes("region")){ // se incluir a regiao excluo
          country = this.dataLocation.features[0].context[4].text
        }else{
          country = this.dataLocation.features[0].context[3].text
        }
        this.mapService.locationData.push(
          country,
          this.dataLocation.features[0].context[2].text,
          postalCode,
          address
        )
      }
    )
    this.mapService.getLocations(lastRoute).subscribe(
      data =>{
        this.dataLocation = data
        let country
        let postalCode
        let address
        if(this.dataLocation.features[0].context[0].id.includes("neighborhood")){
            postalCode = this.dataLocation.features[0].context[1].text
            address = this.dataLocation.features[0].context[2].text
        }else{
          postalCode = this.dataLocation.features[0].context[0].text
          address = this.dataLocation.features[0].context[1].text
        }
        if(this.dataLocation.features[0].context[3].id.includes("region")){ // se incluir a regiao excluo
          country = this.dataLocation.features[0].context[4].text
        }else{
          country = this.dataLocation.features[0].context[3].text
        }
        this.mapService.locationData.push(
          country,
          this.dataLocation.features[0].context[2].text,
          postalCode,
          address
        )
      }
    )

    var coords = this.arrayCoordsAux.map(function (item) {
        return item;
    }).join(';'); //separação dos indices do array passa de , para ;  
    console.warn(coords)
    let wayPoints = "0"+";"+(this.index+1) // supostamente devia ser -1 mas como estamos a adicionar as rotas entre o camiao e o reboque, sao +2 rotas, logo tem de ser index+1 ( (index-1) +2 = (index+1) )
    this.mapService.generateRoute(coords.toString(),wayPoints).subscribe( // Enviar as cordenadas para a api (conforme a api pede) e o número de "vertíces"
      data => {
        let seconds = (data.routes[0].duration)
        let days = Math.floor(seconds / 86400);
        seconds -= days * 86400;
        let hours = Math.floor(seconds / 3600) % 24;
        seconds -= hours * 3600;
        var minutes = Math.floor(seconds / 60) % 60;

        this.time = days + 'd ' + hours + 'h ' + minutes + 'm '
        this.mapService.time = this.time
       
        this.distance = (data.routes[0].distance)/1000 //a distancia que nos dão é em metros entao tenho de converter para KM
        console.log(this.distance)
        let mapBoxCoordsInters = []
        this.clearMap() // limpo as linhas anteriores pq já não precisamos delas
        const steps = data.routes[0].legs[0].steps // para aceder ao campo das localizações
        steps.forEach( (currentValue, index) => {
          steps[index].intersections.forEach((currentValue, indexAux) => {
            mapBoxCoordsInters.push(steps[index].intersections[indexAux])
          })
        })

        var arrayLocation = [] // vai guardar o array de location para depois converter para json e enviar para a BD
        mapBoxCoordsInters.forEach((currentValue, index) => {
          arrayLocation.push(mapBoxCoordsInters[index].location[1], mapBoxCoordsInters[index].location[0])
          if(index > 0)
            {
              var pointA = new L.LatLng(mapBoxCoordsInters[index-1].location[1], mapBoxCoordsInters[index-1].location[0]);
              var pointB = new L.LatLng(mapBoxCoordsInters[index].location[1], mapBoxCoordsInters[index].location[0]);
              var pointList = [pointA, pointB];
          
              var polyline = new L.Polyline(pointList, {
                  color: 'red',
                  weight: 3,
                  opacity: 0.5,
                  smoothFactor: 1
              });
              polyline.addTo(this.map); // traça as linhas no mapa
          }
        })
        var myJsonString = JSON.stringify(arrayLocation);
        this.mapService.createRoute(myJsonString,id,plate, this.distance).subscribe(
          data => {

          }
        )
      }, 
    );
  }

  clearHistoryTruckRoutes(){
    let i : any
    for(i in this.map._layers) {
        if(this.map._layers[i]._path != undefined) {
            try {
                this.map.removeLayer(this.map._layers[i]);
            }
            catch(e) {
                console.log("problem with " + e + this.map._layers[i]);
            }
        }
    }

    this.truckCoordsHistory = []
    this.truckCoordsHistoryAux = []

  }

  clearMap() {
    let i : any
    for(i in this.map._layers) {
        if(this.map._layers[i]._path != undefined) {
            try {
                this.map.removeLayer(this.map._layers[i]);
            }
            catch(e) {
                console.log("problem with " + e + this.map._layers[i]);
            }
        }
    }
    this.arrayCoords = [] // Limpamos o array de coords
    this.arrayCoordsAux = []
    // Zeramos os index dos arrays
    this.index = 0
    this.polyLineIndex = 0;
  }

  delay(ms: number) {
    return new Promise( resolve => setTimeout(resolve, ms) );
  }

  ngOnDestroy() {
    if (this.loopForever) {
      clearInterval(this.loopForever);
    }
  }

}
