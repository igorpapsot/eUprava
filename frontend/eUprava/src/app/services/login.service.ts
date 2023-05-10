import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor() {
    if (localStorage.getItem('jmbg')) {
      this.logged = true
    }
    else {
      this.logged = false
    }
   }

  public getLogged() : boolean{
    return this.logged
  }

  public login() {
    this.logged = true
  }

  public logout() {
    this.logged = false
  }

  logged = false;
}
