import { Component } from '@angular/core';
//import * as jwtDecode from 'jwt-decode';

@Component({
  selector: 'app-granica',
  templateUrl: './granica.component.html',
  styleUrls: ['./granica.component.css']
})
export class GranicaComponent {
  //jwtToken: string = localStorage.getItem('pjwt');
  decodedToken: any;

  constructor() {
    //this.decodedToken = jwtDecode(this.jwtToken);
    console.log(this.decodedToken);
  }

}
