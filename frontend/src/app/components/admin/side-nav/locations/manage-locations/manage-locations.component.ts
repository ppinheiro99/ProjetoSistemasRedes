import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import { LocationService } from 'src/app/services/Locations/locations.service';
export interface LocationsData {
  ID: any;
  name: String;
  latitude: any;
  longitude: any;
}
@Component({
  selector: 'app-manage-company',
  templateUrl: './manage-locations.component.html',
  styleUrls: ['./manage-locations.component.scss']
})
export class ManageLocationComponent implements OnInit {
  displayedColumns = ['select','ID','name','latitude','longitude','delete','update']
  dataSource: MatTableDataSource<LocationsData>
  selection: SelectionModel<LocationsData>
  validationForm: FormGroup
  editForm = false
  errorMessage
  name_edit
  location_id
  latitude_edit
  longitude_edit
  location 

  addLocationDialog = false;
  @ViewChild(MatPaginator, { static: true }) paginator: MatPaginator;
  @ViewChild(MatSort, { static: true }) sort: MatSort;


  constructor(private locationService : LocationService, public fb: FormBuilder,private readonly router: Router) {
    this.validationForm = fb.group({
      name: [null, Validators.required],
      latitude: [null, Validators.required],
      longitude: [null, Validators.required],
      name_edit: [null, Validators.required],
      latitude_edit: [null, Validators.required],
      longitude_edit: [null, Validators.required],
      id_editLocation: [null, Validators.required],
     });
   }

  ngOnInit(): void {
    this.locationService.getData().subscribe(data =>{
      this.dataSource = new MatTableDataSource(data.data);
    })
    
    this.selection = new SelectionModel<LocationsData>(true, []);
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
  addLocation(){
    this.addLocationDialog = true;
  }

  closeAddLocation(){
    this.addLocationDialog = false;
  }

  closeEdit(){
    this.editForm = false;
  }

  insertLocation(){
    console.log(this.validationForm.value)
    this.locationService.addLocation(this.validationForm.value).subscribe(
      data => {
        this.addLocationDialog = false;
        this.locationService.getData().subscribe
        this.reloadPage()
      },
      err => {
        this.errorMessage = err.error.message;
      }
    );
  }

  deleteLocation(id){
    this.locationService.deleteLocation(id).subscribe(data =>{
      this.reloadPage();
    })
  }
 
  editLocation(id){
    this.addLocationDialog = false;
    this.location = this.locationService.getLocation(id).subscribe(
    data => {
      console.log(data.data.ID) 
      this.editForm = true;
      this.name_edit = data.data.name
      this.latitude_edit = data.data.latitude
      this.longitude_edit = data.data.longitude
      this.location_id = data.data.ID
      
    },
   );
      
  }
 
  get nameEdit() { return this.validationForm.get('name_edit') }
  get latitudeEdit() { return this.validationForm.get('latitude_edit') }
  get longitudeEdit() { return this.validationForm.get('longitude_edit') }
  get locationIDEdit() { return this.validationForm.get('id_editLocation') }
  
  confirmEditLocation(){
    this.locationService.updateLocation(this.validationForm.value).subscribe(
      data => {
        console.log(this.validationForm.value)
        this.editForm = false;
        this.reloadPage();
      },
      err => {
        this.errorMessage = err.error.message;
      }
    );
  }

  reloadPage(): void {
    window.location.reload();
  }

}
