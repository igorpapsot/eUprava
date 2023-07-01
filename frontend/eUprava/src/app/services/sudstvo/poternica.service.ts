import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Poternica } from 'src/app/model/sudstvo/poternica';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PoternicaService {

  constructor(private client: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postPoternica(poternica : Poternica){
    console.log(poternica)
    return this.client.post<unknown>(environment.apiUrl + "/mup/poternica", {
      gradjaninId: poternica.gradjaninId,
      sudijaId: poternica.sudijaId,
      naslov: poternica.naslov,
      opis: poternica.opis
    }, this.options())
  }


  getPoternica(id: string) : Observable<Poternica[]> {
    return this.client.get<Poternica[]>(environment.apiUrl + "/mup/poternica" + id);
  }

  getPoternicaSudija(sudijaId: string) : Observable<Poternica[]> {
    return this.client.get<Poternica[]>(environment.apiUrl + "/mup/poternica" + sudijaId);
  }

  getPoternicaGradjanin(gradjaninId: string) : Observable<Poternica> {
    return this.client.get<Poternica>(environment.apiUrl + "/mup/poternica" + gradjaninId);
  }
}
