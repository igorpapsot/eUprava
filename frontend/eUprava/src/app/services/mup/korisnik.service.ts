import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Korisnik } from "src/app/model/mup/korisnik";

@Injectable()
export class KorisnikService {

    constructor(private httpClient: HttpClient){}

    loginUser(jmbg: string, sifra: string) {
        return this.httpClient.post<Korisnik>("http://localhost:8004/user/login", {
            jmbg: jmbg,
            sifra: sifra
        })
    }

    getUserByJmbg(jmbg: string) {
        return this.httpClient.get<Korisnik>("http://localhost:8004/user/jmbg?jmbg=" + jmbg)
    }
}