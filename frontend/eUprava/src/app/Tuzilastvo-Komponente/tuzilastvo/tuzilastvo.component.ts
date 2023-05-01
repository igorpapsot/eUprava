import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { KrivicnaPrijavaServiceService } from 'src/app/services/tuzilastvo/krivicna-prijava-service.service';
import { Status } from 'src/app/model/tuzilastvo/statusEnum';
import { StatusPipe } from 'src/app/model/tuzilastvo/statusEnum';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-tuzilastvo',
  templateUrl: './tuzilastvo.component.html',
  styleUrls: ['./tuzilastvo.component.css']
})
export class TuzilastvoComponent {
[x: string]: any;

  constructor(private prijavaService : KrivicnaPrijavaServiceService, private modalService : NgbModal){
    this.getPrijave()
  }

  getPrijave(){
    this.prijave = this.prijavaService.getPrijave()
  }

  declinePrijava(id: string){
    this.prijavaService.declinePrijava(id).subscribe(data => {
      console.log(data);
      this.getPrijave();
    })
  }

  confirmPrijava(id: string){
    this.prijavaService.confirmPrijava(id).subscribe(data => {
      console.log(data);
      this.getPrijave();
    })
  }
 
  prijave!: Observable<KrivicnaPrijava[]>

  closeResult = '';

	openOdbaci(odbacivanje: any, id: string) {
    console.log(id)
		this.modalService.open(odbacivanje, { ariaLabelledBy: 'odbaci' }).result.then(
			(result) => {
        console.log("1");
        this.declinePrijava(id);

				this['closeResult'] = `Closed with: ${result}`;
			},
			(reason) => {
        console.log("2")
				this['closeResult'] = `Dismissed ${(reason)}`;
			},
		);
	}

  openPrihvati(prihvatanje: any, id: string) {
    console.log(id)
		this.modalService.open(prihvatanje, { ariaLabelledBy: 'prihvati' }).result.then(
			(result) => {
        console.log("3");
        this.confirmPrijava(id);

				this['closeResult'] = `Closed with: ${result}`;
			},
			(reason) => {
        console.log("4")
				this['closeResult'] = `Dismissed ${(reason)}`;
			},
		);
	}


}
