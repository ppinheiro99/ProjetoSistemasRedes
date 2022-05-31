import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { FlexLayoutModule } from '@angular/flex-layout';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CoreModule } from './core/core.module';
import { SharedModule } from './shared/shared.module';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { authInterceptorProviders } from './helpers/auth.interceptor';
import { CommonModule } from '@angular/common';
import { PassrecoverComponent } from './components/passrecover/passrecover.component';
import { PassrecoverFormComponent } from './components/passrecover-form/passrecover-form.component';
import { MatSelectCountryModule } from '@angular-material-extensions/select-country';


@NgModule({
  declarations: [
    AppComponent,
    PassrecoverComponent,
    PassrecoverFormComponent
  ],
  imports: [
    BrowserModule,
    CommonModule,
    MatButtonModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatSelectCountryModule.forRoot('de'), // you can use 'br' | 'de' | 'en' | 'es' | 'fr' | 'hr' | 'it' | 'nl' | 'pt' --> MatSelectCountrySupportedLanguages
    BrowserAnimationsModule,
    ReactiveFormsModule,
    FlexLayoutModule,
    MatDialogModule,
    AppRoutingModule,
    HttpClientModule,
    CoreModule,
    SharedModule
  ],
  providers: [authInterceptorProviders],
  bootstrap: [AppComponent]
})
export class AppModule {}
