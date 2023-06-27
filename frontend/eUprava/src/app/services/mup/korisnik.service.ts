import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Korisnik } from "src/app/model/mup/korisnik";
import { environment } from "src/environments/environment";

@Injectable()
export class KorisnikService {

    constructor(private httpClient: HttpClient){}

    loginUser(jmbg: string, sifra: string) {
        return this.httpClient.post<Korisnik>(environment.apiUrl + "/mup/user/login", {
            jmbg: jmbg,
            sifra: sifra
        })
    }

    getUserByJmbg(jmbg: string) {
        return this.httpClient.get<Korisnik>(environment.apiUrl + "/mup/user/jmbg?jmbg=" + jmbg)
    }
}