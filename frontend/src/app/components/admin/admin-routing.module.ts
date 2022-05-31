import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LayoutComponent } from './layout/layout.component';
import { adminRoutes } from './layout/routes/admin-routes';
import { chefeTrafegoRoutes } from './layout/routes/chefe-trafego-routes';
import { superAdminRoutes } from './layout/routes/super-admin-routes';


const routes: Routes = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: '',
        redirectTo: 'dashboard'
      },
      {
        path: 'profile',
        loadChildren: () => import('./edit-user/profile/profile.module').then(m => m.ProfileModule),   
      },
      {
        path: 'inBox',
        loadChildren: () => import('./inBox/inBox.module').then(m => m.InBoxModule),   
      },
      {
        path: 'changePass',
        loadChildren: () => import('./edit-user/edit-pass/edit-pass.module').then(m => m.EditPassModule),
      },
      ...superAdminRoutes,
      ...chefeTrafegoRoutes,
      ...adminRoutes
      
    ]
  }
  
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule {}
