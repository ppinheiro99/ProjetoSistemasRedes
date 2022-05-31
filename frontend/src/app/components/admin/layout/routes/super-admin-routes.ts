export const superAdminRoutes = [
  {
    path: 'dashboard',
    loadChildren: () =>
      import('../../dashboard/dashboard.module').then(m => m.DashboardModule),
    data: { icon: 'dashboard', text: 'Dashboard' }
  },
  {
    path: 'tables',
    loadChildren: () =>
      import('../../side-nav/show-users/tables.module').then(m => m.TablesModule),
    data: { icon: 'table_chart', text: 'Listar Funcionarios' }
  },
  {
    path: 'trucks',
    loadChildren: () => import('../../side-nav/trucks/trucks.module').then(m => m.TrucksModule),
    data: { icon: 'directions_car', text: 'Camiões/Reboques' }
  },
  {
    path: 'register',
    loadChildren: () => import('../../side-nav/register/register.module').then(m => m.RegisterModule),
    data: { icon: 'assignment', text: 'Registar Utilizadores' }
  },
  {
    path: 'maps',
    loadChildren: () =>
      import('../../side-nav/map/map.module').then(
        m => m.MapModule
      ),
    data: { icon: 'place', text: 'Mapa' }
  },
  {
    path: 'location',
    loadChildren: () =>
      import('../../side-nav/locations/location.module').then(m => m.CompanyModule),
    data: { icon: 'store', text: 'Localizações' }
  },
];
