import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { TokenService } from "../../../../services/token/token.service";
import { UsersService } from "../../../../services/user/users.service";

@Component({
  selector: 'app-top-nav',
  templateUrl: './top-nav.component.html',
  styleUrls: ['./top-nav.component.scss']
})
export class TopNavComponent implements OnInit {
  @Output() sideNavToggled = new EventEmitter<void>();
  @Output() messageNavToggled = new EventEmitter<void>();
  name: string
  id:string
  isLoggedIn = false;
  displayedColumns = ['icon','message'];

  constructor(private readonly router: Router, private tokenService: TokenService, private userService:UsersService) {}

  ngOnInit() {
    this.isLoggedIn = !!this.tokenService.getToken();
    if (this.isLoggedIn) {
      const user = this.tokenService.getUser();
      this.name = user.Name
      this.id = user.ID
    }
  }

  toggleSidebar() {
    this.sideNavToggled.emit()
  }

  toggleMessageBar(){
    this.messageNavToggled.emit()
  }

  onLoggedout() {
    localStorage.removeItem('isLoggedin')
    this.tokenService.signOut()
    this.router.navigate(['/login'])
    window.location.reload();
  }
}
