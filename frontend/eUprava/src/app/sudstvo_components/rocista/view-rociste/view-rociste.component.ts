import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Rociste } from 'src/app/model/sudstvo/rociste';
import { RocisteService } from 'src/app/services/sudstvo/rociste.service';

@Component({
  selector: 'app-view-rociste',
  templateUrl: './view-rociste.component.html',
  styleUrls: ['./view-rociste.component.css']
})
export class ViewRocisteComponent implements OnInit{

  rocista : Rociste[];
  constructor(private rocisteService: RocisteService, private router: Router) {

  }

  ngOnInit(): void {
      this.getRocista();
  }

  private getRocista(){
    this.rocisteService.getRocista().subscribe(data => {
      this.rocista = data;
    })
  }

  // rocistaDetails(id:number) {
  //   this.router.navigate(['rociste-list', id]);
  // }

}
