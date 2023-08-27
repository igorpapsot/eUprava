import { Component } from '@angular/core';
import {tap} from "rxjs";
import {GpolicajacService} from "../../services/gp/gpolicajac.service";
import {PrelazakService} from "../../services/gp/prelazak.service";
import {ProveraService} from "../../services/gp/provera.service";
import {KrivicnaPrijavaServiceService} from "../../services/tuzilastvo/krivicna-prijava-service.service";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {ProveraGradjanina} from "../../model/GP/proveraGradjanina";

@Component({
  selector: 'app-provere',
  templateUrl: './provere.component.html',
  styleUrls: ['./provere.component.css']
})
export class ProvereComponent {

  provere!: ProveraGradjanina[]
  constructor(private proveraService: ProveraService) {
    this.getProvereNaCekanju()
  }

  getProvereNaCekanju() {
    this.proveraService.getProvere().pipe(
      tap(data => console.log('Recived data', data))
    ).subscribe(data =>{
      this.provere = data;
      console.log('1', this.provere)
    })
  }
}
