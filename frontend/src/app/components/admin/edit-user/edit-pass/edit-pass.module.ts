import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EditPassFormComponent } from './edit-password-form/edit-pass-form.component';
import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { EditPassRoutingModule } from "./edit-pass-routing.module";
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatDialogModule } from '@angular/material/dialog';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
    declarations: [EditPassFormComponent],
    imports: [
      CommonModule,
      MatInputModule,
      MatTabsModule,
      MatCardModule,
      EditPassRoutingModule,
      MatButtonModule,
      MatCheckboxModule,
      MatFormFieldModule,
      MatToolbarModule,
      MatSelectModule,
      FormsModule,
      ReactiveFormsModule,
      HttpClientModule,
      MatDialogModule,
    ]
  })
  export class EditPassModule { }