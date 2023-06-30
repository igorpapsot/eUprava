import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Rociste } from 'src/app/model/sudstvo/rociste';
import { RocisteService } from 'src/app/services/sudstvo/rociste.service';

@Component({
  selector: 'app-rociste-list',
  templateUrl: './rociste-list.component.html',
  styleUrls: ['./rociste-list.component.css']
})
export class RocisteListComponent implements OnInit{
  id!: string;
  rociste!: Rociste;

  constructor(private route: ActivatedRoute, private router: Router, private rocisteService: RocisteService) {

  }

  ngOnInit(): void {
      this.id = this.route.snapshot.params['id'];
      this.rocisteService.getRociste(this.id).subscribe(data => {
        this.rociste = data;
      })
  }

}
