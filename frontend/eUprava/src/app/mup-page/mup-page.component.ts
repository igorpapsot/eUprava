import { Component, OnInit } from '@angular/core';
import { KorisnikService } from '../services/mup/korisnik.service';
import { ZahtevService } from '../services/mup/zahtev.service';
import { Zahtev } from '../model/mup/zahtev';
import { DokumentTip } from '../model/mup/dokumentTipEnum';
import { Router } from '@angular/router';

@Component({
    selector: 'app-mup-page',
    templateUrl: './mup-page.component.html',
    styleUrls: ['./mup-page.component.css']
})
export class MupPageComponent implements OnInit {
    constructor(private korisnikService: KorisnikService, private zahtevService: ZahtevService, private router: Router) { }
    ngOnInit(): void {
        this.korisnikService.getUserByJmbg(localStorage.getItem("jmbg") as string).subscribe(korisnik => {
            this.zahtevService.getRequests(korisnik.id).subscribe(zahtevi => {
                this.zahtevi = zahtevi;
            })
        })
    }

    convertDokumentTipToString(tip: number) {
        if (tip == DokumentTip.LICNA) {
            return "Licna karta"
        }
        if (tip == DokumentTip.PASOS) {
            return "Pasos"
        }
        return ""
    }

    zahtevTip: number

    zahtevi: Zahtev[] = []

    getZahteviZaProduzavanje() {
        return this.zahtevi.filter(z => z.zahtevTip === 0)
    }

    getZahteviZaIzgubljena() {
        return this.zahtevi.filter(z => z.zahtevTip === 1)
    }

    getZahteviZaDete() {
        return this.zahtevi.filter(z => z.zahtevTip === 2)
    }

    createRequest() {
        const zahtevTip = document.getElementById("zahtevTip") as HTMLSelectElement;
        const dokumentTip = document.getElementById("dokumentTip") as HTMLSelectElement;
        const zakazivanjeDatum = document.getElementById("zakazivanjeDatum") as HTMLInputElement;
        const zakazivanjeSat = document.getElementById("sati") as HTMLInputElement;
        const zakazivanjeMinut = document.getElementById("minuti") as HTMLInputElement;
        const zakazivanjeDatumVreme = new Date(zakazivanjeDatum.value);
        let datumIsticanja: string | undefined = undefined
        let jmbgDeteta: string | undefined = undefined
        if (zahtevTip.value == "0") {
            datumIsticanja = (document.getElementById("isticanjeDatum") as HTMLInputElement).value;
        }
        if (zahtevTip.value == "2") {
            jmbgDeteta = (document.getElementById("deteJmbg") as HTMLInputElement).value
        }
        zakazivanjeDatumVreme.setHours(Number.parseInt(zakazivanjeSat.value) + zakazivanjeDatumVreme.getTimezoneOffset()/-60, Number.parseInt(zakazivanjeMinut.value))
        this.korisnikService.getUserByJmbg(localStorage.getItem("jmbg") as string).subscribe(korisnik => {
            this.zahtevService.createRequest(korisnik.id, Number.parseInt(zahtevTip.value), Number.parseInt(dokumentTip.value), zakazivanjeDatumVreme.toISOString(), datumIsticanja, jmbgDeteta).subscribe(data => { 
                window.location.reload()
            })
        })
    }


}
