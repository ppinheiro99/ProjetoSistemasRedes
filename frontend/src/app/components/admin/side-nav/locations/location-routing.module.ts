import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ManageLocationComponent } from './manage-locations/manage-locations.component';

const routes: Routes = [
  {
    path: '',
    component: ManageLocationComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class LocationRoutingModule {}
