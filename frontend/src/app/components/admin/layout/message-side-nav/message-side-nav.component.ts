import { Component, OnInit } from '@angular/core';
import { UsersService } from "../../../../services/user/users.service";
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TokenService } from "../../../../services/token/token.service";
import { ChatService } from "src/app/services/socket/chat.service";

@Component({
  selector: 'app-message-side-nav',
  templateUrl: './message-side-nav.component.html',
  styleUrls: ['./message-side-nav.component.scss']
})
export class MessageSideNavComponent implements OnInit {
  users : any[]
  validationForm : FormGroup
  send_message = ''
  user_id : any
  receider_id: any

  constructor(private userService:UsersService, public fb: FormBuilder, private tokenService: TokenService, public chat: ChatService) {
    this.validationForm = fb.group({
      message:[null, Validators.required],
    });
   }

   get message() { return this.validationForm.get('message')}

  ngOnInit(): void {
    this.userService.getData().subscribe(data =>{
      this.users = data.data
    })
    const user = this.tokenService.getUser()
    this.user_id = user.ID
    this.receider_id = this.chat.receiver_id
  }

  onSubmit(){
    let message = this.validationForm.value.message
    let receiver_id = this.chat.receiver_id
    if(message != ""){
      this.chat.sendMessage(message,receiver_id, this.user_id).subscribe(
       )
      this.send_message = ''
    }
  }

}
