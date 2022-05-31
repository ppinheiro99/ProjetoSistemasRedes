import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { InboxRoutingModule } from "./inBox-routing.module";
import { InBoxComponent } from "./inBox-form/in-box.component";
@NgModule({
  imports: [
    CommonModule,
    InboxRoutingModule,
    FlexLayoutModule,
    MatCardModule,
    MatButtonModule
  ],
  declarations: [InBoxComponent]
})
export class InBoxModule {}
