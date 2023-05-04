import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { KorisnikService } from '../services/mup/korisnik.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {

  constructor(private router : Router, private korisnikService: KorisnikService){
  }

  //Ovo treba da se promeni u gradjanin i da se doda sta treba
  //Samo sam kopirao iz prethodnog proj. , cisto da ima nesto
  user : any = new Object;

  submitted : boolean = false;

   login() {
    const jmbg = document.getElementById("username") as HTMLInputElement;
    const sifra = document.getElementById("password") as HTMLInputElement;
    this.korisnikService.loginUser(jmbg.value, sifra.value).subscribe(data => {
        localStorage.setItem("jmbg", data.jmbg)
        this.router.navigateByUrl("/mup")
    })
  }
}
