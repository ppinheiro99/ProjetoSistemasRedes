import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Rx';
import { io, Socket } from 'socket.io-client';
import { TokenService } from "../token/token.service";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import {Observable} from 'rxjs';
import { UsersService } from '../user/users.service';

const API_URL = 'http://localhost:80/api/';
//const API_URL = 'http://18.130.231.194:8080/api/';
const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  })
};

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  receiver_id : any
  socket: any
  user_id:any
  userName:any
  messageList: {message: string, userName: string, senderID: string, userId : string, mine: boolean}[] = []
  users_connected = []
  users: {userID:number, userName:string, online:boolean}[]= []
  userList = []
  messages : {message: string, userName: string, senderID: string, userId : string, mine: boolean}[] = []
  user_select: any

  constructor(private tokenService:TokenService, private http: HttpClient, private userService: UsersService) {
   
   }

  init(){
    const user = this.tokenService.getUser()
    this.user_id = user.ID
    this.userName = user.Name
    //this.socket = io(`http://18.130.231.194:3000?userName=${this.userName}&id=${this.user_id}`)
    this.socket = io(`http://localhost:3000?userName=${this.userName}&id=${this.user_id}`)
    
    this.socket.on('user connect', (userList) =>{
    })

     this.socket.on('user-list', (userList) =>{
        this.userList = userList
        this.userService.getData().subscribe(data =>{

        // quando recebo um user-list é pq alguem novo se conectou e como tal vou receber o array de novos users, tendo entao que limpar o antigo
        this.users = []  
        data.data.forEach( (currentValue, index) => { // para saber dos users todos quais estao online e quais estao offline
          if(userList.find(x => x.id == data.data[index].ID ) != null){ 
            this.users.push({
              userID: data.data[index].ID,
              userName: data.data[index].first_name,
              online: true
            })
          }else{
            this.users.push({
              userID: data.data[index].ID,
              userName: data.data[index].first_name,
              online: false
            })
          }
        })
      })
    })

    this.socket.on("message", (users) =>{
      this.messageList.push({
          message:users.message,
          userName:users.userName,
          senderID:users.sender,
          userId: users.received,
          mine:users.mine
      })
        let aux = 0
        this.messageList.forEach( (currentValue, index) => {
        if(this.messageList[index].senderID == this.user_select  || this.messageList[index].userId == this.user_select){
          this.messages[aux] = this.messageList[index]
          aux++
        }
      })
    })
  }

  getMessagesByUser(id){
     return this.http.get(API_URL + 'getMessagesByUser/' + id, httpOptions);
  }

  sendMessage(message,receiver_id, sender_id):Observable<any>{
      this.socket.emit('message', { 
        data:message,
        receivedID: receiver_id
      })
      return this.http.post(API_URL + 'messages', {
        message: message,
        sender_id: sender_id,
        receiver_id: receiver_id
      }, httpOptions);
  }
}