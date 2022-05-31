import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegisterFormComponent } from './register-form/register-form.component';
import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { RegisterRoutingModule } from "./register-routing.module";
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { MatSelectCountryModule } from '@angular-material-extensions/select-country';
import { MarkdownModule } from 'ngx-markdown';
import { HttpClientModule } from '@angular/common/http';
import { authInterceptorProviders } from 'src/app/helpers/auth.interceptor';

@NgModule({
  declarations: [RegisterFormComponent],
  imports: [
    CommonModule,
    MatInputModule,
    MatTabsModule,
    MatCardModule,
    RegisterRoutingModule,
    MatButtonModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatToolbarModule,
    MatSelectModule,
    FormsModule,
    MarkdownModule.forRoot(),
    HttpClientModule,
    MatSelectCountryModule.forRoot('pt'), // you can use 'br' | 'de' | 'en' | 'es' | 'fr' | 'hr' | 'it' | 'nl' | 'pt' --> MatSelectCountrySupportedLanguages
    ReactiveFormsModule,
    MatDialogModule,
  ],
  providers: [authInterceptorProviders]
})
export class RegisterModule { }
