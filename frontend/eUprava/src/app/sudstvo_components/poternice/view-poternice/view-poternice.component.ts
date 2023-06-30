import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Poternica } from 'src/app/model/sudstvo/poternica';
import { PoternicaService } from 'src/app/services/sudstvo/poternica.service';

@Component({
  selector: 'app-view-poternice',
  templateUrl: './view-poternice.component.html',
  styleUrls: ['./view-poternice.component.css']
})
export class ViewPoterniceComponent implements OnInit{

  poternice!: Poternica[];
  constructor(private poternicaService: PoternicaService, private router: Router) {

  }

  ngOnInit(): void {
      // this.getPoternice();
  }

  // private getPoternice(){
  //   this.poternicaService.getPoternica().subscribe(data => {
  //     this.poternice = data;
  //   })
  // }

  poterniceDetails(sudijaId:string) {
    this.router.navigate(['poternice-list', sudijaId]);
  }

}
