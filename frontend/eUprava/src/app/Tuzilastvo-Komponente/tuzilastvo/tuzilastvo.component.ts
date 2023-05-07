import { Component, Directive, EventEmitter, Input, Output, QueryList, ViewChildren } from '@angular/core';
import { Observable } from 'rxjs';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { KrivicnaPrijavaServiceService } from 'src/app/services/tuzilastvo/krivicna-prijava-service.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Optuzeni } from 'src/app/model/tuzilastvo/optuzeni';
import { Mesto } from 'src/app/model/tuzilastvo/mesto';
import { DatePipe, DecimalPipe, NgFor } from '@angular/common';
import { Tuzilastvo } from 'src/app/model/tuzilastvo/tuzilastvo';
import { TuzilastvoService } from 'src/app/services/tuzilastvo/tuzilastvo.service';
import { Status } from 'src/app/model/tuzilastvo/statusEnum';
import { Optuznica } from 'src/app/model/tuzilastvo/optuznica';
import { OptuzniceService } from 'src/app/services/tuzilastvo/optuznice.service';
import { Event } from '@angular/router';


export type SortColumn = keyof KrivicnaPrijava | '';
export type SortDirection = 'asc' | 'desc' | '';
const rotate: { [key: string]: SortDirection } = { asc: 'desc', desc: '', '': 'asc' };

const compare = (v1: string | number, v2: string | number) => (v1 < v2 ? -1 : v1 > v2 ? 1 : 0);

export interface SortEvent {
	column: SortColumn;
	direction: SortDirection;
}

@Directive({
	selector: 'th[sortable]',
	host: {
		'[class.asc]': 'direction === "asc"',
		'[class.desc]': 'direction === "desc"',
		'(click)': 'rotate()',
	},
})
export class NgbdSortableHeader {
	@Input() sortable: SortColumn = '';
	@Input() direction: SortDirection = '';
	@Output() sort = new EventEmitter<SortEvent>();

	rotate() {
		this.direction = rotate[this.direction];
		this.sort.emit({ column: this.sortable, direction: this.direction });
	}
}


@Component({
  selector: 'app-tuzilastvo',
  templateUrl: './tuzilastvo.component.html',
  styleUrls: ['./tuzilastvo.component.css']
})
export class TuzilastvoComponent {
[x: string]: any;

  datepipe: DatePipe = new DatePipe('en-US');

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
    this.prijavaService.getPrijave().subscribe(data => {
      this.prijave = data;
  })
  }

  getJavnePrijave(){
    this.prijavaService.getJavnePrijave().subscribe(data => {
      this.javnePrijave = data;
  })
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

    this.prijava.datum != null ? this.datepipe.transform(this.prijava.datum, 'dd/MM/YYYY') : String;
    
    this.prijava.optuzeni.datumRodjenja = this.datepipe.transform(this.prijava.optuzeni.datumRodjenja, 'dd/MM/YYYY');

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

  public prijave!: KrivicnaPrijava[]
  javnePrijave!: KrivicnaPrijava[]
  optuznice!: Observable<Optuznica[]>
  tuzilastva! : Observable<Tuzilastvo[]>;

  prikazOptuznice: boolean = false;

  @ViewChildren(NgbdSortableHeader) headers: QueryList<NgbdSortableHeader>;
  
  onSort({ column, direction }: SortEvent) {
    this.headers.forEach((header: { sortable: string; direction: string; }) => {
			if (header.sortable !== column) {
				header.direction = '';
			}
		});

		if (direction === '' || column === '') {
			this.prijave = this.prijave;
		} else {
      this.prijave = this.prijave
			this.prijave = [...this.prijave].sort((a, b) => {
        if (column == "clanZakonika") {
				  const res = compare(a[column], b[column]);
				  return direction === 'asc' ? res : -res;
        }
        else if (column == "mestoPrijave") {
				  const res = compare(a[column], b[column]);
				  return direction === 'asc' ? res : -res;
        }
        else if (column == "datum") {

				  const [dayA, monthA, yearA] = a[column].split('/');
          const [dayB, monthB, yearB] = b[column].split('/');
          
          var valueA = new Date(+yearA, +monthA - 1, +dayA)
          var valueB = new Date(+yearB, +monthB - 1, +dayB)
          
          
          console.log(valueA)
          console.log(valueB)

          if(valueA.getTime() > valueB.getTime()) {
            console.log("1")
            return direction === 'asc' ? 1 : -1;
          }
          else if (valueA.getTime() < valueB.getTime()) {
            console.log("-1")
            return direction === 'asc' ? -1 : 1;
          }
          else {
            console.log("0")
            return direction === 'asc' ? 0 : -0;
          }
          
        }
        else {
          const res = compare("1", "1");
          return direction === 'asc' ? res : -res;
        }
			});
		}
	}


  onSortJavne({ column, direction }: SortEvent) {
    this.headers.forEach((header: { sortable: string; direction: string; }) => {
			if (header.sortable !== column) {
				header.direction = '';
			}
		});

		if (direction === '' || column === '') {
			this.javnePrijave = this.javnePrijave;
		} else {
      this.javnePrijave = this.javnePrijave
			this.javnePrijave = [...this.javnePrijave].sort((a, b) => {
        if (column == "clanZakonika") {
				  const res = compare(a[column], b[column]);
				  return direction === 'asc' ? res : -res;
        }
        else if (column == "datum") {

          const [dayA, monthA, yearA] = a[column].split('/');
          const [dayB, monthB, yearB] = b[column].split('/');
          
          var valueA = new Date(+yearA, +monthA - 1, +dayA)
          var valueB = new Date(+yearB, +monthB - 1, +dayB)
          
          
          console.log(valueA)
          console.log(valueB)

          if(valueA.getTime() > valueB.getTime()) {
            console.log("1")
            return direction === 'asc' ? 1 : -1;
          }
          else if (valueA.getTime() < valueB.getTime()) {
            console.log("-1")
            return direction === 'asc' ? -1 : 1;
          }
          else {
            console.log("0")
            return direction === 'asc' ? 0 : -0;
          }
  
        }
        else if (column == "mestoPrijave") {
				  const res = compare(a[column], b[column]);
				  return direction === 'asc' ? res : -res;
        }
        else {
          const res = compare("1", "1");
          return direction === 'asc' ? res : -res;
        }
			});
		}
	}
}