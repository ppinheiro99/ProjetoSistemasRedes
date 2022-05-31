import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { EditPassFormComponent } from './edit-password-form/edit-pass-form.component';

const routes: Routes = [
  {
    path: '',
    component: EditPassFormComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class EditPassRoutingModule {}
