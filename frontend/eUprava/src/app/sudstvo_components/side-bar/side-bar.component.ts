import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.css']
})
export class SideBarComponent implements OnInit{
  constructor(private router: Router) {}

  ngOnInit(): void {
      
  }

  createPoternica() {
    this.router.navigateByUrl('/create-poternica');
  }

  createRociste() {
    this.router.navigateByUrl('/create-rociste');

  }

  registerSudija() {
    this.router.navigateByUrl('/register-sudija');

  }
}
