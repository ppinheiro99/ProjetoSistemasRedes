import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { InBoxComponent } from './inBox-form/in-box.component';

const routes: Routes = [
  {
    path: '',
    component: InBoxComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class InboxRoutingModule {}
