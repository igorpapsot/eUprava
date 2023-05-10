import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { KorisnikService } from '../services/mup/korisnik.service';
import { LoginService } from '../services/login.service';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {

  constructor(private router : Router, private korisnikService: KorisnikService, private loginS: LoginService){
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
        this.loginS.login()
        this.router.navigateByUrl("/mup")
    })
  }
}
