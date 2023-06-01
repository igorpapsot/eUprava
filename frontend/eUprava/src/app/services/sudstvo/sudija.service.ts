import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Sudija } from 'src/app/model/sudstvo/sudija';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SudijaService {

  constructor(private client : HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  registerSudija(sudija : Sudija) {
    console.log(sudija)
    return this.client.post<unknown>(environment.apiUrl + "/sudstvo/register", {
      ime : sudija.ime,
      prezime : sudija.prezime,
      pol : sudija.pol,
      jmbg : sudija.jmbg,
      lozinka : sudija.lozinka,
      sud : sudija.sud,
    }, this.options())
  }

  getSudije() : Observable<Sudija[]> {
    return this.client.get<Sudija[]>(environment.apiUrl + "/sudstvo/sudije")
  }

  getSudija(jmbg: string) : Observable<Sudija> {
    return this.client.get<Sudija>(environment.apiUrl + "/sudstvo/sudija" + jmbg);
  }

}
