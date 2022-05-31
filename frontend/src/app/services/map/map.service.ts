import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

//const API_URL = 'http://18.130.231.194:8080/api/';
const API_URL = 'http://localhost:80/api/';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})

export class MapService {
    public distanceAndIdTruckHistory:{
      truck_id : any,
      km:any
    }[]= []

    public locationData:{
      country:any,
      city: any,
      postalCode: any,
      address:any,
    }[] =[]
    
    public distance = 0
    public time: any
    
    truckCoordsHistory = []
    truckCoordsHistoryAux = []
    distanceHistoryAux: any
    index = 0
    wayPoints:any

    constructor(private http: HttpClient) { }


    generateRoute(coords,wayPoints): Observable<any>{
      return this.http.get("https://api.mapbox.com/directions/v5/mapbox/driving/"+coords+"?waypoints="+wayPoints+"&steps=true&voice_instructions=true&banner_instructions=true&voice_units=imperial&roundabout_exits=true&access_token=pk.eyJ1IjoicHBpbmhlaXJvOTkiLCJhIjoiY2tvbXI3N3NuMXIzejJxcm1uZTJxOWlmYyJ9.kJ7suLcfyHd8DBr34XygRw", httpOptions);
    }

    getLocations(coords){
      return this.http.get("https://api.mapbox.com/geocoding/v5/mapbox.places/"+coords+".json?types=address&access_token=pk.eyJ1IjoicHBpbmhlaXJvOTkiLCJhIjoiY2tvbXI3N3NuMXIzejJxcm1uZTJxOWlmYyJ9.kJ7suLcfyHd8DBr34XygRw")
    }

    createRoute(jsonCoords, id, plate, distance): Observable<any>{
      console.warn(this.locationData)
      console.warn(this.time)
      return this.http.post(API_URL + 'createDisplacement', {
       coords: jsonCoords,
       truck_driver:id,
       trailer_plate: plate,
       distance: distance,
       time: this.time,
       start_country:this.locationData[0],
       start_city: this.locationData[1],
       start_postal_code: this.locationData[2],
       start_address:this.locationData[3],
       end_country:this.locationData[4],
       end_city: this.locationData[5],
       end_postal_code: this.locationData[6],
       end_address:this.locationData[7],
      }, httpOptions);
    }

    listTrucksDrivers(): Observable<any>{
      return this.http.get(API_URL +"listTrucksDrivers", httpOptions);
    }
    listTrailers(): Observable<any>{
      return this.http.get(API_URL +"listTrailers", httpOptions);
    }
    listTrucks(): Observable<any>{
      return this.http.get(API_URL +"trucks", httpOptions);
    }
    listTruckAndDriver(): Observable<any>{
      return this.http.get(API_URL + 'listTruckAndDriver', httpOptions)
    }

    // Vai pegar no histórico do camionista e nos KM que ele percorreu
    getTruckHistory(data,id){
      this.index = 0
      this.distance = 0
      this.wayPoints = 0
      this.truckCoordsHistoryAux=[]
      this.distanceHistoryAux = 0
      console.log(data.data)
      if(data.data.length > 0){ 
        this.truckCoordsHistory = data.data
        this.truckCoordsHistory.forEach( (currentValue, index) => {
          this.truckCoordsHistoryAux.push(this.truckCoordsHistory[index].longitude+","+this.truckCoordsHistory[index].latitude)
        })
        let cycle;
        cycle = (this.truckCoordsHistoryAux.length) / 24
        if(Number.isInteger(cycle)==false){
          cycle++
        }
        console.warn(cycle + " tam " + this.truckCoordsHistoryAux.length)
        let auxTamArray = this.truckCoordsHistoryAux.length
        let auxTamArrayIndex = 0
        for(let i = 1; i <= cycle; i++){
          if(this.truckCoordsHistoryAux.length <= 24){
            // CASO O HISTÓRICO TENHA MENOS DE 24 PONTOS 
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
      return 0
    }

    private apiMapBox(coords,wayPoints, boolean, id){
      this.generateRoute(coords.toString(),wayPoints).subscribe( // Enviar as cordenadas para a api (conforme a api pede) e o número de "vertíces"
      data => {
        console.log(data)
        this.distanceHistoryAux += (data.routes[0].distance)/1000 //a distancia que nos dão é em metros entao tenho de converter para KM
        if(boolean){
          this.distance =this.distanceHistoryAux
          console.log(this.distance)
        }
      })
    }

    //
    private getDisplacements(truckID){

    }

}