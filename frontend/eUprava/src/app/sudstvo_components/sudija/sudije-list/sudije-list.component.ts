import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Sudija } from 'src/app/model/sudstvo/sudija';
import { SudijaService } from 'src/app/services/sudstvo/sudija.service';

@Component({
  selector: 'app-sudije-list',
  templateUrl: './sudije-list.component.html',
  styleUrls: ['./sudije-list.component.css']
})
export class SudijeListComponent implements OnInit{
  jmbg!: string;
  sudija!: Sudija;

  constructor(private router: Router, private route: ActivatedRoute, private sudijaService: SudijaService) {

  }

  ngOnInit(): void {
    this.jmbg = this.route.snapshot.params['jmbg'];
    this.sudijaService.getSudija(this.jmbg).subscribe(data => {
        this.sudija = data;
    })
  }

}
