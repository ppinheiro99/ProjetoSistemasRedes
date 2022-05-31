import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';
import { TrucksService } from "../../../../../services/trucks/trucks.service";
import { TrailersService } from "../../../../../services/trailers/trailers.service";
import { UsersService } from "../../../../../services/user/users.service";
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { Router } from '@angular/router';
export interface TruckData {
  ID: any;
  plate: String;
  year: any;
  month: any;
  km: any;
  brand: String;
}
export interface UserData {
  ID: any;
  email: String;
  role_id: any;
}

export interface TrailerData {
  ID: any;
  plate: String;
  year: any;
}

@Component({
  selector: 'app-trucks',
  templateUrl: './trucks.component.html',
  styleUrls: ['./trucks.component.scss']
})

export class TrucksComponent implements OnInit, AfterViewInit {
  displayedColumns = ['select','ID','plate', 'year','month','km','brand','button','update','associate','camionistas', 'addRoute'];
  displayedTrailerColumns = ['select','ID','plate', 'year', 'button','update'];

  dataSource: MatTableDataSource<TruckData>;
  dataSourceTrailer: MatTableDataSource<TrailerData>;
  selection: SelectionModel<TruckData>;
  selectionTrailer: SelectionModel<TrailerData>;
  validationForm: FormGroup
  dataUserSource: MatTableDataSource<UserData>;
  selectionUser: SelectionModel<UserData>;
  displayedUserColumns = ['select','ID','first_name','last_name','email', 'role_id','bind'];
  idCamiao: any
  idReboque: any
  addRouteDialog = false

  dataTruckDriverSource: any;
  selectionTruckDriver: SelectionModel<TruckData>;
  displayedTruckDriverColumns = ['select','ID','first_name','last_name'];
  
  dialog = false;
  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;
  constructor(private trucksService : TrucksService, private userService : UsersService, public fb: FormBuilder,private readonly router: Router, private trailersService : TrailersService) {
    this.validationForm = fb.group({
      plate: [null, Validators.required],
      year: [null, Validators.required],
      month: [null, Validators.required],
      km: [null, Validators.required],
      brand: [null, Validators.required],
      plate_editTruck: [null, Validators.required],
      year_editTruck: [null, Validators.required],
      month_editTruck: [null, Validators.required],
      km_editTruck: [null, Validators.required],
      brand_editTruck: [null, Validators.required],
      id_editTruck: [null, Validators.required],
      latitudeRoute: [null, Validators.required],
      longitudeRoute: [null, Validators.required], 

      plateTrailer: [null, Validators.required],
      yearTrailer: [null, Validators.required],
      plate_editTrailer: [null, Validators.required],
      year_editTrailer: [null, Validators.required],
      id_editTrailer: [null, Validators.required],
  });

  }
  
  
  ngOnInit() {
    this.trucksService.getData().subscribe(data =>{
      this.dataSource = new MatTableDataSource(data.data);
    })
    
    this.selection = new SelectionModel<TruckData>(true, []);

    this.userService.getTruckDrivers().subscribe(data =>{
      this.dataUserSource = new MatTableDataSource(data.data);
    })
    
    this.selectionUser = new SelectionModel<UserData>(true, []);

    this.trailersService.getData().subscribe(data =>{
      this.dataSourceTrailer = new MatTableDataSource(data.data);
    })  

    this.selectionTrailer = new SelectionModel<TrailerData>(true, []);


  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
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


  ///funçao do botao adicionar , quando clicado abre o formulario para inserir os dados
  openDialog(){
    this.dialog = true;
  }

  closeDialog(){
    this.dialog = false;
  }

  closeEdit(){
    this.editForm = false;
  }

  get plate() { return this.validationForm.get('plate') }
  get year() { return this.validationForm.get('year') }
  get month() { return this.validationForm.get('month') }
  get km() { return this.validationForm.get('km') }
  get brand() { return this.validationForm.get('brand') }


  plateFailed = false;
  yearFailed = false;
  errorMessage 
  insertTruck(){
    this.trucksService.addTruck(this.validationForm.value).subscribe(
      data => {
        this.dialog = false;
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!" ||err.error.message == "Mes Invalido!" ){
          this.yearFailed = true
        }
      }
    );
  }
  deleteTruck(id){
    this.trucksService.deleteTruck(id).subscribe(data =>{
      window.location.reload();
    })
  }
 
  truck
  plate_edit
  truck_id
  year_edit
  month_edit
  km_edit
  brand_edit
  editForm = false
  editTruck(id){
    this.dialog = false;
   this.truck = this.trucksService.getTruck(id).subscribe(
    data => {
      this.editForm = true;
      this.plate_edit = data.data.plate
      this.year_edit = data.data.year
      this.month_edit = data.data.month
      this.km_edit = data.data.km
      this.brand_edit = data.data.brand
      this.truck_id = data.data.ID
      
    },
   );
      
  }
 
