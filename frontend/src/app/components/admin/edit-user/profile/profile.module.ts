import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProfileFormComponent } from './profile-form/profile-form.component';
import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { ProfileRoutingModule } from "./profile-routing.module";
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { HttpClientModule } from '@angular/common/http';
import { MatSelectCountryModule } from '@angular-material-extensions/select-country';
import { MarkdownModule } from 'ngx-markdown';
import {MatButtonToggleModule} from '@angular/material/button-toggle';

@NgModule({
  declarations: [ProfileFormComponent],
  imports: [
    CommonModule,
    MatInputModule,
    MatTabsModule,
    MatCardModule,
    ProfileRoutingModule,
    MatButtonModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatToolbarModule,
    MatSelectModule,
    FormsModule,
    ReactiveFormsModule,
    MarkdownModule.forRoot(),
    HttpClientModule,
    MatButtonToggleModule,
    MatSelectCountryModule.forRoot('pt'), // you can use 'br' | 'de' | 'en' | 'es' | 'fr' | 'hr' | 'it' | 'nl' | 'pt' --> MatSelectCountrySupportedLanguages
    MatDialogModule,
  ]
})
export class ProfileModule { }
