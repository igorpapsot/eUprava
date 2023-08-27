import {Component} from '@angular/core';
import jwtDecode from "jwt-decode";
import {GpolicajacService} from "../../services/gp/gpolicajac.service";
import {PrelazakService} from "../../services/gp/prelazak.service";
import {ProveraService} from "../../services/gp/provera.service";
import {ProveraRequest} from "../../model/GP/proveraRequest";
import {Observable, tap} from "rxjs";
import {ProveraGradjanina} from "../../model/GP/proveraGradjanina";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {KrivicnaPrijava} from "../../model/tuzilastvo/krivicnaPrijava";
import {PrijavaDAO} from "../../model/GP/prijavaDAO";
import {DatePipe} from "@angular/common";
import {Status} from "../../model/tuzilastvo/statusEnum";
import {KrivicnaPrijavaServiceService} from "../../services/tuzilastvo/krivicna-prijava-service.service";
import {PrelazakRequest} from "../../model/GP/prelazakRequest";
import {PrelazakGranice} from "../../model/GP/prelazakGranice";
import {EGPrelaz} from "../../model/GP/EGPrelaz";
import {Optuzeni} from "../../model/tuzilastvo/optuzeni";
import {Tuzilastvo} from "../../model/tuzilastvo/tuzilastvo";
import {TuzilastvoService} from "../../services/tuzilastvo/tuzilastvo.service";


@Component({
  selector: 'app-granica',
  templateUrl: './granica.component.html',
  styleUrls: ['./granica.component.css']
})
export class GranicaComponent {

  jwtToken: string = 'token';
  decodedToken: any;
  logedPolicajac = false

  proveraJmbgInput: string
  provera: ProveraGradjanina = new ProveraGradjanina()

  provere: ProveraGradjanina[]
  prikaz: boolean = false
  provereSve: ProveraGradjanina[]
  prelasci: PrelazakGranice[]

  trenutnoNaPrijavi: string
  tuzilastva! : Observable<Tuzilastvo[]>
  closeResult = '';
  prijavaDao: PrijavaDAO = new PrijavaDAO();
  tuzilastvo: Tuzilastvo = new Tuzilastvo();
  datepipe: DatePipe = new DatePipe('en-US');


  constructor(private gpolicajacService: GpolicajacService, private prelazakService: PrelazakService,
              private proveraService: ProveraService,private prijavaService : KrivicnaPrijavaServiceService,
              private tuzilastvoService : TuzilastvoService,
              private modalService : NgbModal) {

    this.getPid()
    this.getProvereNaCekanju()
    this.getProvereSve()
    this.getPrelasci()
    this.getTuzilastva()
  }

  getPid() {
    if (localStorage.getItem('pjwt') != null){
      // @ts-ignore ne moze bili null jer je provereno u ifu tkd error nema poentu
      this.jwtToken = localStorage.getItem('pjwt')
    }
    if (this.jwtToken != 'token') {
      this.decodedToken = jwtDecode(this.jwtToken);
      this.logedPolicajac = true
      const policajacId: string = this.decodedToken.Id
      localStorage.setItem('gpolicajacId', policajacId)
      console.log(this.decodedToken);
    }
  }

  postProvera() {
    if (this.proveraJmbgInput !== undefined) {
      const proveraReq: ProveraRequest = new ProveraRequest();
      proveraReq.gradjanin = this.proveraJmbgInput.toString()
      // @ts-ignore
      proveraReq.policajacId = localStorage.getItem('gpolicajacId')
      console.log(proveraReq)
      this.proveraService.postProvera(proveraReq).subscribe(response => {
        console.log('Number sent successfully:', response);
        this.getProvereNaCekanju()
      }, error => {
        console.error('Error sending number:', error);
      })
    }
  }

  getProvereNaCekanju() {
    this.proveraService.getProvereNaCekanju().subscribe(data =>{
      this.provere = data;
      console.log('provere', this.provere)
    })
  }

  getProvereSve() {
    this.proveraService.getProvere().subscribe(data =>{
      this.provereSve = data;
    })
  }

  getPrelasci(){
    this.prelazakService.getPrelaske().subscribe(data =>{
      this.prelasci = data;
      console.log('prelasci', this.prelasci)
    })
  }

  getTuzilastva() {
    this.tuzilastva = this.tuzilastvoService.getTuzilastva()
  }

  openPrijava(prijavi: any, provera: ProveraGradjanina) {
    console.log('trenutna provera',provera.id)
    this.provera = provera;
    this.modalService.open(prijavi, { ariaLabelledBy: 'prijavi' }).result.then(
      (result) => {
        console.log("3");
        this.trenutnoNaPrijavi = ''
        this['closeResult'] = `Closed with: ${result}`;
      },
      (reason) => {
        console.log("4")
        this.trenutnoNaPrijavi = ''
        this['closeResult'] = `Dismissed ${(reason)}`;
      },
    );
  }

  sendPrijava() {
    const krivicnaPrijava: KrivicnaPrijava = new KrivicnaPrijava();
    const optuzeni: Optuzeni = new Optuzeni();
    optuzeni.ime = this.provera.gradjanin.ime
    optuzeni.prezime = this.provera.gradjanin.prezime
    optuzeni.jmbg = this.provera.gradjanin.jmbg
    const now = new Date();
    krivicnaPrijava.datum != null ? this.datepipe.transform(now, 'dd/MM/YYYY') : String;
    krivicnaPrijava.optuzeni = optuzeni
    krivicnaPrijava.privatnost = false;
    krivicnaPrijava.status = Status.NACEKANJU
    krivicnaPrijava.clanZakonika = this.prijavaDao.ClanZakonika;
    krivicnaPrijava.obrazlozenje = this.prijavaDao.Obrazlozenje;
    krivicnaPrijava.tuzilastvoId = this.prijavaDao.TuzilastvoId;
    krivicnaPrijava.mestoPrijave = this.provera.policajac.prelaz;
    krivicnaPrijava.gradjaninId = this.provera.gradjanin.id;
    krivicnaPrijava.tuzilastvoId = this.tuzilastvo.id;

    this.proveraService.postZabrani(this.provera.id).subscribe(data => {
      console.log(data);
      this.getProvereNaCekanju()
    })
    this.prijavaService.createPrijava(krivicnaPrijava).subscribe(data => {
      console.log(data);
    })

  }

  acceptProvera(proveraId: string){
    const prelazakreq: PrelazakRequest = new PrelazakRequest()
    prelazakreq.proveraId = proveraId
    // @ts-ignore
    prelazakreq.policajacId = localStorage.getItem('gpolicajacId')
    console.log(prelazakreq)
    this.proveraService.postPusti(proveraId).subscribe(data => {
      console.log(data);
      this.getProvereNaCekanju()
      this.prelazakService.postPrelazak(prelazakreq).subscribe(data => {
        console.log(data);
        this.getPrelasci()
      })
    })

  }

  hideShow() {
    if (!this.prikaz) {
      this.prikaz = true
    } else {
      this.prikaz = false
    }
  }

}