  get plateEdit() { return this.validationForm.get('plate_editTruck') }
  get yearEdit() { return this.validationForm.get('year_editTruck') }
  get monthEdit() { return this.validationForm.get('month_editTruck') }
  get kmEdit() { return this.validationForm.get('km_editTruck') }
  get brandEdit() { return this.validationForm.get('brand_editTruck') }
  get truckIDEdit() { return this.validationForm.get('id_editTruck') }
  
  confirmEditTruck(){
    this.trucksService.updateTruck(this.validationForm.value).subscribe(
      data => {
        this.editForm = false;
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!" ||err.error.message == "Mes Invalido!" ){
          this.yearFailed = true
        }
      }
    );
  }

  assocCamionista = false;
  idTruck = 0 ;
  assocCamionistaform(id){
    this.dialog = false;
    this.editForm = false;
    this.assocCamionista = true;
    this.idTruck=id;
    ///recebemos o id do camiao que queremos associar
    ///De seguida listamos no formulário a lista dos camionistas


  }

  addRoute(id){
    this.idCamiao = id
    this.addRouteDialog = !this.addRouteDialog
  }


  get latitudeRoute() { return this.validationForm.get('latitudeRoute') }
  get longitudeRoute() { return this.validationForm.get('longitudeRoute') }
  insertRoute(){
    this.trucksService.createRoute(this.validationForm.value,this.idCamiao).subscribe(
      data => {
        this.addRouteDialog = !this.addRouteDialog
       // this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
      }
    );
  }
  
  bindTruckAndDriver(DriverId,TruckId){
    //Recebemos por argumento o id do camionista e o do camiao a associar
    
    this.trucksService.bindTruckAndDriver(DriverId,TruckId).subscribe(
      data => {
        
        this.dialog = false;
        this.editForm = false;
        this.assocCamionista = false;
        
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!" ||err.error.message == "Mes Invalido!" ){
          this.yearFailed = true
        }
      }
    );
  }

  showCamionistas = false;
  showCamionistasTable(id){
    this.dialog = false;
    this.editForm = false;
    this.assocCamionista = false;
    this.showCamionistas = true;
    this.idTruck=id;
    
    this.trucksService.getTruckDriver(this.idTruck).subscribe(data =>{
      this.dataTruckDriverSource = data.data;
      
    })

    this.selectionTruckDriver = new SelectionModel<TruckData>(true, []);
    ///recebemos o id do camiao que queremos associar
    ///De seguida listamos no formulário a lista dos camionistas

  }

  removerCamionista(id){
    this.trucksService.unbindTruckDriver(id).subscribe(
      data => {
        
        this.dialog = false;
        this.editForm = false;
        this.assocCamionista = false;
        this.showCamionistas = false;
        
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!" ||err.error.message == "Mes Invalido!" ){
          this.yearFailed = true
        }
      }
    );
  }

  closeDesassociar(){
    this.showCamionistas = false;
  }

  // TRAILER
  dialogTrailer = false;
  trailer: any
  plate_editTrailer: any
  trailer_id: any
  year_editTrailer: any
  editFormTrailer = false

  get plateTrailer() { return this.validationForm.get('plateTrailer') }
  get yearTrailer() { return this.validationForm.get('yearTrailer') }

  get plateEditTrailer() { return this.validationForm.get('plate_editTrailer') }
  get yearEditTrailer() { return this.validationForm.get('year_editTrailer') }
  get trailerIDEdit() { return this.validationForm.get('id_editTrailer') }

  openDialogTrailer(){
    this.dialogTrailer = true;
    this.editFormTrailer = false;
  }

  insertTrailer(){
    this.trailersService.addTrailer(this.validationForm.value).subscribe(
      data => {
        this.dialogTrailer = false;
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!"){
          this.yearFailed = true
        }
      }
    );
  }

  deleteTrailer(id){
    this.trailersService.deleteTrailer(id).subscribe(data =>{
      window.location.reload();
    })
  }

  editTrailer(id){
    this.trailer = this.trailersService.getTrailer(id).subscribe(
     data => {
       this.editFormTrailer = true;
       this.dialogTrailer = false;
       this.plate_editTrailer = data.data.plate
       this.year_editTrailer = data.data.year
       this.trailer_id = data.data.ID
     })
  }
 
  confirmEditTrailer(){
    this.trailersService.updateTrailer(this.validationForm.value).subscribe(
      data => {
        this.editFormTrailer = false;
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
        if(err.error.message == "Matricula Invalida!" || err.error.message == "Matricula já existe!"){
          this.plateFailed=true
        }else if (err.error.message == "Ano Invalido!"){
          this.yearFailed = true
        }
      }
    );
  }

  closeDialogTrailer(){
    this.dialogTrailer = false;
  }

  closeEditTrailer(){
    this.editFormTrailer = false;
  }

  reloadPage(): void {
    window.location.reload();
  }

}

