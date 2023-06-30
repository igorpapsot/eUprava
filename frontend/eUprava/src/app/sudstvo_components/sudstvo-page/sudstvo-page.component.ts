import { Component } from '@angular/core';
import { PoternicaService } from '../../services/sudstvo/poternica.service';
import { RocisteService } from '../../services/sudstvo/rociste.service';
import { SudstvoService } from '../../services/sudstvo/sudstvo.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { OptuznicaService } from 'src/app/services/sudstvo/optuznica.service';
import { Poternica } from 'src/app/model/sudstvo/poternica';
import { Rociste } from 'src/app/model/sudstvo/rociste';
import { KonacnaPresuda } from 'src/app/model/sudstvo/konacnaPresuda';
import { Optuznica } from 'src/app/model/sudstvo/optuznica';
import { Sud } from 'src/app/model/sudstvo/sud';
import { Observable } from 'rxjs';
import { Sudija } from 'src/app/model/sudstvo/sudija';
import { SudijaService } from 'src/app/services/sudstvo/sudija.service';
import { KonacnaPresudaService } from 'src/app/services/sudstvo/konacna-presuda.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-sudstvo-page',
  templateUrl: './sudstvo-page.component.html',
  styleUrls: ['./sudstvo-page.component.css']
})
export class SudstvoPageComponent {

  constructor(private router: Router ,private poternicaService: PoternicaService, private rocisteService: RocisteService, private sudstvoService: SudstvoService, private modalService: NgbModal,
    private optuznicaService: OptuznicaService, private sudijaService: SudijaService, private konacnaPresudaService: KonacnaPresudaService) {
      
  }

  openLg(content : any) {
    this.modalService.open(content, { size: 'lg' });
  }


  poternica: Poternica = new Poternica();
  rociste: Rociste = new Rociste();
  konacnaPresuda: KonacnaPresuda = new KonacnaPresuda();
  optuznica: Optuznica = new Optuznica();
  sud: Sud = new Sud();

  poternice!: Observable<Poternica[]>
  optuznice!: Observable<Optuznica[]>
  sudovi! : Observable<Sud[]>
  rocista! : Observable<Rociste[]>
  konacnePresude! : Observable<KonacnaPresuda[]>

  prikazSudovi: boolean = true;
  prikazOptuznice: boolean = true;
  prikazPoternice: boolean = true;
  prikazRocista: boolean = true;
  prikazKonacnuPresudu: boolean = true;
  sudija: Sudija = new Sudija()
  sudijaUlogovan: boolean = false
  jmbg: string | null = null



  newPoternica() {
    this.poternicaService.postPoternica(this.poternica).subscribe(data => {
      console.log(data);
    })
  }

  newRociste(){
    this.rocisteService.postRociste(this.rociste).subscribe(data => {
      console.log(data);
    })
  }

  newKonacnaPresuda(){
    this.konacnaPresudaService.postKonacnaPresuda(this.konacnaPresuda).subscribe(data => {
      console.log(data);
    })
  }

  newSudija() {
    this.sudijaService.registerSudija(this.sudija).subscribe(data => {
      console.log(data);
    })
  }


  prikaziOptuznice(){
    this.prikazOptuznice = true;
    this.getOptuznice();
  }

  prikaziSudove() {
    this.prikazSudovi = true;
    this.getSudovi();
  }

  prikaziRocista() {
    this.prikazRocista = true;
    this.getRocista();
  }

  getSudovi() {
    this.sudovi = this.sudstvoService.getSudovi();
  }

  getOptuznice() {
    this.optuznice = this.optuznicaService.getOptuznice();
  }

  getRocista() {
    this.rocista = this.rocisteService.getRocista();
  }

  getKonacnaPresuda(id: string | null) : KonacnaPresuda {
    if(id) {
      this.konacnaPresudaService.getKonacnaPresuda(id).subscribe(data => {
        console.log("Konacna Presuda")
        console.log(data)
        return data
      })
    }
    return new KonacnaPresuda
  }


  getPoternica(id: string | null) : Poternica {
    if(id) {
      this.poternicaService.getPoternica(id).subscribe(data => {
        console.log("Poternica")
        console.log(data)
        return data
      })
    }
    return new Poternica
  }

  getSudija(jmbg: string | null) : Sudija {
    if(jmbg) {
      this.sudijaService.getSudija(jmbg).subscribe(data => {
        console.log("Sudija")
        console.log(data)
        return data
      })
    }
    return new Sudija
  }
  
}