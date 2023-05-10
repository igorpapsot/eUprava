import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from '../services/login.service';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent {
  constructor(private router: Router, private loginS: LoginService) {
    this.logOut()
  }

  logOut() {
    this.loginS.logout()
    localStorage.removeItem('jmbg');
    this.router.navigateByUrl('')
  }
}
