import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Zahtev } from "src/app/model/mup/zahtev";

@Injectable()
export class ZahtevService {

    constructor(private httpClient: HttpClient){}

    createRequest(korisnikId: string, zahtevTip: number, dokumentTip: number, zakazanDatumVreme: string, datumIsticanja: string|undefined, jmbgDeteta: string|undefined) {
        return this.httpClient.post<Zahtev>("http://localhost:8004/req", {
            korisnikId: korisnikId,
            zahtevTip: zahtevTip,
            dokumentTip: dokumentTip,
            zakazanDatumVreme: zakazanDatumVreme,
            datumIsticanja: datumIsticanja,
            jmbgDeteta: jmbgDeteta
        })
    }

    getRequests(korisnikId: string) {
        return this.httpClient.get<Zahtev[]>("http://localhost:8004/req/user?id=" + korisnikId)
    }
}