import {Component} from '@angular/core';
import jwtDecode from "jwt-decode";
import {GpolicajacService} from "../../services/gp/gpolicajac.service";
import {PrelazakService} from "../../services/gp/prelazak.service";
import {ProveraService} from "../../services/gp/provera.service";
import {ProveraRequest} from "../../model/GP/proveraRequest";
import {Observable} from "rxjs";
import {ProveraGradjanina} from "../../model/GP/proveraGradjanina";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {KrivicnaPrijava} from "../../model/tuzilastvo/krivicnaPrijava";
import {PrijavaDAO} from "../../model/GP/prijavaDAO";
import {DatePipe} from "@angular/common";
import {Status} from "../../model/tuzilastvo/statusEnum";
import {KrivicnaPrijavaServiceService} from "../../services/tuzilastvo/krivicna-prijava-service.service";
import {PrelazakRequest} from "../../model/GP/prelazakRequest";
import {PrelazakGranice} from "../../model/GP/prelazakGranice";


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
  provere!: Observable<ProveraGradjanina[]>
  prelasci!: Observable<PrelazakGranice[]>
  provera: ProveraGradjanina = new ProveraGradjanina()
  trenutnoNaPrijavi: string
  closeResult = '';
  prijavaDao: PrijavaDAO = new PrijavaDAO();
  datepipe: DatePipe = new DatePipe('en-US');


  constructor(private gpolicajacService: GpolicajacService, private prelazakService: PrelazakService,
              private proveraService: ProveraService,private prijavaService : KrivicnaPrijavaServiceService, private modalService : NgbModal) {
    this.getPid()
    this.getProvereNaCekanju()
    this.getPrelasci()
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
      proveraReq.Gradjanin = this.proveraJmbgInput
      // @ts-ignore
      proveraReq.PolicajacId = localStorage.getItem('gpolicajacId')
      console.log(proveraReq)
      this.proveraService.postProvera(proveraReq).subscribe(response => {
        console.log('Number sent successfully:', response);
      }, error => {
        console.error('Error sending number:', error);
      })
    }
  }

  getProvereNaCekanju() {
    this.provere = this.proveraService.getProvereNaCekanju()
    console.log(this.provere)
  }

  getPrelasci(){
    this.prelasci = this.prelazakService.getPrelaske()
  }

  openPrijava(prijavi: any, provera: ProveraGradjanina) {
    console.log(provera.Id)
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
    krivicnaPrijava.privatnost = false;
    krivicnaPrijava.clanZakonika = this.prijavaDao.ClanZakonika;
    const now = new Date();
    krivicnaPrijava.datum != null ? this.datepipe.transform(now, 'dd/MM/YYYY') : String;
    krivicnaPrijava.mestoPrijave = this.provera.Policajac.GPrelaz;
    krivicnaPrijava.obrazlozenje = this.prijavaDao.Obrazlozenje
    krivicnaPrijava.status = Status.NACEKANJU
    krivicnaPrijava.optuzeni.ime = this.provera.Gradjanin.Ime
    krivicnaPrijava.optuzeni.prezime = this.provera.Gradjanin.Prezime
    krivicnaPrijava.optuzeni.jmbg = this.provera.Gradjanin.Jmbg
    krivicnaPrijava.gradjaninId = this.provera.Gradjanin.Id

    this.proveraService.postZabrani(this.provera.Id).subscribe(data => {
      console.log(data);
      this.getProvereNaCekanju()
    })
    this.prijavaService.createPrijava(krivicnaPrijava).subscribe(data => {
      console.log(data);
    })

  }

  acceptProvera(proveraId: string){

    this.proveraService.postPusti(proveraId).subscribe(data => {
      console.log(data);
      this.getProvereNaCekanju()
    })
    const prelazakreq: PrelazakRequest = new PrelazakRequest()
    prelazakreq.ProveraId = proveraId
    // @ts-ignore
    prelazakreq.PolicajacId = localStorage.getItem('gpolicajacId')
    this.prelazakService.postPrelazak(prelazakreq).subscribe(data => {
      console.log(data);
      this.getPrelasci()
    })
  }



}
