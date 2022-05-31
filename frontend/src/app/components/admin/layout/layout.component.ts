import { Component, OnInit, OnDestroy} from '@angular/core';
import { Subscription } from 'rxjs';
import { UsersService } from 'src/app/services/user/users.service';
import { TokenService } from "../../../services/token/token.service";
import { ChatService } from "src/app/services/socket/chat.service";

export interface MessageData {
  id : any;
  message : any;
}

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit, OnDestroy { 
  displayedColumns = ['icon','message'];
  sideNavOpened = false;
  messageSideNavOpened = false;
  showFiller = false;
  select = 0
  showMessagesId:any
  sideNavMode: 'side' | 'over' = 'side';
  messageSideNavMode: 'side' | 'over' = 'side';
  toolBarHeight = 64;
  users : any[]
  user_id : any 
  userName = ''
  showUserMouseOver = true

  private readonly mediaWatcher: Subscription;
  constructor(private userService: UsersService, private tokenService: TokenService, public chat: ChatService) {
  }
  ngOnInit() { 
    this.userService.getData().subscribe(data =>{
      this.users = data.data
    })
    const user = this.tokenService.getUser()
    this.user_id = user.ID
    this.userName = user.Name
    // Carregamos as mensagens todas do user da BD
    this.chat.getMessagesByUser(this.user_id).subscribe(
      data =>{
        var auxMessages
        auxMessages = data
        if(auxMessages.messagesArray.length != 0){
          auxMessages.messagesArray.forEach( (currentValue, index) => {
          // Verifico quais sao as mensagens enviadas por mim e quais é que não são adicionando depois ao array de mensagens
          if(auxMessages.messagesArray[index].sender_id == this.user_id){
            this.chat.messageList.push({
              message: auxMessages.messagesArray[index].message,
              userName: this.userName,
              senderID:auxMessages.messagesArray[index].sender_id,
              userId:auxMessages.messagesArray[index].receiver_id,
              mine: true
            })
          }else{
              this.chat.messageList.push({
                message: auxMessages.messagesArray[index].message,
                userName: this.userName,
                senderID:auxMessages.messagesArray[index].sender_id,
                userId:auxMessages.messagesArray[index].receiver_id,
                mine: false
              })
            }
          })
        }
    })
    // Mostramos as mensagens (mandamos para o array de mensagens)
    this.chat.init()
    this.updateMessage(this.select)
  }

  ngOnDestroy(): void {
    this.mediaWatcher.unsubscribe();
  }

  toggleMessageBar(id){
    this.showMessagesId = id
    this.chat.receiver_id = id 
    this.showFiller = !this.showFiller
    this.chat.messages = []
  
    // percorrer as mensagens que tenho do user que carreguei (id)
    this.updateMessage(id)
  }

  updateMessage(id){
    let aux = 0
    this.chat.user_select = id
    this.chat.messageList.forEach( (currentValue, index) => {
      if(this.chat.messageList[index].senderID == id  || this.chat.messageList[index].userId == id){
        this.chat.messages[aux] = this.chat.messageList[index]
        aux++
      }
    });
  }

}
