import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LeafletModule } from "@asymmetrik/ngx-leaflet";
import { MapRoutingModule } from './map-routing.module';
import { HomeComponent } from './home/home.component';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { authInterceptorProviders } from 'src/app/helpers/auth.interceptor';
@NgModule({
  declarations: [HomeComponent],
  imports: [
    CommonModule,
    MapRoutingModule,
    LeafletModule,
    MatMenuModule,
    MatIconModule,
    MatSelectModule,
    FormsModule
  ],
  providers: [DatePipe,authInterceptorProviders]
})
export class MapModule { }
