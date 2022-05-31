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
export class LocationService {
  constructor(private http: HttpClient) { }

  getData():Observable<any>{
    return this.http.get(API_URL + 'locations', httpOptions)
  }

  addLocation(data): Observable<any> {
    return this.http.post(API_URL + 'locations/addLocation', {
      name: data.name,
      latitude: data.latitude,
      longitude: data.longitude,
    }, httpOptions);
  }

  deleteLocation(id):Observable<any> {
    return this.http.delete(API_URL + 'locations/deleteLocation/' + id, httpOptions);
  }

  getLocation(id):Observable<any> {
    return this.http.get(API_URL + 'locations/getLocation/' + id, httpOptions);
  }

  updateLocation(data): Observable<any> {
    return this.http.post(API_URL + 'locations/updateLocation', {
      id: data.id_editLocation,
      name: data.name_edit,
      latitude: data.latitude_edit,
      longitude: data.longitude_edit,
    }, httpOptions);
  }

}
