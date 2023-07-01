import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Optuznica } from 'src/app/model/sudstvo/optuznica';
import { OptuznicaService } from 'src/app/services/sudstvo/optuznica.service';

@Component({
  selector: 'app-view-optuznice',
  templateUrl: './view-optuznice.component.html',
  styleUrls: ['./view-optuznice.component.css']
})
export class ViewOptuzniceComponent implements OnInit{

  optuznice!: Optuznica[];
  constructor(private optuznicaService: OptuznicaService, private router: Router) {

  }

  ngOnInit(): void {
      this.getOptuznice();
  }

  private getOptuznice(){
    this.optuznicaService.getOptuznice().subscribe(data => {
      this.optuznice = data;
    })
  }

}
