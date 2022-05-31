import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { TrucksComponent } from './trucks/trucks.component';

const routes: Routes = [
  {
    path: '',
    component: TrucksComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TrucksRoutingModule {}
