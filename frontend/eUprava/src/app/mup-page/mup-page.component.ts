import { Component } from '@angular/core';
import { KorisnikService } from '../services/mup/korisnik.service';
import { ZahtevService } from '../services/mup/zahtev.service';

@Component({
  selector: 'app-mup-page',
  templateUrl: './mup-page.component.html',
  styleUrls: ['./mup-page.component.css']
})
export class MupPageComponent {
    constructor(private korisnikService: KorisnikService, private zahtevService: ZahtevService) {}

    createRequest() {
        const zahtevTip = document.getElementById("zahtevTip") as HTMLSelectElement;
        const dokumentTip = document.getElementById("dokumentTip") as HTMLSelectElement;
        this.korisnikService.getUserByJmbg(localStorage.getItem("jmbg") as string).subscribe(korisnik => {
            this.zahtevService.createRequest(korisnik.id, Number.parseInt(zahtevTip.value), Number.parseInt(dokumentTip.value)).subscribe(data => {console.log(data)})
        })
    }
}
