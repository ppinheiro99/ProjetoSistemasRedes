export const chefeTrafegoRoutes = [
  {
    path: 'dashboard',
    loadChildren: () =>
      import('../../dashboard/dashboard.module').then(m => m.DashboardModule),
    data: { icon: 'dashboard', text: 'Dashboard' }
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
  {
    path: 'trucks',
    loadChildren: () => import('../../side-nav/trucks/trucks.module').then(m => m.TrucksModule),
    data: { icon: 'directions_car', text: 'Camiões/Reboques' }
  },
];
