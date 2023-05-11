import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Rociste } from 'src/app/model/sudstvo/rociste';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RocisteService {

  constructor(private client: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postRociste(rociste : Rociste){
    console.log(rociste)
    return this.client.post<unknown>(environment.apiUrl + "/rocista", {
      datum: rociste.datum,
      mesto: rociste.mesto,
      sud: rociste.sud
    }, this.options())
  }

  getRocista() : Observable<Rociste[]> {
    return this.client.get<Rociste[]>(environment.apiUrl + "/rocista")
  }
}
