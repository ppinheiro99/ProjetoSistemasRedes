import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";

//const API_URL = 'http://18.130.231.194:8080/api/';
const API_URL = 'http://localhost:80/api/';

import {Observable} from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})

export class TrucksService {

  email : string

  constructor(private http: HttpClient) {
  }

  getData():Observable<any>{
    return this.http.get(API_URL + 'trucks', httpOptions)
  }

  getCountTruck():Observable<any>{
    return this.http.get(API_URL + 'trucks/truckCount', httpOptions)
  }

  getAllStates():Observable<any>{
    return this.http.get(API_URL + 'trucks/alltruckState', httpOptions)
  }

  getRpmPlates():Observable<any>{
    return this.http.get(API_URL + 'trucksRpmPlates', httpOptions)
  }

  getTruckHistory(id):Observable<any>{
    return this.http.get(API_URL + 'trucks/getTruckHistory/' + id, httpOptions)
  }

  addTruck(data): Observable<any> {
    return this.http.post(API_URL + 'trucks/register', {
      plate: data.plate,
      year: data.year,
      month: data.month,
      km: data.km,
      brand: data.brand,
    }, httpOptions);
  }

  deleteTruck(id):Observable<any> {
    return this.http.delete(API_URL + 'trucks/deleteTruck/' + id, httpOptions);
  }

  getTruck(id):Observable<any> {
    return this.http.get(API_URL + 'trucks/getTruck/' + id, httpOptions);
  }


  updateTruck(data): Observable<any> {
    return this.http.post(API_URL + 'trucks/updateTruck', {
      id: data.id_editTruck,
      plate: data.plate_editTruck,
      year: data.year_editTruck,
      month: data.month_editTruck,
      km: data.km_editTruck,
      brand: data.brand_editTruck,
    }, httpOptions);
  }

  bindTruckAndDriver(DriverId,TruckId): Observable<any> {
    return this.http.post(API_URL + 'trucks/bindTruckAndDriver', {
      driver_id: DriverId,
      truck_id: TruckId,
    }, httpOptions);
  }

  createRoute(data,idCamiao): Observable<any> {
    return this.http.post(API_URL + 'trucks/addRoute', {
      truck_id: idCamiao,
      latitude: data.latitudeRoute,
      longitude: data.longitudeRoute,
    }, httpOptions);
  }
  getTruckDriver(id):Observable<any> {
    return this.http.get(API_URL + 'trucks/getTruckDriver/' + id, httpOptions);
  }

  unbindTruckDriver(id):Observable<any> {
    return this.http.delete(API_URL + 'trucks/unbindTruckDriver/' + id, httpOptions);
  }

  getRoutes(email):Observable<any>{
    return this.http.get(API_URL + 'mobile/getDriverRoutes/' + email, httpOptions);
  }

}

