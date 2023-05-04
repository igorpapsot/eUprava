import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Zahtev } from "src/app/model/mup/zahtev";

@Injectable()
export class ZahtevService {

    constructor(private httpClient: HttpClient){}

    createRequest(korisnikId: string, zahtevTip: number, dokumentTip: number) {
        return this.httpClient.post<Zahtev>("http://localhost:8004/req", {
            korisnikId: korisnikId,
            zahtevTip: zahtevTip,
            dokumentTip: dokumentTip
        })
    }
}