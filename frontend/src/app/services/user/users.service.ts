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

export class UsersService {
  constructor(private http: HttpClient) {}

  getData():Observable<any>{
    return this.http.get(API_URL + 'users', httpOptions)
  }

  getTruckDrivers():Observable<any>{
    return this.http.get(API_URL + 'usersDrivers', httpOptions)
  }

  deleteUser(id):Observable<any> {
    return this.http.delete(API_URL + 'deleteUser/' + id, httpOptions);
  }

  updateUser(user, email): Observable<any> {
    return this.http.post(API_URL + 'updateUser', {
      email: email,
      first_name: user.first_name,
      last_name: user.last_name,
      address: user.address,
      country:user.country.name,
    }, httpOptions);
  }

  sendMsg(sender_id,id, message):Observable<any> {
    return this.http.post(API_URL + 'sendMsg/' +id, {
      message:message,
      sender_id: sender_id,
    }, httpOptions);
  }

  receivedMessage(id):Observable<any>{
    return this.http.get(API_URL + 'inbox/' +id , httpOptions);
  }

  getTravelMap(id){
    return this.http.get(API_URL +"travelMap/"+ id, httpOptions);
  }
}
