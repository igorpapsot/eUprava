import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Zahtev } from "src/app/model/mup/zahtev";
import { environment } from "src/environments/environment";

@Injectable()
export class ZahtevService {

    constructor(private httpClient: HttpClient){}

    createRequest(korisnikId: string, zahtevTip: number, dokumentTip: number, zakazanDatumVreme: string, datumIsticanja: string|undefined, jmbgDeteta: string|undefined) {
        return this.httpClient.post<Zahtev>(environment.apiUrl + "/mup/req", {
            korisnikId: korisnikId,
            zahtevTip: zahtevTip,
            dokumentTip: dokumentTip,
            zakazanDatumVreme: zakazanDatumVreme,
            datumIsticanja: datumIsticanja,
            jmbgDeteta: jmbgDeteta
        })
    }

    getRequests(korisnikId: string) {
        return this.httpClient.get<Zahtev[]>(environment.apiUrl + "/mup/req/user?id=" + korisnikId)
    }
}