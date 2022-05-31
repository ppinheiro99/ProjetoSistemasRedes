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

export class TrailersService {

  email : string

  constructor(private http: HttpClient) {
  }

  getData():Observable<any>{
    return this.http.get(API_URL + 'trailers', httpOptions)
  }

  addTrailer(data): Observable<any> {
    return this.http.post(API_URL + 'trailers/register', {
      plate: data.plateTrailer,
      year: data.yearTrailer,
    }, httpOptions);
  }

  deleteTrailer(id):Observable<any> {
    return this.http.delete(API_URL + 'trailers/deleteTrailer/'+ id, httpOptions);
  }

  getTrailer(id):Observable<any> {
    return this.http.get(API_URL + 'trailers/getTrailer/' + id, httpOptions);
  }

  updateTrailer(data): Observable<any> {
    return this.http.post(API_URL + 'trailers/updateTrailer', {
      id: data.id_editTrailer,
      plate: data.plate_editTrailer,
      year: data.year_editTrailer,
    }, httpOptions);
  }

}

