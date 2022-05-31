import { Component, OnInit } from '@angular/core';
import { UsersService } from "../../../../services/user/users.service";
import { TokenService } from "../../../../services/token/token.service";
@Component({
  selector: 'app-in-box',
  templateUrl: './in-box.component.html',
  styleUrls: ['./in-box.component.scss']
})
export class InBoxComponent implements OnInit {
  id:any
  msg:any
  senderName:any
  senderLastName:any

  constructor(private userService : UsersService, private tokenService: TokenService) { }

  ngOnInit(): void {
    const user = this.tokenService.getUser()
    this.id = user.ID
    console.log("tem de escrever a mensagem aqui em baixo:")
    this.userService.receivedMessage(this.id).subscribe(
      data => { // ele devia retornar para aqui como nos fazemos sempre, mas nao esta a dar
        console.log(data)
        this.msg = data.message
        this.senderName = data.senderName
        this.senderLastName = data.senderLastName
        console.log("sucesso")
      },
    )
    console.log("tem de escrever a mensagem aqui em cima:")
  }

}
