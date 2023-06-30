import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Poternica } from 'src/app/model/sudstvo/poternica';
import { PoternicaService } from 'src/app/services/sudstvo/poternica.service';

@Component({
  selector: 'app-poternice-list',
  templateUrl: './poternice-list.component.html',
  styleUrls: ['./poternice-list.component.css']
})
export class PoterniceListComponent implements OnInit{ 
  id!: string;
  poternica!: Poternica;

  constructor(private route: ActivatedRoute, private router: Router, private poternicaService: PoternicaService) {

  }

  ngOnInit(): void {
      this.id = this.route.snapshot.params['id'];
      this.poternicaService.getPoternica(this.id).subscribe(data => {
          this.poternica = data;
      })
  }
}
