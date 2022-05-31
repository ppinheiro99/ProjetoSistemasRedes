import { Component, OnInit } from '@angular/core';
import { superAdminRoutes } from '../../layout/routes/super-admin-routes';
import { adminRoutes } from '../../layout/routes/admin-routes';
import { chefeTrafegoRoutes } from '../../layout/routes/chefe-trafego-routes';
import { TokenService } from "../../../../services/token/token.service";

@Component({
  selector: 'app-side-nav',
  templateUrl: './side-nav.component.html',
  styleUrls: ['./side-nav.component.scss']
})
export class SideNavComponent implements OnInit {
  showMenu = false;
  routes :any;
  user : any;
  constructor(private tokenService: TokenService) {}

  ngOnInit() {
    this.routes = this.getRoutes()
  }

  getRoutes(){
    this.user = this.tokenService.getUser()
    if(this.user.Role == 1){
      this.routes = superAdminRoutes
      return this.routes
    }else if(this.user.Role == 2){
      this.routes = adminRoutes
      return this.routes
    }else if(this.user.Role == 3){
      this.routes = chefeTrafegoRoutes
      return this.routes
    }
  }
}
