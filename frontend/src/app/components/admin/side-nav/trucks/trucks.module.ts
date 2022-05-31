import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { TrucksRoutingModule } from './trucks-routing.module';
import { TrucksComponent } from './trucks/trucks.component';
import { authInterceptorProviders } from "../../../../helpers/auth.interceptor";
import { MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
@NgModule({
  imports: [
    CommonModule,
    TrucksRoutingModule,
    MatTableModule,
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatPaginatorModule,
    MatCardModule,
    MatDialogModule,
    MatSortModule,
    MatInputModule,
    MatButtonModule,
    MatCheckboxModule
  ],
  declarations: [TrucksComponent],
  providers: [authInterceptorProviders]
})
export class TrucksModule {}
