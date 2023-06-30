import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Sudija } from 'src/app/model/sudstvo/sudija';
import { SudijaService } from 'src/app/services/sudstvo/sudija.service';

@Component({
  selector: 'app-view-sudije',
  templateUrl: './view-sudije.component.html',
  styleUrls: ['./view-sudije.component.css']
})
export class ViewSudijeComponent implements OnInit{

  sudije!: Sudija[];
  constructor(private sudijaService: SudijaService, private router: Router) {

  }

  ngOnInit(): void {
    this.getSudije();  
  }

  private getSudije() {
    this.sudijaService.getSudije().subscribe(data => {
      this.sudije = data;
    })
  }

  sudijeDetails(jmbg: string) {
    this.router.navigate(['sudije-list', jmbg]);
  }
}
