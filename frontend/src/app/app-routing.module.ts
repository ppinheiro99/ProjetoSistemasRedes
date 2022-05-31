import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from './core/auth.guard';
import { PageNotFoundComponent } from './shared/page-not-found/page-not-found.component';
import { PassrecoverComponent } from './components/passrecover/passrecover.component';
import { PassrecoverFormComponent } from './components/passrecover-form/passrecover-form.component';

const routes: Routes = [
  {
    path: '',
    loadChildren: () =>
      import('./components/admin/admin.module').then((m) => m.AdminModule),
    canActivate: [AuthGuard],
  },
  {
    path:'passrecover',component: PassrecoverComponent
  },
  {
    path:'passrecover/:token',component: PassrecoverFormComponent
  },
  {
    path: 'login',
    loadChildren: () =>
      import('./components/login/login.module').then((m) => m.LoginModule),
  },
  {
    path: '**',
    component: PageNotFoundComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
