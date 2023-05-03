import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {

  constructor(private router : Router){
  }

  //Ovo treba da se promeni u gradjanin i da se doda sta treba
  //Samo sam kopirao iz prethodnog proj. , cisto da ima nesto
  user : any = new Object;

  submitted : boolean = false;

  login() {
    this.router.navigateByUrl("/");
  }
}