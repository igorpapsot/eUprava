import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { KrivicnaPrijavaServiceService } from 'src/app/services/tuzilastvo/krivicna-prijava-service.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Optuzeni } from 'src/app/model/tuzilastvo/optuzeni';
import { Mesto } from 'src/app/model/tuzilastvo/mesto';
import { DatePipe } from '@angular/common';
import { Tuzilastvo } from 'src/app/model/tuzilastvo/tuzilastvo';
import { TuzilastvoService } from 'src/app/services/tuzilastvo/tuzilastvo.service';
import { Status } from 'src/app/model/tuzilastvo/statusEnum';
import { Optuznica } from 'src/app/model/tuzilastvo/optuznica';
import { OptuzniceService } from 'src/app/services/tuzilastvo/optuznice.service';

@Component({
  selector: 'app-tuzilastvo',
  templateUrl: './tuzilastvo.component.html',
  styleUrls: ['./tuzilastvo.component.css']
})
export class TuzilastvoComponent {
[x: string]: any;

  constructor(private prijavaService : KrivicnaPrijavaServiceService, private modalService : NgbModal, private tuzilastvoService : TuzilastvoService,
    private optuzniceService: OptuzniceService){
    this.getPrijave();
    this.getTuzilastva();
    this.getJavnePrijave();
    this.prijava.optuzeni = this.optuzeni;
    this.prijava.optuzeni.mestoPrebivalista = this.mesto;
  }

  prikaziPrijave() {
    this.prikazOptuznice = false;
    this.getPrijave();
  }

  prikaziOptuznice() {
    this.prikazOptuznice = true;
    this.getOptuznice();
  }

  getPrijave(){
    this.prijave = this.prijavaService.getPrijave()
  }

  getJavnePrijave(){
    this.javnePrijave = this.prijavaService.getJavnePrijave()
  }

  getTuzilastva() {
    this.tuzilastva = this.tuzilastvoService.getTuzilastva()
  }

  getOptuznice() {
    this.optuznice = this.optuzniceService.getOptuznice()
  }

  declinePrijava(id: string){
    this.prijavaService.declinePrijava(id).subscribe(data => {
      console.log(data);
      this.getPrijave();
    })
  }

  confirmPrijava(id: string){
    this.okp.id = id;
    this.optuznica.krivicnaPrijava = this.okp;
    this.optuznica.idGradjanina = "tempIdGradjanina";
    this.optuznica.aktivna = true;
    this.prijavaService.confirmPrijava(this.optuznica).subscribe(data => {
      console.log(data);
      this.getPrijave();
    })
  }

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

  openDodavanje(dodavanje: any) {
		this.modalService.open(dodavanje, { ariaLabelledBy: 'dodavanje' }).result.then(
			(result) => {
        console.log("5");

				this['closeResult'] = `Closed with: ${result}`;
			},
			(reason) => {
        console.log("6")
				this['closeResult'] = `Dismissed ${(reason)}`;
			},
		);
	}


  createPrijava() {
    const datepipe: DatePipe = new DatePipe('en-US');

    this.prijava.datum = datepipe.transform(this.prijava.datum, 'dd/MM/YYYY');
    this.prijava.optuzeni.datumRodjenja = datepipe.transform(this.prijava.optuzeni.datumRodjenja, 'dd/MM/YYYY');

    if(this.privatnost == "Javna") {
      this.prijava.privatnost = true
      console.log("da")
    }
    else {
      this.prijava.privatnost = false
      console.log("ne")
    }

    //Dodaj da se cita iz jwta
    this.prijava.gradjaninId = "tempIdGradjanina"
    this.prijava.tuzilastvoId = this.tuzilastvo.id
    this.prijava.status = Status.NACEKANJU

    this.prijavaService.createPrijava(this.prijava).subscribe(data => {
      console.log(data);
      this.getPrijave();
    })
    
  }

  prijava: KrivicnaPrijava = new KrivicnaPrijava();
  optuzeni: Optuzeni = new Optuzeni();
  mesto: Mesto = new Mesto();
  privatnost: string;
  tuzilastvo: Tuzilastvo = new Tuzilastvo();
  optuznica: Optuznica = new Optuznica();
  okp: KrivicnaPrijava = new KrivicnaPrijava();

  prijave!: Observable<KrivicnaPrijava[]>
  javnePrijave!: Observable<KrivicnaPrijava[]>
  optuznice!: Observable<Optuznica[]>
  tuzilastva! : Observable<Tuzilastvo[]>;

  prikazOptuznice: boolean = false;
}