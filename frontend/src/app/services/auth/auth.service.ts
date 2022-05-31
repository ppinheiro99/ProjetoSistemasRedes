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
export class AuthService {
  constructor(private http: HttpClient) { }

  login(credentials): Observable<any> {
    return this.http.post(API_URL + 'auth/login', {
      email: credentials.email,
      password: credentials.password
    }, httpOptions);
  }
  
  passRecover(credentials): Observable<any> {
    return this.http.post(API_URL + 'auth/passRecover', {
      email: credentials.email,
    }, httpOptions);
  }

  passRecoverVerififyToken(token): Observable<any> {
    return this.http.post(API_URL + 'auth/passRecover/checktoken', {
      token: token,
    }, httpOptions);
  }

  passRecoverSetPassword(credentials,token): Observable<any> {
    return this.http.post(API_URL + 'auth/passRecover/setPassword', {
      password: credentials.password,
      confirm_password: credentials.confirme_password,
      token: token,
    }, httpOptions);
  }


  register(user): Observable<any> {
    return this.http.post(API_URL + 'register', {
      email: user.email,
      password: user.password,
      role_id: parseInt(user.role),
      first_name: user.first_name,
      last_name: user.last_name,
      address: user.address,
      country:user.country.name,
    }, httpOptions);
  }

  changePassword(user,email): Observable<any> {
    
    return this.http.post(API_URL + 'auth/changePassword', {
      oldpassword: user.password_old,
      newpassword: user.password_new,
      newpasswordC: user.password_newConf,
      email: email,
    }, httpOptions);
  }
}
