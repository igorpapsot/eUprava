import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { KrivicnaPrijavaServiceService } from 'src/app/services/tuzilastvo/krivicna-prijava-service.service';
import { Status } from 'src/app/model/tuzilastvo/statusEnum';
import { StatusPipe } from 'src/app/model/tuzilastvo/statusEnum';

@Component({
  selector: 'app-tuzilastvo',
  templateUrl: './tuzilastvo.component.html',
  styleUrls: ['./tuzilastvo.component.css']
})
export class TuzilastvoComponent {
[x: string]: any;

  constructor(private prijavaService : KrivicnaPrijavaServiceService){
    this.getPrijave()
  }

getPrijave(){
  this.prijave = this.prijavaService.getPrijave()
}
  prijave!: Observable<KrivicnaPrijava[]>

}
