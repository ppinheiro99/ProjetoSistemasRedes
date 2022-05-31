import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';
import { UsersService } from "../../../../../services/user/users.service";
import { TokenService } from "../../../../../services/token/token.service";
import { MapService } from 'src/app/services/map/map.service';
import { TrucksService } from 'src/app/services/trucks/trucks.service';

export interface UserData {
  ID: any;
  email: String;
  role_id: any;
}

export interface truckDriverMapData {
  ID: any;
  truck_id:any
	trailer_id:any
	coords:any
	distance:any
	time:any
	start_country:any
	start_city:any
	start_postal_code:any
	start_address:any
	end_country:any
	end_city:any
	end_postal_code:any
	end_address:any
}

@Component({
  selector: 'app-tables',
  templateUrl: './tables.component.html',
  styleUrls: ['./tables.component.scss']
})

export class TablesComponent implements OnInit, AfterViewInit {
  user_role : any
  user_role_id:any
  user_id:any
  displayedColumns = ['select','ID','email', 'role_id', 'button', 'truckDriverMap'];
  displayedColumnsTruckDriver = ['ID','truck_id', 'trailer_id', 'distance', 'time', 'start_country', 'start_city','start_postal_code', 'start_address', 'end_country', 'end_city','end_postal_code', 'end_address'];
  dataSource: MatTableDataSource<UserData>;
  selection: SelectionModel<UserData>;
  dataSourceTruckDriver: MatTableDataSource<truckDriverMapData>;
  selectionTruckDriver: SelectionModel<truckDriverMapData>;
  arrayTrucks = []
  arrayTruckAndDriver = []
  showDataTruckDriver= false
  dataLocation: any
  distance = 0
  distance1 = 0

  
  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;
  constructor(private tokenService: TokenService, private userService : UsersService, private mapService: MapService, private trucksService: TrucksService) {}

  ngOnInit() {
    const user = this.tokenService.getUser()
    this.user_role_id = user.Role
    this.user_id = user.ID
      this.userService.getData().subscribe(data =>{
        this.dataSource = new MatTableDataSource(data.data);
    })

    this.mapService.listTrucks().subscribe(
      data => {
        this.arrayTrucks= data.data
      }
    )
    this.mapService.listTruckAndDriver().subscribe(
      data => {
        this.arrayTruckAndDriver= data.data
      }
    )

    this.selection = new SelectionModel<UserData>(true, []);
  }

  ngAfterViewInit() { /// Ver isto melhor e meter a funcionar
   // this.dataSource.paginator = this.paginator;
   // this.dataSource.sort = this.sort;
  }

  applyFilter(filterValue: string) {
    this.dataSource.filter = filterValue.trim().toLowerCase();
    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected()
      ? this.selection.clear()
      : this.dataSource.data.forEach(row => this.selection.select(row));
  }

  delete(id_remove){
    this.userService.deleteUser(id_remove).subscribe(data =>{
      window.location.reload();
    })
  }

  truckDriverMap(id_truckDriver){
    this.distance = 0;
    this.distance1 = 0;
    this.showDataTruckDriver = false
    // buscar o camiao correspondente ao truckDriver
    const auxTruckAndDriver = this.arrayTruckAndDriver.find(x => x.first_driver_id == id_truckDriver )
    let truck
    if(auxTruckAndDriver != null){
      truck = this.arrayTrucks.find(x => x.ID == auxTruckAndDriver.truck_id)
    }else{
      const auxTruckAndDriver = this.arrayTruckAndDriver.find(x => x.second_driver_id == id_truckDriver )
      truck = this.arrayTrucks.find(x => x.ID == auxTruckAndDriver.truck_id)
    }

    // Buscar o mapa do motorista e no .html criar uma tabele e exibir as informações todas
    this.userService.getTravelMap(truck.ID).subscribe(
      data =>{
        this.dataLocation = data
        if(this.dataLocation.data.length > 0){
          this.showDataTruckDriver = true
          this.dataSourceTruckDriver = new MatTableDataSource(this.dataLocation.data);
          this.distance1 = this.dataLocation.data[0].distance
        }
      }
    )
    // mandar para o método já feito do percurso e buscar os KM
    console.warn(truck.ID)
    this.trucksService.getTruckHistory(truck.ID).subscribe(
      async data =>{
        console.warn(data)
        console.warn(this.mapService.getTruckHistory(data,truck.ID))
        await this.delay(500); // delay para nao fazerem sincronos e esperar que o this.mapService.distance receba os valores
        console.warn(this.mapService.distance)
        this.distance = this.mapService.distance
    })
  }

  closeTruckDriverData(){
    this.showDataTruckDriver = false
  }

  delay(ms: number) {
    return new Promise( resolve => setTimeout(resolve, ms) );
  }
}
