import { Component } from '@angular/core';
import { PoternicaService } from '../../services/sudstvo/poternica.service';
import { RocisteService } from '../../services/sudstvo/rociste.service';
import { SudstvoService } from '../../services/sudstvo/sudstvo.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Poternica } from '../../model/sudstvo/poternica';

@Component({
  selector: 'app-sudstvo-page',
  templateUrl: './sudstvo-page.component.html',
  styleUrls: ['./sudstvo-page.component.css']
})
export class SudstvoPageComponent {


  constructor(private poternicaService: PoternicaService, private rocisteService: RocisteService, private sudstvoService: SudstvoService, private modalService: NgbModal) {
   
  }


}
