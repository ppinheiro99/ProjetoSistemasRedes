import { Component, OnInit } from '@angular/core';
import { TrucksService } from "../../../../services/trucks/trucks.service";
interface Place {
  imgSrc: string;
  name: string;
  description: string;
  charge: string;
  location: string;
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  places: Array<Place> = [];
  teste = 50;
  rzero;
  rmid;
  rrun ;
  constructor(private trucksService : TrucksService) {}
   bgClass: string;
   icon: string;
  count: number;
   label: string;
  data: number;


  ngOnInit() {
this.icon =  "directions_car";
    this.trucksService.getCountTruck().subscribe(data =>{
      this.teste = data.data;
    })
    this.trucksService.getRpmPlates().subscribe(data =>{
      this.rzero = data.rpmzero;
      this.rmid = data.rpmmid;
      this.rrun = data.rpmrun;
    })



    this.places = [
      {
        imgSrc: 'assets/images/card-1.jpg',
        name: 'Cozy 5 Stars Apartment',
        description: `The place is close to Barceloneta Beach and bus stop just 2 min by walk and near to "Naviglio"
              where you can enjoy the main night life in Barcelona.`,
        charge: '$899/night',
        location: 'Barcelona, Spain'
      },
      {
        imgSrc: 'assets/images/card-2.jpg',
        name: 'Office Studio',
        description: `The place is close to Metro Station and bus stop just 2 min by walk and near to "Naviglio"
              where you can enjoy the night life in London, UK.`,
        charge: '$1,119/night',
        location: 'London, UK'
      },
      {
        imgSrc: 'assets/images/card-3.jpg',
        name: 'Beautiful Castle',
        description: `The place is close to Metro Station and bus stop just 2 min by walk and near to "Naviglio"
              where you can enjoy the main night life in Milan.`,
        charge: '$459/night',
        location: 'Milan, Italy'
      }
    ];
  }

  showRpmzero = false
  listPlatesRpmZero(){
    this.showRpmzero=true
  }
  showRpmmid = false
  listPlatesRpmMid(){
    this.showRpmmid=true
  }
  showRpmrun = false
  listPlatesRpmRun(){
    this.showRpmrun=true
  }
}


