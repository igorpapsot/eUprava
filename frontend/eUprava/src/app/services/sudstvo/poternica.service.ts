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
    return this.client.post<unknown>(environment.apiUrl + "/sudstvo/poternice", {
      ime: poternica.ime,
      opis: poternica.opis,
    }, this.options())
  }

  getPoternice() : Observable<Poternica[]> {
    return this.client.get<Poternica[]>(environment.apiUrl + "/sudstvo/poternice")
  }

  getPoternica(id: string) : Observable<Poternica> {
    return this.client.get<Poternica>(environment.apiUrl + "/sudstvo/poternice" + id);
  }
}
